package worker

import (
	"time"

	"github.com/kardianos/service"

	"tdp-cloud/module/worker"
)

type program struct{}

func (p *program) Start(s service.Service) error {

	svclog.Info("TDP Worker start")

	go p.run()
	return nil

}

func (p *program) Stop(s service.Service) error {

	svclog.Info("TDP Worker stop")

	return nil

}

func (p *program) run() {

	defer p.timer()

	if err := worker.Connect(); err != nil {
		svclog.Error(err)
	}

}

func (p *program) timer() {

	svclog.Warning("Connection disconnected, retry in 5 seconds.")

	time.Sleep(5 * time.Second)
	p.run()

}
