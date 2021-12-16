.PHONY: test
test:
	go test ./...

.PHONY: check-tidy
check-tidy:
	go mod tidy

.PHONY: ci
ci: check-tidy test
