@ECHO OFF
::

SET TDP_DEBUG=1

SET CGO_ENABLED=0
SET GO111MODULE=on

::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

CD /d %~dp0

CALL go mod tidy
CALL go run main.go --listen 127.0.0.1:7800

IF "%1" == "" CMD /K
