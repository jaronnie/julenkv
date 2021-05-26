.PHONY: build
BIN_SERVER=julenkv-server
BIN-CLI=julenkv-cli

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/${BIN_SERVER} cmd/server/main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/${BIN-CLI} cmd/client/main.go
