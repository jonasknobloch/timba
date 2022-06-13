#!/usr/bin/env bash

GOOS=darwin GOARCH=amd64 go build -o timba-amd64-darwin
GOOS=linux GOARCH=amd64 go build -o timba-amd64-linux
GOOS=windows GOARCH=amd64 go build -o timba-amd64.exe
