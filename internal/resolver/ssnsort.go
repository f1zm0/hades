package resolver

import (
	"errors"
	"sort"
	"strings"
	"unsafe"

	"github.com/f1zm0/hades/pkg/hashing"
	rrd "github.com/f1zm0/hades/pkg/rawreader"

	"github.com/Binject/debug/pe"
)

type ssnSortResolver struct {
	// Hashing provider
	djb2 *hashing.DJB2

	// Map of InMemProc structs indexed by their djb2 hash
	ntdllApi map[int64]InMemProc
}

var _ Resolver = (*ssnSortResolver)(nil)

func newSSNResolver() (Resolver, error) {
	r := &ssnSortResolver{
		djb2: hashing.NewDJB2(),
	}
	m, err := r.resolveSyscallIDs()
	if err != nil {
		return nil, err
	}
	r.ntdllApi = m
	return r, nil
}

// GetModuleHandleByHash returns a PEModule struct containing the base address and a pointer to the PE file of the module.
func (r *ssnSortResolver) GetModuleHandleByHash(modNameHash int64) (*PEModule, error) {
	entries := GetLdrTableEntries()
	for _, entry := range entries {
		if r.djb2.HashString(entry.BaseDllName.String()) == modNameHash {

			modBaseAddr := uintptr(unsafe.Pointer(entry.DllBase))
			modSize := int(uintptr(unsafe.Pointer(entry.SizeOfImage)))
			rr := rrd.NewRawReader(modBaseAddr, modSize)

			p, err := pe.NewFileFromMemory(rr)
			if err != nil {
				return nil, errors.New("Error reading module from memory")
			}

			pm := &PEModule{
				BaseAddr: modBaseAddr,
				File:     p,
			}
			return pm, nil
		}
	}
	return nil, errors.New("Module not found. Probably not loaded.")
}

// GetProcAddressByHash returns the base address of a function.
func (r *ssnSortResolver) GetProcAddressByHash(p *PEModule, funcNameHash int64) (int64, error) {
	ex, err := p.File.Exports()
	if err != nil {
		return 0, err
	}
	for _, exp := range ex {
		if r.djb2.HashString(exp.Name) == funcNameHash {
			return (int64(p.BaseAddr) + int64(exp.VirtualAddress)), nil
		}
	}

	return 0, errors.New("Function not found")
}

func getAllProcs(p *PEModule) ([]InMemProc, error) {
	var procs []InMemProc

	ex, err := p.File.Exports()
	if err != nil {
		return procs, err
	}
	for _, exp := range ex {
		memAddr := int64(p.BaseAddr) + int64(exp.VirtualAddress)
		procs = append(procs, InMemProc{
			Name:     exp.Name,
			BaseAddr: uintptr(memAddr),
		})
	}

	return procs, nil
}

func (r *ssnSortResolver) resolveSyscallIDs() (map[int64]InMemProc, error) {
	procMap := make(map[int64]InMemProc)
	var ntProcs []InMemProc

	hNtdll, err := r.GetModuleHandleByHash(249899979757565421)
	if err != nil {
		return procMap, err
	}
	procs, err := getAllProcs(hNtdll)
	if err != nil {
		return procMap, err
	}

	for _, p := range procs {
		if strings.HasPrefix(p.Name, "Zw") {
			ntProcs = append(ntProcs, p)
		}
	}

	sort.Slice(ntProcs, func(i, j int) bool {
		return ntProcs[i].BaseAddr < ntProcs[j].BaseAddr
	})

	for i := range ntProcs {
		ntProcs[i].SSN = i
		ntProcs[i].Name = "Nt" + ntProcs[i].Name[2:]
		procMap[r.djb2.HashString(ntProcs[i].Name)] = ntProcs[i]
	}

	return procMap, nil
}

// GetSyscallSSN returns the syscall ID of a native API function by its djb2 hash.
// If the function is not found, -1 is returned.
func (r *ssnSortResolver) GetSyscallSSN(fnHash int64) int {
	if v, ok := r.ntdllApi[fnHash]; ok {
		return v.SSN
	}
	return -1
}
