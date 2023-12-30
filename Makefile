#!/bin/zsh


.PHONY: build
build:
#	GOOS=linux  GOARCH=amd64  go build -o oj main.go
	go build -o oj main.go
