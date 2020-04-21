all: build
version=`git describe --tags --abbrev=12`

build:
		go build -ldflags="-X 'main.version=$(version)'"

clean:
	go clean
