NAME=go-md2man
BINARYNAME=md2man
OWNER=cpuguy83
ARCH=$(shell uname -m)
VERSION=1.0.6

build:
	go build -a \
		-installsuffix cgo \
		-ldflags "-X main.Version=$(VERSION)"

build-cross: build-darwin build-linux

build-linux:
	mkdir -p build/Linux && GOOS=linux CGO_ENABLED=0 go build -a \
		-installsuffix cgo \
		-ldflags "-X main.Version=$(VERSION)" \
		-o build/Linux/$(BINARYNAME)

build-darwin:
	mkdir -p build/Darwin && GOOS=darwin CGO_ENABLED=0 go build -a \
		-installsuffix cgo \
		-ldflags "-X main.Version=$(VERSION)" \
		-o build/Darwin/$(BINARYNAME)

check:
	script/validate/vendor
	gometalinter --config .gometalinter.json ./...

clean:
	rm -rf build

deps:
	dep ensure -vendor-only

deps-dev:
	script/setup/dev-tools

test:
	go test -v ./...

.PHONY: build build-cross build-darwin build-linux check clean deps deps-dev test
