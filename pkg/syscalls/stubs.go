//go:build direct_syscalls

package syscalls

func execDirectSyscall(ssn uint16, argh ...uintptr) (errcode uint32)

func DirectSyscall(ssn uint16, argh ...uintptr) uint32 {
	return execDirectSyscall(ssn, argh...)
}
