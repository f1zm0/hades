{
  onEnter(log, args) {
	log("NtWriteVirtualMemory()");
        this.Handle = args[0];
        this.BaseAddress = args[1];
        this.Buffer = args[2];
        this.NumberOfBytesToWrite = args[3];
        this.NumberOfBytesWritten = args[4];
  },
  onLeave(log) {
	if(!(this.Handle == 0xffffffff)){
            log(" |-- Handle: " + this.Handle);
            log(" |-- BaseAddress: " + this.BaseAddress);
            log(" |-- Buffer: " + this.Buffer);
            log(" |-- NumberOfBytesToWrite: " + this.NumberOfBytesToWrite);
            log(" |-- NumberOfBytesWritten: " + this.NumberOfBytesWritten);
            log("")
        }
  }
}
