VERSION := $(shell git describe --tags --abbrev=0)

build:
	go build -o open-dir-${VERSION}

clean: 
	rm -rf open-dir
