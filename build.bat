@echo off
::

SET CGO_ENABLED=0
SET GIN_MODE=release

::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

CALL :build darwin amd64
CALL :build darwin arm64

CALL :build linux 386
CALL :build linux amd64
CALL :build linux arm64

CALL :build windows 386
CALL :build windows amd64
CALL :build windows arm64

cmd /k
GOTO :EOF

::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

:build
  SET GOOS=%1
  SET GOARCH=%2
  echo building for %1/%2
  go build -o build/%1-%2 main.go
  GOTO :EOF
