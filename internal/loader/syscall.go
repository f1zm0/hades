package loader

import "errors"

func execSyscall(callid uint16, argh ...uintptr) (errcode uint32)

func Syscall(syscallID uint16, args ...uintptr) (errcode uint32, err error) {
	errcode = execSyscall(syscallID, args...)

	if errcode != 0 {
		return errcode, errors.New("non-zero return from syscall")
	}
	return errcode, nil
}
