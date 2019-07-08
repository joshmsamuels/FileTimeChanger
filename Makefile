all: fmt run

build:
	go build -o bin/changeModTime cmd/main.go

run-only:
	bin/changeModTime

run: build
	bin/changeModTime

fmt:
	go fmt ./...