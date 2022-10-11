package loader

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"sort"
	"strings"
	"unsafe"

	"github.com/f1zm0/hades/pkg/hashing"
	rrd "github.com/f1zm0/hades/pkg/rawreader"

	"github.com/Binject/debug/pe"
)

type PEModule struct {
	BaseAddr uintptr
	File     *pe.File
}

type InMemProc struct {
	Name      string
	BaseAddr  uintptr
	SyscallID int
}

type Loader struct {
	// Hashing provider
	djb2 *hashing.DJB2

	// Map of ntdll.dll [funcNameHash]syscallID
	ntdllApi map[int64]InMemProc
}

func NewLoader() *Loader {
	djb2 := hashing.NewDJB2()
	pl := &Loader{djb2: djb2}
	ntProcs, err := pl.ResolveSyscallIDs()
	if err != nil {
		panic(err)
	}

	pl.ntdllApi = ntProcs

	return pl
}

func (pl *Loader) GetSysID(funcName string) int {
	fHash := pl.djb2.HashString(funcName)
	if v, ok := pl.ntdllApi[fHash]; ok {
		return v.SyscallID
	}
	return -1
}

func (pl *Loader) GetModuleHandleByHash(modNameHash int64) (*PEModule, error) {
	entries := GetLdrTableEntries()
	for _, entry := range entries {
		if pl.djb2.HashString(entry.BaseDllName.String()) == modNameHash {

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

func (pl *Loader) GetProcAddressByHash(p *PEModule, funcNameHash int64) (int64, error) {
	ex, err := p.File.Exports()
	if err != nil {
		return 0, err
	}
	for _, exp := range ex {
		if pl.djb2.HashString(exp.Name) == funcNameHash {
			return (int64(p.BaseAddr) + int64(exp.VirtualAddress)), nil
		}
	}

	return 0, errors.New("Function not found")
}

func GetAllProcs(p *PEModule) ([]InMemProc, error) {
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

func (p *InMemProc) IsHooked() bool {
	safeBytes := []byte{0x4c, 0x8b, 0xd1, 0xb8}
	stub := make([]byte, len(safeBytes))

	rr := rrd.NewRawReader(p.BaseAddr, len(safeBytes))

	sr := io.NewSectionReader(rr, 0, 1<<63-1)
	binary.Read(sr, binary.LittleEndian, &stub)

	if bytes.Compare(stub, safeBytes) == 0 {
		return true
	}
	return false
}

func (pl *Loader) ResolveSyscallIDs() (map[int64]InMemProc, error) {
	procMap := make(map[int64]InMemProc)
	var ntProcs []InMemProc

	hNtdll, err := pl.GetModuleHandleByHash(249899979757565421)
	if err != nil {
		return procMap, err
	}
	procs, err := GetAllProcs(hNtdll)
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
		ntProcs[i].SyscallID = i
		ntProcs[i].Name = "Nt" + ntProcs[i].Name[2:]
		procMap[pl.djb2.HashString(ntProcs[i].Name)] = ntProcs[i]
	}

	return procMap, nil
}
