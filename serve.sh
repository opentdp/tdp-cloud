#!/bin/sh
#

export GIN_MODE=debug

####################################################################

go mod tidy
go run main.go
