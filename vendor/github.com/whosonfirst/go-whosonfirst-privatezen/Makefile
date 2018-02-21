CWD=$(shell pwd)
GOPATH := $(CWD)

prep:
	if test -d pkg; then rm -rf pkg; fi

rmdeps:
	if test -d src; then rm -rf src; fi 

self:   prep
	if test -d src; then rm -rf src; fi
	mkdir -p src/github.com/whosonfirst/go-whosonfirst-privatezen
	cp *.go src/github.com/whosonfirst/go-whosonfirst-privatezen/
	cp -r vendor/* src/

deps:
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-geojson-v2"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-placetypes"

vendor-deps: rmdeps deps
	if test -d vendor; then rm -rf vendor; fi
	cp -r src vendor
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

fmt:
	go fmt *.go

bin:	self
