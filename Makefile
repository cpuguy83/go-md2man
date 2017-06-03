NAME=go-md2man
BINARYNAME=md2man
OWNER=cpuguy83
ARCH=$(shell uname -m)
VERSION=0.0.0

build: build-darwin build-linux

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
	gometalinter --config .gometalinter.json ./...

circleci:
	rm ~/.gitconfig
	rm -rf /home/ubuntu/.go_workspace/src/github.com/$(OWNER)/$(NAME) && cd .. \
		&& mkdir -p /home/ubuntu/.go_workspace/src/github.com/$(OWNER) \
		&& mv $(NAME) /home/ubuntu/.go_workspace/src/github.com/$(OWNER)/$(NAME) \
		&& ln -s /home/ubuntu/.go_workspace/src/github.com/$(OWNER)/$(NAME) $(NAME)

clean:
	rm -rf build

deps:
	dep ensure -vendor-only

test: src
	true

.PHONY: build build-darwin build-linux check clean deps test
