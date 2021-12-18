all: format check-tidy
	go test ./...

.PHONY: test
test:
	go test -v ./...

.PHONY: check-tidy
check-tidy:
	go mod tidy

.PHONY: ci
ci: format check-tidy test

.PHONY: format
format:
	go fmt ./...

