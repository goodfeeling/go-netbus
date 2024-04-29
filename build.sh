#!/bin/bash

go mod tidy
# 构建 Go 项目
go build -o netbus main.go

echo "Build completed."