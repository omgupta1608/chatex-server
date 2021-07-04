@echo off

set JWT_SECRET="secret"
go run cmd/main.go %*
