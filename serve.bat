@ECHO OFF
::

SET GIN_MODE=debug

::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

CD /d %~dp0

go mod tidy
go run main.go

IF "%1" == "" CMD /K
