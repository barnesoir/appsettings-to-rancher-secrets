#!/bin/bash
FN="ATRS"

echo "Building for Windows"
GOOS=windows
GOARCH=amd64
go build -o "bin/$FN.exe"

echo "Building for NIX"
GOOS=linux
GOARCH=amd64
go build -o "bin/$FN"

echo "done"