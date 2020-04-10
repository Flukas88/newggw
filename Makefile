all: gofmt test

test:
	go test -v ./...

gofmt:
	go fmt ./...

build_server:
	$(MAKE) -C server build

build_client:
	$(MAKE) -C client build

clean_server:
	$(MAKE) -C server clean

clean_client:
	$(MAKE) -C client clean
