@ECHO OFF
::

SET TDP_DEBUG=1

SET CGO_ENABLED=0
SET GO111MODULE=on

::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

CD /d %~dp0

CALL go mod tidy
CALL go run main.go server --listen 127.0.0.1:7800 --dsn cloud.db?_pragma=busy_timeout=5000&_pragma=journa_mode(WAL)

IF "%1" == "" CMD /K
