@ECHO OFF
::

SET IS_DEBUG=1

::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

CD /d %~dp0

CALL go mod tidy
CALL go run main.go --address 127.0.0.1:7800

IF "%1" == "" CMD /K
