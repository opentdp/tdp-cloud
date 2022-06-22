#!/bin/sh
#

export IS_DEBUG=1

####################################################################

go mod tidy
go run main.go --address 127.0.0.1:7800
