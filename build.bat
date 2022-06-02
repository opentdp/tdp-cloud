@echo off
::

SET CGO_ENABLED=0

::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

CALL :build android arm64

CALL :build darwin amd64
CALL :build darwin arm64

CALL :build linux 386
CALL :build linux amd64
CALL :build linux arm64

CALL :build windows 386 .exe
CALL :build windows amd64 .exe
CALL :build windows arm64 .exe

cmd /k

GOTO :EOF

::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

:build
  SET GOOS=%1
  SET GOARCH=%2
  echo building for %1/%2
  go build -o build/%1-%2%3 main.go
  GOTO :EOF
