<p align="center">
    <img src="static/hades-banner.png" title="hades banner" width="80%"/>
</p>
<p align="center">
  <a href="https://github.com/f1zm0/hades/releases">
    <img alt="Made with Go" src="https://img.shields.io/badge/Made%20with%20Go-00ADD8?style=for-the-badge&logo=Go&logoColor=white" style="max-width: 100%;">
</a>
<a href="https://github.com/f1zm0/hades">
    <img src="https://img.shields.io/github/license/f1zm0/hades?style=for-the-badge&color=aabbcc&logo=bookstack&logoColor=white&labelColor=2b2c34" alt="project license">
</a>
<a href="https://github.com/f1zm0/hades/issues">
    <img alt="Issues" src="https://img.shields.io/github/issues/f1zm0/hades?style=for-the-badge&logo=dependabot&color=aabbcc&logoColor=d9e0ee&labelColor=2b2c33" style="max-width: 100%;">
  </a>
<a href="#"> <img src="https://img.shields.io/badge/Status-PoC-aabbcc?style=for-the-badge&labelColor=2b2c33&logo=curl" alt="project status"> </a>
</p>

<p align="center">
  <i>SSN sorting and direct syscall invocation for AV/EDR evasion in Go and Go ASM</i>
</p>

## Disclaimer

The techniques used in this project are not new. This project is merely a proof of concept, and has been created for educational purposes only, to experiment with malware dev in Go, and learn more about the [unsafe](https://pkg.go.dev/unsafe) package and the weird [Go Assembly](https://go.dev/doc/asm) syntax.

Also, this project is not intended to be used to bypass any particular EDR or anti malware solution.

## Credits

Big thanks to the following people that shared their knowledge and code that inspired this tool:

- [@smelly\_\_vx](https://twitter.com/@RtlMateusz) and [@am0nsec](https://twitter.com/am0nsec) creators of [Hell's Gate](https://github.com/am0nsec/HellsGate)
- [@modexp](https://twitter.com/modexpblog)'s blog post on [Bypassing User-Mode Hooks and syscall invocation in C](https://www.mdsec.co.uk/2020/12/bypassing-user-mode-hooks-and-direct-invocation-of-system-calls-for-red-teams/)
- [@ElephantSe4l](ElephantSe4l) creator of [FreshyCalls](https://github.com/crummie5/FreshyCalls)
- [@thefLink](https://twitter.com/theflink_) creator of [RecycledGate](https://github.com/thefLink/RecycledGate)
- [@C_Sto](https://twitter.com/C_Sto) creator of [BananaPhone](https://github.com/C-Sto/BananaPhone)

## License

This project is licensed under the GPLv3 License - see the [LICENSE](LICENSE) file for details
