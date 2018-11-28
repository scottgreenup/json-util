
default: main

.PHONY: setup
setup:
	go get -u golang.org/x/tools/cmd/goimports
	go get -t -v ./...

.PHONY: fmt
fmt:
	goimports -w=true -d .

.PHONY: test
test:
	go test -cover -race ./...

.PHONY: release
release:
	go build -ldflags '-s' -o ju .
