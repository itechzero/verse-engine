.PHONY: tidy
tidy:
	go mod tidy

.PHONY: lint
lint:
	golangci-lint run
