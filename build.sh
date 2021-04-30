#!/bin/bash

set -e

cov(){
	for GOOS in darwin linux windows; do
		go test -v -race -coverprofile=coverage-$GOOS.txt ./...
		go tool cover -html=coverage-$GOOS.txt -o cover-$GOOS.html
	done
	open cover-darwin.html
}

bf1(){
	for GOOS in $(go tool dist list|awk -F'/' '{print $1}'|sort -u); do echo -e "\n\nTESTING FOR $GOOS ...\n"; go test -v -race -coverprofile=coverage-$GOOS.txt -test.run=^TestDirTimestamps$ ./dir/ || exit; done
}

fmt(){
	echo fmt...
	gofmt -l -w -s .
}

lint(){
	echo lint...
	golint ./...
}

cyclo(){
	echo cyclo...
	gocyclo -top 10 .
}

all(){
	fmt && lint && cyclo
}


if [[ $# -eq 0 ]]; then
	cmd=cov
else
	cmd=${1:-cov} && shift
fi
$cmd "$@"