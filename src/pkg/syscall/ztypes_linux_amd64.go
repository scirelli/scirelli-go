// godefs -gsyscall -f-m64 types_linux.c

// MACHINE GENERATED - DO NOT EDIT.

package syscall

// Constants
const (
	sizeofPtr               = 0x8
	sizeofShort             = 0x2
	sizeofInt               = 0x4
	sizeofLong              = 0x8
	sizeofLongLong          = 0x8
	PathMax                 = 0x1000
	SizeofSockaddrInet4     = 0x10
	SizeofSockaddrInet6     = 0x1c
	SizeofSockaddrAny       = 0x70
	SizeofSockaddrUnix      = 0x6e
	SizeofSockaddrLinklayer = 0x14
	SizeofSockaddrNetlink   = 0xc
	SizeofLinger            = 0x8
	SizeofIpMreq            = 0x8
	SizeofMsghdr            = 0x38
	SizeofCmsghdr           = 0x10
	SizeofUcred             = 0xc
	IFA_UNSPEC              = 0
	IFA_ADDRESS             = 0x1
	IFA_LOCAL               = 0x2
	IFA_LABEL               = 0x3
	IFA_BROADCAST           = 0x4
	IFA_ANYCAST             = 0x5
	IFA_CACHEINFO           = 0x6
	IFA_MULTICAST           = 0x7
	IFLA_UNSPEC             = 0
	IFLA_ADDRESS            = 0x1
	IFLA_BROADCAST          = 0x2
	IFLA_IFNAME             = 0x3
	IFLA_MTU                = 0x4
	IFLA_LINK               = 0x5
	IFLA_QDISC              = 0x6
	IFLA_STATS              = 0x7
	IFLA_COST               = 0x8
	IFLA_PRIORITY           = 0x9
	IFLA_MASTER             = 0xa
	IFLA_WIRELESS           = 0xb
	IFLA_PROTINFO           = 0xc
	IFLA_TXQLEN             = 0xd
	IFLA_MAP                = 0xe
	IFLA_WEIGHT             = 0xf
	IFLA_OPERSTATE          = 0x10
	IFLA_LINKMODE           = 0x11
	IFLA_LINKINFO           = 0x12
	IFLA_NET_NS_PID         = 0x13
	IFLA_IFALIAS            = 0x14
	IFLA_MAX                = 0x14
	SizeofNlMsghdr          = 0x10
	SizeofNlMsgerr          = 0x14
	SizeofRtGenmsg          = 0x1
	SizeofNlAttr            = 0x4
	SizeofRtAttr            = 0x4
	SizeofIfInfomsg         = 0x10
	SizeofIfAddrmsg         = 0x8
	SizeofRtmsg             = 0xc
	SizeofInotifyEvent      = 0x10
)

// Types

type _C_short int16

type _C_int int32

type _C_long int64

type _C_long_long int64

type Timespec struct {
	Sec  int64
	Nsec int64
}

type Timeval struct {
	Sec  int64
	Usec int64
}

type Timex struct {
	Modes         uint32
	Pad_godefs_0  [4]byte
	Offset        int64
	Freq          int64
	Maxerror      int64
	Esterror      int64
	Status        int32
	Pad_godefs_1  [4]byte
	Constant      int64
	Precision     int64
	Tolerance     int64
	Time          Timeval
	Tick          int64
	Ppsfreq       int64
	Jitter        int64
	Shift         int32
	Pad_godefs_2  [4]byte
	Stabil        int64
	Jitcnt        int64
	Calcnt        int64
	Errcnt        int64
	Stbcnt        int64
	Tai           int32
	Pad_godefs_3  int32
	Pad_godefs_4  int32
	Pad_godefs_5  int32
	Pad_godefs_6  int32
	Pad_godefs_7  int32
	Pad_godefs_8  int32
	Pad_godefs_9  int32
	Pad_godefs_10 int32
	Pad_godefs_11 int32
	Pad_godefs_12 int32
	Pad_godefs_13 int32
}

type Time_t int64

type Tms struct {
	Utime  int64
	Stime  int64
	Cutime int64
	Cstime int64
}

type Utimbuf struct {
	Actime  int64
	Modtime int64
}

type Rusage struct {
	Utime    Timeval
	Stime    Timeval
	Maxrss   int64
	Ixrss    int64
	Idrss    int64
	Isrss    int64
	Minflt   int64
	Majflt   int64
	Nswap    int64
	Inblock  int64
	Oublock  int64
	Msgsnd   int64
	Msgrcv   int64
	Nsignals int64
	Nvcsw    int64
	Nivcsw   int64
}

type Rlimit struct {
	Cur uint64
	Max uint64
}

type _Gid_t uint32

type Stat_t struct {
	Dev       uint64
	Ino       uint64
	Nlink     uint64
	Mode      uint32
	Uid       uint32
	Gid       uint32
	X__pad0   int32
	Rdev      uint64
	Size      int64
	Blksize   int64
	Blocks    int64
	Atim      Timespec
	Mtim      Timespec
	Ctim      Timespec
	X__unused [3]int64
}

type Statfs_t struct {
	Type    int64
	Bsize   int64
	Blocks  uint64
	Bfree   uint64
	Bavail  uint64
	Files   uint64
	Ffree   uint64
	Fsid    [8]byte /* __fsid_t */
	Namelen int64
	Frsize  int64
	Spare   [5]int64
}

type Dirent struct {
	Ino          uint64
	Off          int64
	Reclen       uint16
	Type         uint8
	Name         [256]int8
	Pad_godefs_0 [5]byte
}

type RawSockaddrInet4 struct {
	Family uint16
	Port   uint16
	Addr   [4]byte /* in_addr */
	Zero   [8]uint8
}

type RawSockaddrInet6 struct {
	Family   uint16
	Port     uint16
	Flowinfo uint32
	Addr     [16]byte /* in6_addr */
	Scope_id uint32
}

type RawSockaddrUnix struct {
	Family uint16
	Path   [108]int8
}

type RawSockaddrLinklayer struct {
	Family   uint16
	Protocol uint16
	Ifindex  int32
	Hatype   uint16
	Pkttype  uint8
	Halen    uint8
	Addr     [8]uint8
}

type RawSockaddrNetlink struct {
	Family uint16
	Pad    uint16
	Pid    uint32
	Groups uint32
}

type RawSockaddr struct {
	Family uint16
	Data   [14]int8
}

type RawSockaddrAny struct {
	Addr RawSockaddr
	Pad  [96]int8
}

type _Socklen uint32

type Linger struct {
	Onoff  int32
	Linger int32
}

type Iovec struct {
	Base *byte
	Len  uint64
}

type IpMreq struct {
	Multiaddr [4]byte /* in_addr */
	Interface [4]byte /* in_addr */
}

type Msghdr struct {
	Name         *byte
	Namelen      uint32
	Pad_godefs_0 [4]byte
	Iov          *Iovec
	Iovlen       uint64
	Control      *byte
	Controllen   uint64
	Flags        int32
	Pad_godefs_1 [4]byte
}

type Cmsghdr struct {
	Len   uint64
	Level int32
	Type  int32
}

type Ucred struct {
	Pid int32
	Uid uint32
	Gid uint32
}

type NlMsghdr struct {
	Len   uint32
	Type  uint16
	Flags uint16
	Seq   uint32
	Pid   uint32
}

type NlMsgerr struct {
	Error int32
	Msg   NlMsghdr
}

type RtGenmsg struct {
	Family uint8
}

type NlAttr struct {
	Len  uint16
	Type uint16
}

type RtAttr struct {
	Len  uint16
	Type uint16
}

type IfInfomsg struct {
	Family     uint8
	X__ifi_pad uint8
	Type       uint16
	Index      int32
	Flags      uint32
	Change     uint32
}

type IfAddrmsg struct {
	Family    uint8
	Prefixlen uint8
	Flags     uint8
	Scope     uint8
	Index     uint32
}

type RtMsg struct {
	Family   uint8
	Dst_len  uint8
	Src_len  uint8
	Tos      uint8
	Table    uint8
	Protocol uint8
	Scope    uint8
	Type     uint8
	Flags    uint32
}

type InotifyEvent struct {
	Wd     int32
	Mask   uint32
	Cookie uint32
	Len    uint32
}

type PtraceRegs struct {
	R15      uint64
	R14      uint64
	R13      uint64
	R12      uint64
	Rbp      uint64
	Rbx      uint64
	R11      uint64
	R10      uint64
	R9       uint64
	R8       uint64
	Rax      uint64
	Rcx      uint64
	Rdx      uint64
	Rsi      uint64
	Rdi      uint64
	Orig_rax uint64
	Rip      uint64
	Cs       uint64
	Eflags   uint64
	Rsp      uint64
	Ss       uint64
	Fs_base  uint64
	Gs_base  uint64
	Ds       uint64
	Es       uint64
	Fs       uint64
	Gs       uint64
}

type FdSet struct {
	Bits [16]int64
}

type Sysinfo_t struct {
	Uptime       int64
	Loads        [3]uint64
	Totalram     uint64
	Freeram      uint64
	Sharedram    uint64
	Bufferram    uint64
	Totalswap    uint64
	Freeswap     uint64
	Procs        uint16
	Pad          uint16
	Pad_godefs_0 [4]byte
	Totalhigh    uint64
	Freehigh     uint64
	Unit         uint32
	X_f          [0]int8
	Pad_godefs_1 [4]byte
}

type Utsname struct {
	Sysname    [65]int8
	Nodename   [65]int8
	Release    [65]int8
	Version    [65]int8
	Machine    [65]int8
	Domainname [65]int8
}

type Ustat_t struct {
	Tfree        int32
	Pad_godefs_0 [4]byte
	Tinode       uint64
	Fname        [6]int8
	Fpack        [6]int8
	Pad_godefs_1 [4]byte
}

type EpollEvent struct {
	Events uint32
	Fd     int32
	Pad    int32
}
