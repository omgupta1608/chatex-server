@echo off

set JWT_SECRET="SECRET"
go run cmd/main.go %*
