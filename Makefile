#!bin/bash


#
# Makefile
# hzsunshx, 2015-02-11 13:17
#

install:
	go get .
	go get github.com/vektra/mockery/v2/.../

test:
	mockery --dir repository --all --output ./mocks/repository/ --keeptree --inpackage
	go test -timeout 30s github.com/yellyoshua/elections-app/modules/graphql
	go test -timeout 30s github.com/yellyoshua/elections-app/modules/authentication
	go test -timeout 30s github.com/yellyoshua/elections-app/api
	go test -timeout 30s github.com/yellyoshua/elections-app/utils
	go test -timeout 30s github.com/yellyoshua/elections-app/middlewares
	go test -timeout 30s github.com/yellyoshua/elections-app/handlers
	go test -timeout 30s github.com/yellyoshua/elections-app/modules/storage
	go test -timeout 30s github.com/yellyoshua/elections-app/repository

clean-dependencies:
	go mod tidy

build:
	go build .

dev: clean
	npm start

prod: bundle binary

bundle:
	webpack -p --progress --config webpack.config.prod.js

binary:
	go get github.com/jteeuwen/go-bindata/...
	(cd public; go-bindata -pkg public favicon.ico bundle.js js/ css/)
	(cd templates; go-bindata -pkg templates ./...)
	go get -tags "bindata"
	go build -tags "bindata"

install-deps:
	sudo apt-get update -qq
	sudo apt-get install -qq nodejs npm

deps:
	npm install

cross-build:
	GOOS=windows GOARCH=386 go build
	GOOS=linux GOARCH=386 go build -o fileserv-linux-386
	GOOS=linux GOARCH=amd64 go build -o fileserv-linux-amd64

webpack:
	webpack

clean:
	rm -f public/bundle.js
# vim:ft=make
#