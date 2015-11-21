package system

//NOTES:
//* syscall constants in the "syscall" package are nice, but some syscalls there are missing
//* future versions will include more than just the syscall name
//* 32bit (x86/i386) and 64bit (x86_64) syscall numbers are different

const (
	SyscallX86MinNum      = 0
	SyscallX86UnknownNum  = -1
	SyscallX86UnknownName = "unknown_syscall"

	SyscallX86MaxNum32   = 358
	SyscallX86LastName32 = "execveat"

	SyscallX86MaxNum64   = 322
	SyscallX86LastName64 = "execveat"
)

type NumberResolverFunc func(int16) string

func CallNumberResolver(arch ArchName) NumberResolverFunc {
	switch arch {
	case ArchName386:
		return callNameX86Family32
	case ArchNameAmd64:
		return callNameX86Family64
	default:
		return nil
	}
}

var syscallNumTableX86Family64 = [...]string{
	"read",
	"write",
	"open",
	"close",
	"stat",
	"fstat",
	"lstat",
	"poll",
	"lseek",
	"mmap",
	"mprotect",
	"munmap",
	"brk",
	"rt_sigaction",
	"rt_sigprocmask",
	"rt_sigreturn",
	"ioctl",
	"pread",
	"pwrite",
	"readv",
	"writev",
	"access",
	"pipe",
	"select",
	"sched_yield",
	"mremap",
	"msync",
	"mincore",
	"madvise",
	"shmget",
	"shmat",
	"shmctl",
	"dup",
	"dup2",
	"pause",
	"nanosleep",
	"getitimer",
	"alarm",
	"setitimer",
	"getpid",
	"sendfile",
	"socket",
	"connect",
	"accept",
	"sendto",
	"recvfrom",
	"sendmsg",
	"recvmsg",
	"shutdown",
	"bind",
	"listen",
	"getsockname",
	"getpeername",
	"socketpair",
	"setsockopt",
	"getsockopt",
	"clone",
	"fork",
	"vfork",
	"execve",
	"_exit",
	"wait4",
	"kill",
	"uname",
	"semget",
	"semop",
	"semctl",
	"shmdt",
	"msgget",
	"msgsnd",
	"msgrcv",
	"msgctl",
	"fcntl",
	"flock",
	"fsync",
	"fdatasync",
	"truncate",
	"ftruncate",
	"getdents",
	"getcwd",
	"chdir",
	"fchdir",
	"rename",
	"mkdir",
	"rmdir",
	"creat",
	"link",
	"unlink",
	"symlink",
	"readlink",
	"chmod",
	"fchmod",
	"chown",
	"fchown",
	"lchown",
	"umask",
	"gettimeofday",
	"getrlimit",
	"getrusage",
	"sysinfo",
	"times",
	"ptrace",
	"getuid",
	"syslog",
	"getgid",
	"setuid",
	"setgid",
	"geteuid",
	"getegid",
	"setpgid",
	"getppid",
	"getpgrp",
	"setsid",
	"setreuid",
	"setregid",
	"getgroups",
	"setgroups",
	"setresuid",
	"getresuid",
	"setresgid",
	"getresgid",
	"getpgid",
	"setfsuid",
	"setfsgid",
	"getsid",
	"capget",
	"capset",
	"rt_sigpending",
	"rt_sigtimedwait",
	"rt_sigqueueinfo",
	"rt_sigsuspend",
	"sigaltstack",
	"utime",
	"mknod",
	"uselib",
	"personality",
	"ustat",
	"statfs",
	"fstatfs",
	"sysfs",
	"getpriority",
	"setpriority",
	"sched_setparam",
	"sched_getparam",
	"sched_setscheduler",
	"sched_getscheduler",
	"sched_get_priority_max",
	"sched_get_priority_min",
	"sched_rr_get_interval",
	"mlock",
	"munlock",
	"mlockall",
	"munlockall",
	"vhangup",
	"modify_ldt",
	"pivot_root",
	"_sysctl",
	"prctl",
	"arch_prctl",
	"adjtimex",
	"setrlimit",
	"chroot",
	"sync",
	"acct",
	"settimeofday",
	"mount",
	"umount2",
	"swapon",
	"swapoff",
	"reboot",
	"sethostname",
	"setdomainname",
	"iopl",
	"ioperm",
	"create_module",
	"init_module",
	"delete_module",
	"get_kernel_syms",
	"query_module",
	"quotactl",
	"nfsservctl",
	"getpmsg",
	"putpmsg",
	"afs_syscall",
	"tuxcall",
	"security",
	"gettid",
	"readahead",
	"setxattr",
	"lsetxattr",
	"fsetxattr",
	"getxattr",
	"lgetxattr",
	"fgetxattr",
	"listxattr",
	"llistxattr",
	"flistxattr",
	"removexattr",
	"lremovexattr",
	"fremovexattr",
	"tkill",
	"time",
	"futex",
	"sched_setaffinity",
	"sched_getaffinity",
	"set_thread_area",
	"io_setup",
	"io_destroy",
	"io_getevents",
	"io_submit",
	"io_cancel",
	"get_thread_area",
	"lookup_dcookie",
	"epoll_create",
	"epoll_ctl_old",
	"epoll_wait_old",
	"remap_file_pages",
	"getdents64",
	"set_tid_address",
	"restart_syscall",
	"semtimedop",
	"fadvise64",
	"timer_create",
	"timer_settime",
	"timer_gettime",
	"timer_getoverrun",
	"timer_delete",
	"clock_settime",
	"clock_gettime",
	"clock_getres",
	"clock_nanosleep",
	"exit_group",
	"epoll_wait",
	"epoll_ctl",
	"tgkill",
	"utimes",
	"vserver",
	"mbind",
	"set_mempolicy",
	"get_mempolicy",
	"mq_open",
	"mq_unlink",
	"mq_timedsend",
	"mq_timedreceive",
	"mq_notify",
	"mq_getsetattr",
	"kexec_load",
	"waitid",
	"add_key",
	"request_key",
	"keyctl",
	"ioprio_set",
	"ioprio_get",
	"inotify_init",
	"inotify_add_watch",
	"inotify_rm_watch",
	"migrate_pages",
	"openat",
	"mkdirat",
	"mknodat",
	"fchownat",
	"futimesat",
	"newfstatat",
	"unlinkat",
	"renameat",
	"linkat",
	"symlinkat",
	"readlinkat",
	"fchmodat",
	"faccessat",
	"pselect6",
	"ppoll",
	"unshare",
	"set_robust_list",
	"get_robust_list",
	"splice",
	"tee",
	"sync_file_range",
	"vmsplice",
	"move_pages",
	"utimensat",
	"epoll_pwait",
	"signalfd",
	"timerfd_create",
	"eventfd",
	"fallocate",
	"timerfd_settime",
	"timerfd_gettime",
	"accept4",
	"signalfd4",
	"eventfd2",
	"epoll_create1",
	"dup3",
	"pipe2",
	"inotify_init1",
	"preadv",
	"pwritev",
	"rt_tgsigqueueinfo",
	"perf_event_open",
	"recvmmsg",
	"fanotify_init",
	"fanotify_mark",
	"prlimit64",
	"name_to_handle_at",
	"open_by_handle_at",
	"clock_adjtime",
	"syncfs",
	"sendmmsg",
	"setns",
	"getcpu",
	"process_vm_readv",
	"process_vm_writev",
	"kcmp",
	"finit_module",
	"sched_setattr",
	"sched_getattr",
	"renameat2",
	"seccomp",
	"getrandom",
	"memfd_create",
	"kexec_file_load",
	"bpf",
	"execveat",
}

func callNameX86Family64(num int16) string {
	if num > SyscallX86MaxNum64 {
		return SyscallX86UnknownName
	}

	return syscallNumTableX86Family64[num]
}

func callNumTableIsOkX86Family64() bool {
	if (len(syscallNumTableX86Family64) == SyscallX86MaxNum64+1) &&
		syscallNumTableX86Family64[SyscallX86MaxNum64] == SyscallX86LastName64 {
		return true
	}

	return false
}

var syscallNumTableX86Family32 = [...]string{
	"restart_syscall",
	"_exit",
	"fork",
	"read",
	"write",
	"open",
	"close",
	"waitpid",
	"creat",
	"link",
	"unlink",
	"execve",
	"chdir",
	"time",
	"mknod",
	"chmod",
	"lchown",
	"break",
	"oldstat",
	"lseek",
	"getpid",
	"mount",
	"umount",
	"setuid",
	"getuid",
	"stime",
	"ptrace",
	"alarm",
	"oldfstat",
	"pause",
	"utime",
	"stty",
	"gtty",
	"access",
	"nice",
	"ftime",
	"sync",
	"kill",
	"rename",
	"mkdir",
	"rmdir",
	"dup",
	"pipe",
	"times",
	"prof",
	"brk",
	"setgid",
	"getgid",
	"signal",
	"geteuid",
	"getegid",
	"acct",
	"umount2",
	"lock",
	"ioctl",
	"fcntl",
	"mpx",
	"setpgid",
	"ulimit",
	"oldolduname",
	"umask",
	"chroot",
	"ustat",
	"dup2",
	"getppid",
	"getpgrp",
	"setsid",
	"sigaction",
	"sgetmask",
	"ssetmask",
	"setreuid",
	"setregid",
	"sigsuspend",
	"sigpending",
	"sethostname",
	"setrlimit",
	"getrlimit",
	"getrusage",
	"gettimeofday",
	"settimeofday",
	"getgroups",
	"setgroups",
	"oldselect",
	"symlink",
	"oldlstat",
	"readlink",
	"uselib",
	"swapon",
	"reboot",
	"readdir",
	"old_mmap",
	"munmap",
	"truncate",
	"ftruncate",
	"fchmod",
	"fchown",
	"getpriority",
	"setpriority",
	"profil",
	"statfs",
	"fstatfs",
	"ioperm",
	"socketcall",
	"syslog",
	"setitimer",
	"getitimer",
	"stat",
	"lstat",
	"fstat",
	"olduname",
	"iopl",
	"vhangup",
	"idle",
	"vm86old",
	"wait4",
	"swapoff",
	"sysinfo",
	"ipc",
	"fsync",
	"sigreturn",
	"clone",
	"setdomainname",
	"uname",
	"modify_ldt",
	"adjtimex",
	"mprotect",
	"sigprocmask",
	"create_module",
	"init_module",
	"delete_module",
	"get_kernel_syms",
	"quotactl",
	"getpgid",
	"fchdir",
	"bdflush",
	"sysfs",
	"personality",
	"afs_syscall",
	"setfsuid",
	"setfsgid",
	"_llseek",
	"getdents",
	"select",
	"flock",
	"msync",
	"readv",
	"writev",
	"getsid",
	"fdatasync",
	"_sysctl",
	"mlock",
	"munlock",
	"mlockall",
	"munlockall",
	"sched_setparam",
	"sched_getparam",
	"sched_setscheduler",
	"sched_getscheduler",
	"sched_yield",
	"sched_get_priority_max",
	"sched_get_priority_min",
	"sched_rr_get_interval",
	"nanosleep",
	"mremap",
	"setresuid",
	"getresuid",
	"vm86",
	"query_module",
	"poll",
	"nfsservctl",
	"setresgid",
	"getresgid",
	"prctl",
	"rt_sigreturn",
	"rt_sigaction",
	"rt_sigprocmask",
	"rt_sigpending",
	"rt_sigtimedwait",
	"rt_sigqueueinfo",
	"rt_sigsuspend",
	"pread64",
	"pwrite64",
	"chown",
	"getcwd",
	"capget",
	"capset",
	"sigaltstack",
	"sendfile",
	"getpmsg",
	"putpmsg",
	"vfork",
	"ugetrlimit",
	"mmap2",
	"truncate64",
	"ftruncate64",
	"stat64",
	"lstat64",
	"fstat64",
	"lchown32",
	"getuid32",
	"getgid32",
	"geteuid32",
	"getegid32",
	"setreuid32",
	"setregid32",
	"getgroups32",
	"setgroups32",
	"fchown32",
	"setresuid32",
	"getresuid32",
	"setresgid32",
	"getresgid32",
	"chown32",
	"setuid32",
	"setgid32",
	"setfsuid32",
	"setfsgid32",
	"pivot_root",
	"mincore",
	"madvise",
	"getdents64",
	"fcntl64",
	"unused.222", //222 - unused
	"security",   //223 - supposed to be unused too
	"gettid",
	"readahead",
	"setxattr",
	"lsetxattr",
	"fsetxattr",
	"getxattr",
	"lgetxattr",
	"fgetxattr",
	"listxattr",
	"llistxattr",
	"flistxattr",
	"removexattr",
	"lremovexattr",
	"fremovexattr",
	"tkill",
	"sendfile64",
	"futex",
	"sched_setaffinity",
	"sched_getaffinity",
	"set_thread_area",
	"get_thread_area",
	"io_setup",
	"io_destroy",
	"io_getevents",
	"io_submit",
	"io_cancel",
	"fadvise64",
	"unused.251", //251 - unused (available for reuse) was used by sys_set_zone_reclaim
	"exit_group",
	"lookup_dcookie",
	"epoll_create",
	"epoll_ctl",
	"epoll_wait",
	"remap_file_pages",
	"set_tid_address",
	"timer_create",
	"timer_settime",
	"timer_gettime",
	"timer_getoverrun",
	"timer_delete",
	"clock_settime",
	"clock_gettime",
	"clock_getres",
	"clock_nanosleep",
	"statfs64",
	"fstatfs64",
	"tgkill",
	"utimes",
	"fadvise64_64",
	"vserver",
	"mbind",
	"get_mempolicy",
	"set_mempolicy",
	"mq_open",
	"mq_unlink",
	"mq_timedsend",
	"mq_timedreceive",
	"mq_notify",
	"mq_getsetattr",
	"kexec_load",
	"waitid",
	"setaltroot", //285 - sys_setaltroot (sort of used :))
	"add_key",
	"request_key",
	"keyctl",
	"ioprio_set",
	"ioprio_get",
	"inotify_init",
	"inotify_add_watch",
	"inotify_rm_watch",
	"migrate_pages",
	"openat",
	"mkdirat",
	"mknodat",
	"fchownat",
	"futimesat",
	"fstatat64",
	"unlinkat",
	"renameat",
	"linkat",
	"symlinkat",
	"readlinkat",
	"fchmodat",
	"faccessat",
	"pselect6",
	"ppoll",
	"unshare",
	"set_robust_list",
	"get_robust_list",
	"splice",
	"sync_file_range",
	"tee",
	"vmsplice",
	"move_pages",
	"getcpu",
	"epoll_pwait",
	"utimensat",
	"signalfd",
	"timerfd_create",
	"eventfd",
	"fallocate",
	"timerfd_settime",
	"timerfd_gettime",
	"signalfd4",
	"eventfd2",
	"epoll_create1",
	"dup3",
	"pipe2",
	"inotify_init1",
	"preadv",
	"pwritev",
	"rt_tgsigqueueinfo",
	"perf_event_open",
	"recvmmsg",
	"fanotify_init",
	"fanotify_mark",
	"prlimit64",
	"name_to_handle_at",
	"open_by_handle_at",
	"clock_adjtime",
	"syncfs",
	"sendmmsg",
	"setns",
	"process_vm_readv",
	"process_vm_writev",
	"kcmp",
	"finit_module",
	"sched_setattr",
	"sched_getattr",
	"renameat2",
	"seccomp",
	"getrandom",
	"memfd_create",
	"bpf",
	"execveat",
}

func callNameX86Family32(num int16) string {
	if num > SyscallX86MaxNum32 {
		return SyscallX86UnknownName
	}

	return syscallNumTableX86Family32[num]
}

func callNumTableIsOkX86Family32() bool {
	if (len(syscallNumTableX86Family32) == SyscallX86MaxNum32+1) &&
		syscallNumTableX86Family32[SyscallX86MaxNum32] == SyscallX86LastName32 {
		return true
	}

	return false
}
