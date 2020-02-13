all: gofmt test

test:
	go test -v ./...

gofmt:
	go fmt ./...
