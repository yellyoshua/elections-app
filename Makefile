#!bin/bash


#
# Makefile
# hzsunshx, 2015-02-11 13:17
#

go-src-project = github.com/yellyoshua/elections-app

src-ballot-main = ballot/ballot.go
src-suffrage-main = suffrage/suffrage.go

install:
	go get -u -v ./...
	go get github.com/vektra/mockery/v2/.../

# test:
# 	mockery --dir repository --all --output ./commons/mocks/repository/ --keeptree --inpackage
# 	go test -timeout 30s $(go-src-project)/modules/graphql
# 	go test -timeout 30s $(go-src-project)modules/authentication
# 	go test -timeout 30s $(go-src-project)/api
# 	go test -timeout 30s $(go-src-project)/utils
# 	go test -timeout 30s $(go-src-project)/middlewares
# 	go test -timeout 30s $(go-src-project)/handlers
# 	go test -timeout 30s $(go-src-project)/modules/storage
# 	go test -timeout 30s $(go-src-project)/repository

clean-dependencies:
	go mod tidy

build:
	make ballot-build
	make suffrage-build

suffrage-tests:
	echo "$(src-suffrage-main)"

ballot-tests:
	echo "$(src-suffrage-main)"

ballot-build:
	go build -a -o elections-ballot $(src-ballot-main)

suffrage-build:
	go build -a -o elections-suffrage $(src-suffrage-main)

ballot-cross-build:
	GOOS=windows GOARCH=386 go build -a -o elections-ballot-win32 $(src-ballot-main)
	GOOS=linux GOARCH=386 go build -a -o elections-ballot-linux-386 $(src-ballot-main)
	GOOS=linux GOARCH=amd64 go build -a -o elections-ballot-linux-amd64 $(src-ballot-main)

suffrage-cross-build:
	GOOS=windows GOARCH=386 go build -a -o elections-suffrage-win32 $(src-suffrage-main)
	GOOS=linux GOARCH=386 go build -a -o elections-suffrage-linux-386 $(src-suffrage-main)
	GOOS=linux GOARCH=amd64 go build -a -o elections-suffrage-linux-amd64 $(src-suffrage-main)

cross-build:
	GOOS=windows GOARCH=386 go build -a -o elections-suffrage-win32 $(src-suffrage-main)
	GOOS=linux GOARCH=386 go build -a -o elections-suffrage-linux-386 $(src-suffrage-main)
	GOOS=linux GOARCH=amd64 go build -a -o elections-suffrage-linux-amd64 $(src-suffrage-main)

	GOOS=windows GOARCH=386 go build -a -o elections-ballot-win32 $(src-ballot-main)
	GOOS=linux GOARCH=386 go build -a -o elections-ballot-linux-386 $(src-ballot-main)
	GOOS=linux GOARCH=amd64 go build -a -o elections-ballot-linux-amd64 $(src-ballot-main)
