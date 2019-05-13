#-----------------------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See LICENSE in the project root for license information.
#-----------------------------------------------------------------------------------------

FROM golang:1

RUN go get -u -v \
    github.com/mdempsky/gocode \
    github.com/uudashr/gopkgs/cmd/gopkgs \
    github.com/ramya-rao-a/go-outline \
    github.com/acroca/go-symbols \
    golang.org/x/tools/cmd/guru \
    golang.org/x/tools/cmd/gorename \
    github.com/rogpeppe/godef \
    github.com/zmb3/gogetdoc \
    github.com/sqs/goreturns \
    golang.org/x/tools/cmd/goimports \
    golang.org/x/lint/golint \
    github.com/alecthomas/gometalinter \
    honnef.co/go/tools/... \
    github.com/golangci/golangci-lint/cmd/golangci-lint \
    github.com/mgechev/revive \
    github.com/derekparker/delve/cmd/dlv 2>&1

# gocode-gomod
RUN go get -x -d github.com/stamblerre/gocode \
    && go build -o gocode-gomod github.com/stamblerre/gocode \
    && mv gocode-gomod $GOPATH/bin/

# Copy default endpoint specific user settings overrides into container to specify Go path
COPY settings.vscode.json /root/.vscode-remote/data/Machine/settings.json

# Install dep
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Install git, process tools and other dependencies
RUN apt-get update && apt-get -y install \
  curl \
  git \
  make \
  gcc \
  tree \
  vim \
  procps

# Drop user into correct working folder for go code in the terminal
RUN echo "cd /go/src/vscode-remote-try-go" >> ~/.bashrc

WORKDIR /go/src

# Clean up
RUN apt-get autoremove -y \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/*

