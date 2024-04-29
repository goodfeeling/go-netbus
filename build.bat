@echo off
go mod tidy
go build -o netbus.exe main.go

echo Build completed.