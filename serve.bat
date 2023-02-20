@ECHO OFF
::

SET TDP_DEBUG=1

SET CGO_ENABLED=0
SET GO111MODULE=on

::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

CD /d %~dp0

go mod tidy
go run main.go server --listen 127.0.0.1:7800
