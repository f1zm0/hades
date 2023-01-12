{
  onEnter(log, args) {
    log('NtQueueApcThread()');
    this.ThreadHandle = args[0].toInt32();
    this.ApcRoutine = args[1].toInt32();
    this.ApcArgument = args[2].toInt32();
    this.ApcEnvironment = args[3].toInt32();
  },
  onLeave(log) {
    log(" |-- ThreadHandle: " + ThreadHandle);
    log(" |-- ApcRoutine: " + ApcRoutine);
    log(" |-- ApcArgument: " + ApcArgument);
    log(" |-- ApcEnvironment: " + ApcEnvironment);
    log("")
  }
}
