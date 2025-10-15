include Makefile.def.mk

.PHONY: tidy
tidy:
	export GOPROXY=https://goproxy.cn,direct
	go mod tidy

.PHONY: lint
lint: fmt vet golangci-lint

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: golangci-lint
	golangci-lint run ./...