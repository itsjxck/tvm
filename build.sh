#!/bin/bash

GOOS=darwin GOARCH=amd64 go build -o builds/tvm_darwin_amd64
GOOS=linux GOARCH=amd64 go build -o builds/tvm_linux_amd64