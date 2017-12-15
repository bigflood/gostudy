#!/bin/bash

export GOOS=linux
export GOARCH=amd64
go build -o webchat-svr main.go

docker build . -t webchat-svr:latest
