#!/bin/bash


GOOS=darwin GOARCH=arm64 go build -o bin/gommit_darwin_arm64 -ldflags "-X main.debugMode=false"
GOOS=darwin GOARCH=amd64 go build -o bin/gommit_darwin_amd64 -ldflags "-X main.debugMode=false"

GOOS=windows GOARCH=amd64 go build -o bin/gommit_windows_amd64.exe -ldflags "-X main.debugMode=false"

GOOS=linux GOARCH=arm64 go build -o bin/gommit_linux_amd64 -ldflags "-X main.debugMode=false"
GOOS=linux GOARCH=amd64 go build -o bin/gommit_linux_amd64 -ldflags "-X main.debugMode=false"
# GOOS=linux GOARCH=amd64 go build -o bin/app-amd64-linux app.go