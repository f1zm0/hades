# --- Go Variables
GO=$(shell which go)
GOINSTALL=${GO} install
GOGET=${GO} get
GOBUILD=${GO} build
GOTEST=${GO} test

# --- Config Variables
WIN_ARCHS=amd64 # 386 (not supported yet)
COMMIT_ID=$(shell git rev-parse --short HEAD)
TODAY=$(shell date +%d/%m/%y)

ifdef VERSION
	VERSION := $(VERSION)
else
	VERSION := dev
endif

# --- Project Vars
PROJ_NAME=hades
PROJ_MOD_PREFIX=github.com/f1zm0/hades
BUILD_PATH=${CURDIR}/dist
ENTRYPOINT=${CURDIR}/cmd/hades/main.go

# --- Compiler Vars
GCFLAGS=-gcflags=all=-trimpath=$(GOPATH)
ASMFLAGS=-asmflags=all=-trimpath=$(GOPATH)
# LDFLAGS="-s -w -H=windowsgui"
LDFLAGS="-s -w"


# --- Targets
.PHONY: default
default: build


.PHONY: help
## help: prints an help message with all the available targets
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'


.PHONY: clean
## clean: delete all binaries
clean:
	@if [ -d "${BUILD_PATH}" ]; then ${RM} ${BUILD_PATH}/* ; fi


.PHONY: test
## test: test code base using go test
test:
	${GOTEST} ./... -v -cover
 

.PHONY: build
## build: builds binary for Windows
build:
	@for ARCH in ${WIN_ARCHS}; do \
		echo "Building binaries for Windows $${ARCH} ..."; \
		GOOS=windows GOARCH=$${ARCH} ${GOBUILD} -ldflags=${LDFLAGS} ${GCFLAGS} ${ASMFLAGS} \
			-o ${BUILD_PATH}/${PROJ_NAME}-win-$${ARCH}-${VERSION}.exe ${ENTRYPOINT} || exit 1;\
	done;
