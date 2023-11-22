@echo off
::

SET CGO_ENABLED=0
SET GO111MODULE=on

::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

CALL :build android arm64

CALL :build darwin amd64
CALL :build darwin arm64

CALL :build freebsd 386
CALL :build freebsd amd64
CALL :build freebsd arm64

CALL :build linux 386
CALL :build linux amd64
CALL :build linux arm64
CALL :build linux ppc64le
CALL :build linux s390x

CALL :build netbsd amd64

CALL :build openbsd amd64
CALL :build openbsd arm64

CALL :build windows amd64
CALL :build windows arm64

GOTO :EOF

::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

:build
  SET GOOS=%1
  SET GOARCH=%2
  SET target=build/tdp-cloud-%1-%2
  IF "%1"=="windows" (
    SET target=%target%.exe
  )
  ECHO building for %1/%2
  go build -ldflags="-s -w" -o %target% main.go
  GOTO :EOF

IF "%1" == "" CMD /K
