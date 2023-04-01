package resolver

import (
	"unsafe"

	wt "github.com/f1zm0/hades/internal/types"
)

func GetLdrTableEntryPtr(listptr uintptr, i int64) *wt.LdrDataTableEntry

func GetInMemoryOrderModuleListPtr() uintptr

func GetLdrTableEntries() []*wt.LdrDataTableEntry {
	entries := []*wt.LdrDataTableEntry{}
	var (
		entry      *wt.LdrDataTableEntry
		firstEntry *wt.LdrDataTableEntry
	)

	// addr of Ldr->InMemoryOrderModuleList
	modListPtr := GetInMemoryOrderModuleListPtr()

	firstEntry = GetLdrTableEntryPtr(modListPtr, 0)
	entries = append(entries, firstEntry)

	i := int64(1)
	for {
		entry = GetLdrTableEntryPtr(modListPtr, i)
		if entry == firstEntry || unsafe.Pointer(entry.DllBase) == unsafe.Pointer(nil) {
			break
		}
		entries = append(entries, entry)
		i = i + 1
	}

	return entries
}
