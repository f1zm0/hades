package resolver

// Resolver is the interface that wraps the basic Resolve method.
type Resolver interface {
	// GetModuleHandleByHash returns a PEModule struct containing the base address and a pointer to the PE file of the module.
	GetModuleHandleByHash(modNameHash int64) (*PEModule, error)

	// GetProcAddressByHash returns the base address of a function.
	GetProcAddressByHash(p *PEModule, funcNameHash int64) (int64, error)

	// GetSyscallSSN returns the syscall SSN.
	GetSyscallSSN(funcNameHash int64) int
}

func NewResolver() (Resolver, error) {
	return newSSNResolver()
}
