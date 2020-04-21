all: gofmt test

test:
	go test -v ./...

gofmt:
	go fmt ./...

build_server:
	$(MAKE) -C cmd/server build

build_client:
	$(MAKE) -C cmd/client build

clean_server:
	$(MAKE) -C cmd/server clean

clean_client:
	$(MAKE) -C cmd/client clean
