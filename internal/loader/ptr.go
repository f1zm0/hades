package loader

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	nullptr = uintptr(0)
)

type ToPtr interface {
	int | uint | uint32 | uint16 | uint8 | uintptr
}

type WinTypesPtr interface {
	windows.Handle |
		windows.SecurityAttributes |
		windows.StartupInfo |
		windows.ProcessInformation
}

func ptr[T ToPtr](v T) uintptr {
	return uintptr(v)
}

func ptr2ptr[T ToPtr | WinTypesPtr](v *T) uintptr {
	if v == nil {
		return nullptr
	}
	return uintptr(unsafe.Pointer(v))
}
