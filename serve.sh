#!/bin/sh
#

export GIN_MODE=debug

####################################################################

go mod tidy
go run main.go --listen 127.0.0.1:7800
