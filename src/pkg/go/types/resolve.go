// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"fmt"
	"go/ast"
	"go/token"
	"strconv"
)

func (check *checker) declareObj(scope, altScope *Scope, obj Object) {
	alt := scope.Insert(obj)
	if alt == nil && altScope != nil {
		// see if there is a conflicting declaration in altScope
		alt = altScope.Lookup(obj.GetName())
	}
	if alt != nil {
		prevDecl := ""
		if pos := alt.GetPos(); pos.IsValid() {
			prevDecl = fmt.Sprintf("\n\tprevious declaration at %s", check.fset.Position(pos))
		}
		check.errorf(obj.GetPos(), fmt.Sprintf("%s redeclared in this block%s", obj.GetName(), prevDecl))
	}
}

func (check *checker) resolveIdent(scope *Scope, ident *ast.Ident) bool {
	for ; scope != nil; scope = scope.Outer {
		if obj := scope.Lookup(ident.Name); obj != nil {
			check.register(ident, obj)
			return true
		}
	}
	return false
}

func (check *checker) resolve(importer Importer) (pkg *Package, methods []*ast.FuncDecl) {
	// complete package scope
	pkgName := ""
	pkgScope := &Scope{Outer: Universe}

	i := 0
	for _, file := range check.files {
		// package names must match
		switch name := file.Name.Name; {
		case pkgName == "":
			pkgName = name
		case name != pkgName:
			check.errorf(file.Package, "package %s; expected %s", name, pkgName)
			continue // ignore this file
		}

		// keep this file
		check.files[i] = file
		i++

		// insert top-level file objects in package scope
		// (the parser took care of declaration errors)
		for _, decl := range file.Decls {
			switch d := decl.(type) {
			case *ast.BadDecl:
				// ignore
			case *ast.GenDecl:
				if d.Tok == token.CONST {
					check.assocInitvals(d)
				}
				for _, spec := range d.Specs {
					switch s := spec.(type) {
					case *ast.ImportSpec:
						// handled separately below
					case *ast.ValueSpec:
						for _, name := range s.Names {
							if name.Name == "_" {
								continue
							}
							pkgScope.Insert(check.lookup(name))
						}
					case *ast.TypeSpec:
						if s.Name.Name == "_" {
							continue
						}
						pkgScope.Insert(check.lookup(s.Name))
					default:
						check.invalidAST(s.Pos(), "unknown ast.Spec node %T", s)
					}
				}
			case *ast.FuncDecl:
				if d.Recv != nil {
					// collect method
					methods = append(methods, d)
					continue
				}
				if d.Name.Name == "_" || d.Name.Name == "init" {
					continue // blank (_) and init functions are inaccessible
				}
				pkgScope.Insert(check.lookup(d.Name))
			default:
				check.invalidAST(d.Pos(), "unknown ast.Decl node %T", d)
			}
		}
	}
	check.files = check.files[0:i]

	// package global mapping of imported package ids to package objects
	imports := make(map[string]*Package)

	// complete file scopes with imports and resolve identifiers
	for _, file := range check.files {
		// build file scope by processing all imports
		importErrors := false
		fileScope := &Scope{Outer: pkgScope}
		for _, spec := range file.Imports {
			if importer == nil {
				importErrors = true
				continue
			}
			path, _ := strconv.Unquote(spec.Path.Value)
			pkg, err := importer(imports, path)
			if err != nil {
				check.errorf(spec.Path.Pos(), "could not import %s (%s)", path, err)
				importErrors = true
				continue
			}
			// TODO(gri) If a local package name != "." is provided,
			// global identifier resolution could proceed even if the
			// import failed. Consider adjusting the logic here a bit.

			// local name overrides imported package name
			name := pkg.Name
			if spec.Name != nil {
				name = spec.Name.Name
			}

			// add import to file scope
			if name == "." {
				// merge imported scope with file scope
				for _, obj := range pkg.Scope.Entries {
					check.declareObj(fileScope, pkgScope, obj)
				}
			} else if name != "_" {
				// declare imported package object in file scope
				// (do not re-use pkg in the file scope but create
				// a new object instead; the Decl field is different
				// for different files)
				obj := &Package{Name: name, Scope: pkg.Scope, spec: spec}
				check.declareObj(fileScope, pkgScope, obj)
			}
		}

		// resolve identifiers
		if importErrors {
			// don't use the universe scope without correct imports
			// (objects in the universe may be shadowed by imports;
			// with missing imports, identifiers might get resolved
			// incorrectly to universe objects)
			pkgScope.Outer = nil
		}
		i := 0
		for _, ident := range file.Unresolved {
			if !check.resolveIdent(fileScope, ident) {
				check.errorf(ident.Pos(), "undeclared name: %s", ident.Name)
				file.Unresolved[i] = ident
				i++
			}

		}
		file.Unresolved = file.Unresolved[0:i]
		pkgScope.Outer = Universe // reset outer scope
	}

	return &Package{Name: pkgName, Scope: pkgScope, Imports: imports}, methods
}
