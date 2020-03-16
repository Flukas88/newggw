all: build
version=`git rev-parse --short HEAD`

build:
		go build -ldflags="-X 'main.version=$(version)'"