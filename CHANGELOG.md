# Changelog

## [1.0.0](https://github.com/f1zm0/hades/compare/v0.0.1...v1.0.0) (2023-04-25)


### ⚠ BREAKING CHANGES

* include direct syscall logic only with build constraint
* integrate acheron for indirect syscalls

### Features

* add frida handler scripts for testing ([95f2c25](https://github.com/f1zm0/hades/commit/95f2c25b975d08669258c7e9d01fe3f0a334a8d5))
* add generics helpers for nt calls args ([0ef8b7b](https://github.com/f1zm0/hades/commit/0ef8b7bd34312c4892da8704c63b7cb9386c85a8))
* include direct syscall logic only with build constraint ([0ec7c21](https://github.com/f1zm0/hades/commit/0ec7c214a9c718d231b051bad62b15902ccee77c))
* integrate acheron for indirect syscalls ([ffda94d](https://github.com/f1zm0/hades/commit/ffda94de753a6aff5da2845d87852d49785bbe36))


### Continuous Integration

* add github action for autorelease and golangci file ([30b8740](https://github.com/f1zm0/hades/commit/30b87409820381a8d51f401b10001d78c838b071))


### Misc

* bump go version ([11f5cce](https://github.com/f1zm0/hades/commit/11f5cce45ca3df4a74900ae37f1ccfba0e79338f))
* change dir structure ([de4e131](https://github.com/f1zm0/hades/commit/de4e13134581c7d38afc5344ea37786b431e7ed4))
* make sure cli flags are lowercase ([0466fd0](https://github.com/f1zm0/hades/commit/0466fd0c876542e8cbacc7418c22ca57b37cd49e))
* remove unused types and functions ([5b80e14](https://github.com/f1zm0/hades/commit/5b80e141f1c4875daabcd9919e8ac091f4ccec5b))
* rename direct syscall stub func for clarity ([27c53aa](https://github.com/f1zm0/hades/commit/27c53aaf7770fef13c7dffb519b253fb031d022d))
* replace acheron dependency with public version ([c262469](https://github.com/f1zm0/hades/commit/c2624698b2c9cc25bd474190709731a80ff8b948))
* replace hash function in hasher main ([7a3e2d4](https://github.com/f1zm0/hades/commit/7a3e2d4ce93290f57b5f3fb4a05dec1df61e6736))
* update cli banner ([01e1024](https://github.com/f1zm0/hades/commit/01e102461b45cdca10b4da6923164cbf404f58df))
* update description and install cmd in readme ([840451d](https://github.com/f1zm0/hades/commit/840451dd78bba41dd1c4a333c4199dbe32986e3b))
* update go mod and sum ([008c153](https://github.com/f1zm0/hades/commit/008c153e81108ff4049fc36a4db656eeb37e85b4))
* update readme ([8698a57](https://github.com/f1zm0/hades/commit/8698a57b4f13ee504288ea9a5c1bd4e5c40fd339))


### Code Refactoring

* change receiving vars according to acheron api changes ([837d624](https://github.com/f1zm0/hades/commit/837d62404d51df9b2da3f69691cca3b7546be10e))
* move inj technique selection to load method ([113af97](https://github.com/f1zm0/hades/commit/113af97f40f43ff2662726ac54956ef103994459))
* replace djb2 with xored hash func ([7c1bca0](https://github.com/f1zm0/hades/commit/7c1bca085e6fdcf66bbc1d579bbc42a674a7a3c6))
* split loader and resolver into separate pkgs ([4d444af](https://github.com/f1zm0/hades/commit/4d444af72a8cc53320d2af500fcf346d319fb1ad))


### Documentation

* correct installation commands in readme ([84970dc](https://github.com/f1zm0/hades/commit/84970dc4e5faefd3e9b28b21b5815d54d710b371))
* fix handles in credits ([03150b4](https://github.com/f1zm0/hades/commit/03150b46f025f5bb0d7ab2441ee6d8ae94c68005))
* fix image links and add notes ([e0a7839](https://github.com/f1zm0/hades/commit/e0a783996de8667f2ecfe91a4528197d3e1ee918))
* refresh license badge ([8350240](https://github.com/f1zm0/hades/commit/83502407df11e9227187a724198548f4b1ed001b))
* udpate main readme ([5bbe077](https://github.com/f1zm0/hades/commit/5bbe077a66322c66a59816bd6bf29b3bf94949f4))
* update credits and project desc in main readme ([f4bba95](https://github.com/f1zm0/hades/commit/f4bba950dd0d4303d34dd8a8465d290869b77a0a))
* update main readme ([6e85d68](https://github.com/f1zm0/hades/commit/6e85d680d41d9021a63049907ccdf81ce4c8184d))
