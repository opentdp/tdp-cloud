@ECHO OFF
::

SET GIN_MODE=debug

::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

CD /d %~dp0

go mod tidy
go run main.go --address 127.0.0.1:7800

IF "%1" == "" CMD /K
