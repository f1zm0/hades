package loader

import (
	"golang.org/x/sys/windows"
)

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
