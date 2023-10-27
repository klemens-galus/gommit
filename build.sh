#!/bin/bash


GOOS=darwin GOARCH=arm64 go build -o bin/gommit_darwin_arm64
GOOS=darwin GOARCH=amd64 go build -o bin/gommit_darwin_amd64

GOOS=windows GOARCH=arm64
GOOS=windows GOARCH=amd64

GOOS=linux GOARCH=arm64
GOOS=linux GOARCH=amd64
# GOOS=linux GOARCH=amd64 go build -o bin/app-amd64-linux app.go
