all: build
version=`git rev-parse HEAD`

build:
		go build -ldflags="-X 'main.version=$(version)'"