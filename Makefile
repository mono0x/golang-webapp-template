GO=go
GOBIN=$(PWD)/bin
TESTOPTS=-coverprofile=result.coverprofile -v -race ./...
BUILDOPTS=-tags netgo -installsuffix netgo -ldflags "-w -s -extldflags -static"
BINARY=golang-webapp-template

all: deps test build

setup:
	GOBIN=$(GOBIN) GO111MODULE=on go install github.com/google/wire/cmd/wire
	GOBIN=$(GOBIN) GO111MODULE=on go install github.com/jessevdk/go-assets-builder
	GOBIN=$(GOBIN) GO111MODULE=on go install github.com/lestrrat-go/server-starter/cmd/start_server
	GOBIN=$(GOBIN) GO111MODULE=on go install honnef.co/go/tools/cmd/megacheck

deps:
	GO111MODULE=on go mod tidy

test:
	GO111MODULE=on $(GO) mod verify
	GO111MODULE=on $(GO) vet ./...
	GO111MODULE=on $(GO) test $(TESTOPTS)

npm-build:
	cd client; npm run build

generate:
	GO111MODULE=on PATH=$(GOBIN):$(PATH) $(GO) generate -tags=wireinject ./...

build: npm-build generate
	GO111MODULE=on $(GO) build -o $(BINARY) $(BUILDOPTS)

run: generate
	GO111MODULE=on go run -tags dev $(shell go list -tags dev -f '{{range .GoFiles}}{{$$.Dir}}/{{.}}{{"\n"}}{{end}}')

build-linux:
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build -o $(BINARY).linux $(BUILDOPTS)
