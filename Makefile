
default: main

.PHONY: setup
setup: setup-common
	go install github.com/golangci/golangci-lint/...

.PHONY: setup-common
setup-common:
	go get -u golang.org/x/tools/cmd/goimports
	go get -t -v ./...

.PHONY: setup-ci
setup-ci: setup-common
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s v1.12.3

.PHONY: fmt
fmt:
	goimports -w=true -d .

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -cover -race ./...

.PHONY: release
release:
	go build -ldflags '-s' -o ju .
