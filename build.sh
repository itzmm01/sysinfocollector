#!/bin/bash

name="sysinfocollector"
build(){
     
	go build main.go
	mkdir ${name}-$GOOS-$GOARCH
	mv main ${name}-$GOOS-$GOARCH/${name}
	cp install.sh ${name}-$GOOS-$GOARCH/
	cp README.md ${name}-$GOOS-$GOARCH/
	tar zcvf ${name}-$GOOS-$GOARCH.tar.gz ${name}-$GOOS-$GOARCH
	rm -rf ${name}-$GOOS-$GOARCH/
}


# linux
export CGO_ENABLED=0
export GOOS=linux

# amd64
export GOARCH=amd64
echo $GOOS-$GOARCH
build

# arm64
export GOARCH=arm64
echo $GOOS-$GOARCH
build
