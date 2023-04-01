package resolver

import (
	"github.com/Binject/debug/pe"
)

// PEModule is a struct that contains the base address of a PE module and a pointer to the PE file.
type PEModule struct {
	BaseAddr uintptr
	File     *pe.File
}

// InMemProc is a struct that contains the name, base address and SSN of a function.
type InMemProc struct {
	Name     string
	BaseAddr uintptr
	SSN      int
}
