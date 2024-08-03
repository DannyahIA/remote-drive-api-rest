@echo off

REM Set the necessary environment variables
set GOPATH=C:\path\to\your\go\workspace
set PATH=%PATH%;%GOPATH%\bin

REM Build the Go program
go build -o ../bin/windows/main.exe ../cmd/personal_server/main.go 
