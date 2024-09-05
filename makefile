expect: fmt build

.PHONY: build
build:
	go build .

.PHONY: fmt
fmt:
	go fmt .

