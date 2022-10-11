package types

import "golang.org/x/sys/windows"

type UnicodeString struct {
	Length        uint16
	MaximumLength uint16
	Buffer        *uint16
}

func (s UnicodeString) String() string {
	return windows.UTF16PtrToString(s.Buffer)
}
