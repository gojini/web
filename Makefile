export GO111MODULE=on
GO_SRC=$(shell find . -path ./.build -prune -false -o -name \*.go)

.PHONY: all
all: lint test

test_certs:
	./gen_certs.sh

test: $(GO_SRC) test_certs example_tests server_tests

example_tests: $(GO_SRC) test_certs
	go test -v -race -cover -coverpkg ./... -coverprofile=example_coverage.txt -covermode=atomic ./...

lint: ./.golangcilint.yaml
	./bin/golangci-lint --version || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./bin v1.46.2
	./bin/golangci-lint --config ./.golangcilint.yaml run ./...

.PHONY: server_tests
server_tests: $(GO_SRC) test_certs
	cd ./tests && go test -v -race -cover -coverpkg "gojini.dev/web" -coverprofile=../coverage.txt -covermode=atomic ./...