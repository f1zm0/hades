<p align="center">
    <img src="static/hades-banner.png" title="hades banner" width="80%"/>
</p>
<p align="center">
  <a href="https://github.com/f1zm0/hades/releases">
    <img alt="Made with Go" src="https://img.shields.io/badge/Made%20with%20Go-00ADD8?style=for-the-badge&logo=Go&logoColor=white" style="max-width: 100%;">
</a>
<a href="https://github.com/f1zm0/hades">
    <img src="https://img.shields.io/github/license/f1zm0/hades?style=for-the-badge&color=aabbcc&logo=bookstack&logoColor=white&labelColor=2b2c33" alt="project license">
</a>
<a href="https://github.com/f1zm0/hades/issues">
    <img alt="Issues" src="https://img.shields.io/github/issues/f1zm0/hades?style=for-the-badge&logo=dependabot&color=aabbcc&logoColor=d9e0ee&labelColor=2b2c33" style="max-width: 100%;">
  </a>
    <a href="https://twitter.com/intent/follow?screen_name=f1zm0">
      <img alt="follow on Twitter" src="https://img.shields.io/twitter/follow/f1zm0?style=for-the-badge&logo=twitter&color=8aadf3&logoColor=d9e0ee&labelColor=2b2c33" />
    </a>
</p>

## About

`hades` is a proof of concept loader that combines SSN sorting and direct syscall invocation to bypass user-mode hooks in Go and Go-ASM. Needed functions are resolved by walking the PEB and using their djb2 hash, without calling other native APIs.

> **Note** <br/>
> The techniques used in this project are not new. This project has been created for educational purposes only, to experiment with malware dev in Go, and learn more about the [unsafe](https://pkg.go.dev/unsafe) package and the weird [Go Assembly](https://go.dev/doc/asm) syntax.
> Don't use it to on systems you don't own. The developer of this project is not responsible for any damage caused by this tool.

## Usage

The easiest way, is probably building the project on Linux using `make`.

```sh
git clone https://github.com/f1zm0/hades && cd hades
make
```

Then you can bring the executable to a x64 Windows host and run it with `./hades` or `./hades -h` to see the available options.

```
PS > .\hades.exe -h

  '||'  '||'     |     '||''|.   '||''''|   .|'''.|
   ||    ||     |||     ||   ||   ||  .     ||..  '
   ||''''||    |  ||    ||    ||  ||''|      ''|||.
   ||    ||   .''''|.   ||    ||  ||       .     '||
  .||.  .||. .|.  .||. .||...|'  .||.....| |'....|'

          version: dev [11/01/23] :: @f1zm0

Usage:
  hades -f <filepath> [-t selfthread|remotethread|queueuserapc]

Options:
  -f, --file <str>        shellcode file path (.bin)
  -t, --technique <str>   injection technique [selfthread, remotethread, queueuserapc]
```

### Example:

Inject shellcode that spawms `calc.exe` with [queueuserapc](https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-queueuserapc) technique:

```
.\hades.exe -f calc.bin -t queueuserapc
```

## Showcase

Below is a very quick proof of concept of the tools, that is used to inject a simple calc shellcode with APC injection, while intercepting the call to `NtQueueApcThread` with [Frida](https://frida.re). The loader doesn't care about the hook and instead uses the RVAs of `Zw*` functions to calculate the SSN of `NtQueueApcThread` and make a direct system call.

![NtQueueApcThread Frida interceptor](static/frida-poc.gif)

## Credits

Big thanks to the following people that shared their knowledge and code that inspired this tool:

- [@smelly\_\_vx](https://twitter.com/smelly_vx) and [@am0nsec](https://twitter.com/am0nsec) creators of [Hell's Gate](https://github.com/am0nsec/HellsGate)
- [@modexp](https://twitter.com/modexpblog)'s excellent blog post [Bypassing User-Mode Hooks and syscall invocation in C](https://www.mdsec.co.uk/2020/12/bypassing-user-mode-hooks-and-direct-invocation-of-system-calls-for-red-teams/)
- [@ElephantSe4l](https://twitter.com/elephantse4l) creator of [FreshyCalls](https://github.com/crummie5/FreshyCalls)
- [@C_Sto](https://twitter.com/c__sto) creator of [BananaPhone](https://github.com/C-Sto/BananaPhone)

## License

This project is licensed under the GPLv3 License - see the [LICENSE](LICENSE) file for details
