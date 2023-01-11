PROJ_NAME=hades
PROJ_MOD_PREFIX=github.com/f1zm0/hades
BUILD_PATH=${CURDIR}/dist
ENTRYPOINT=${CURDIR}/cmd/hades/main.go
GCFLAGS=-gcflags=all=-trimpath=$(GOPATH)
ASMFLAGS=-asmflags=all=-trimpath=$(GOPATH)
# LDFLAGS="-s -w -H=windowsgui"
LDFLAGS="-s -w"


.PHONY: default
default: build

.PHONY: help
## help: prints an help message with all the available targets
help:
	@echo -e "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'


.PHONY: clean
## clean: delete all binaries
clean:
	@if [ -d "${BUILD_PATH}" ]; then ${RM} ${BUILD_PATH}/* ; fi


.PHONY: build
## build: builds binary for Windows
build:
	@echo "Building binaries for Windows x64 ..."; \
		GOOS=windows GOARCH=amd64 go build -ldflags=${LDFLAGS} ${GCFLAGS} ${ASMFLAGS} \
			-o ${BUILD_PATH}/${PROJ_NAME}.exe ${ENTRYPOINT} || exit 1; \
