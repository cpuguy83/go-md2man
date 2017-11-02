.PHONY: help build vendor clean

help:
	@echo "Targets:"
	@echo "  - build:  Build the go-md2man binary"
	@echo "  - vendor: Installs govend and updates vendor libs required for building"
	@echo "  - clean:  Remove build binary and data in the vendor folder"
	@echo "  - help:   You're reading it!"

build:
	CGO_ENABLED=0 go build

vendor:
	go get -u github.com/govend/govend
	govend -v -u --prune

clean:
	rm -f go-md2man
	rm -rf vendor/*
