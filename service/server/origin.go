package server

import (
	"github.com/kardianos/service"
)

type origin struct{}

func (p *origin) Start(s service.Service) error {

	svclog.Info("TDP Server start")

	return p.run()

}

func (p *origin) Stop(s service.Service) error {

	svclog.Info("TDP Server stop")

	return nil

}

func (p *origin) run() error {

	go inlet()

	return nil

}
