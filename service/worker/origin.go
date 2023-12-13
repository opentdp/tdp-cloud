package worker

import (
	"time"

	"github.com/opentdp/go-helper/logman"

	"tdp-cloud/module/worker"
)

func origin() {

	defer timer()

	if err := worker.Connect(); err != nil {
		logman.Error(err.Error())
	}

}

func timer() {

	logman.Warn("Connection disconnected, retry in 15 seconds")

	time.Sleep(15 * time.Second)
	origin()

}
