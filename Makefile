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

cross-build:
	GOOS=windows GOARCH=386 go build -o fileserv-win32
	GOOS=linux GOARCH=386 go build -o fileserv-linux-386
	GOOS=linux GOARCH=amd64 go build -o fileserv-linux-amd64
