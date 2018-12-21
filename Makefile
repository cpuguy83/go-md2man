.PHONY: build
build:
	go build -mod=vendor

.PHONY: check
check:
	gometalinter --config .gometalinter.json ./...

.PHONY: mod
mod:
	go mod tidy -v
	go mod verify

.PHONY: vendor
vendor:
	go mod vendor
