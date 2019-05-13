all: test build

build: install-dependencies fmt
	go build -o /go/src/vscode-remote-try-go/bin/server /go/src/vscode-remote-try-go/cmd/server

install-dependencies:
	dep ensure

fmt:
	go fmt /go/src/vscode-remote-try-go/pkg/... /go/src/vscode-remote-try-go/cmd/server...

test:
	go test /go/src/vscode-remote-try-go/pkg/... /go/src/vscode-remote-try-go/cmd/server... -coverprofile cover.out