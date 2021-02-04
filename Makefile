.PHONY: fmt lint test

test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic

lint:
	golangci-lint run

fmt:
	`go list -f {{.Target}} mvdan.cc/gofumpt` -w .
