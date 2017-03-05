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

test: src
	true

install: build
	install build/$(shell uname -s)/$(BINARYNAME) /usr/local/bin

deps:
	go get -u github.com/progrium/gh-release/...

deps-update:
	dep ensure -vendor-only

release:
	rm -rf release && mkdir release
	tar -zcf release/$(NAME)_$(VERSION)_Linux_$(ARCH).tgz -C build/Linux $(BINARYNAME)
	tar -zcf release/$(NAME)_$(VERSION)_Darwin_$(ARCH).tgz -C build/Darwin $(BINARYNAME)
	gh-release checksums sha256
	gh-release create $(OWNER)/$(NAME) $(VERSION) $(shell git rev-parse --abbrev-ref HEAD) v$(VERSION)

circleci:
	rm ~/.gitconfig
	rm -rf /home/ubuntu/.go_workspace/src/github.com/$(OWNER)/$(NAME) && cd .. \
		&& mkdir -p /home/ubuntu/.go_workspace/src/github.com/$(OWNER) \
		&& mv $(NAME) /home/ubuntu/.go_workspace/src/github.com/$(OWNER)/$(NAME) \
		&& ln -s /home/ubuntu/.go_workspace/src/github.com/$(OWNER)/$(NAME) $(NAME)

clean:
	rm -rf build release

check:
	gometalinter --config .gometalinter.json ./...

.PHONY: build build-darwin build-linux check deps deps-update release

