package types

import (
	"fmt"
	"unsafe"
)

type LdrDataTableEntry struct {
	InLoadOrderLinks           ListEntry
	InMemoryOrderLinks         ListEntry
	InInitializationOrderLinks ListEntry
	DllBase                    *uintptr
	EntryPoint                 *uintptr
	SizeOfImage                *uintptr
	FullDllName                UnicodeString
	BaseDllName                UnicodeString
	Flags                      uint32
	LoadCount                  uint16
	TlsIndex                   uint16
	HashLinks                  ListEntry
	TimeDateStamp              uint64
}

type ListEntry struct {
	Flink *ListEntry
	Blink *ListEntry
}

func (te *LdrDataTableEntry) DumpInfo() {
	fmt.Printf(`
----
Name: %s
Base: 0x%016x
Size: %d
`,
		te.FullDllName.String(),
		uintptr(unsafe.Pointer(te.DllBase)),
		int(uintptr(unsafe.Pointer(te.SizeOfImage))),
	)
}
