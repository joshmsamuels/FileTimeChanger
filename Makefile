app:
	go build -o bin/changeModTime cmd/main.go

run: app
	bin/changeModTime

fmt:
	go fmt ./...