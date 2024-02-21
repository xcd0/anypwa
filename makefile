BIN           := ./anypwa.exe
REVISION      := `git rev-parse --short HEAD`
FLAGS_UNIX    := -ldflags="-X main.version=$(VERSION) -X main.revision=$(REVISION) -s -w -extldflags='-static' -buildid=" -a -tags=netgo -trimpath
FLAGS_WIN     := -ldflags="-H windowsgui -X main.version=$(VERSION) -X main.revision=$(REVISION) -s -w -extldflags='-static' -buildid=" -a -tags=netgo -trimpath

#FLAG          := $(FLAGS_UNIX)
FLAG          := $(FLAGS_WIN)

all:
	cat ./makefile

build:
	rm -rf ./files
	make generate
	make fmt
	go build

release:
	rm -rf ./files
	make generate
	make fmt
	go build $(FLAG)
	make upx 
	@echo Success!

fmt:
	goimports -w *.go
	gofmt -w *.go

generate:
	go generate

upx:
	upx --lzma $(BIN)

