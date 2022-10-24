#!/bin/sh
set GOARCH=amd64
go env -w GOARCH=amd64
set GOOS=linux
go env -w GOOS=linux
go build -o app  main.go
# 还原
go env -w GOARCH=amd64
go env -w GOOS=windows