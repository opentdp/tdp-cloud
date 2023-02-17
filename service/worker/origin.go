package worker

import (
	"log"
	"time"

	"github.com/kardianos/service"
	"github.com/spf13/viper"

	"tdp-cloud/module/worker"
)

type program struct{}

func (p *program) Start(s service.Service) error {

	log.Println("TDP Worker start")

	go p.run()
	return nil

}

func (p *program) Stop(s service.Service) error {

	log.Println("TDP Worker stop")

	return nil

}

func (p *program) run() {

	defer p.timer()

	remote := viper.GetString("worker.remote")

	if err := worker.Daemon(remote); err != nil {
		log.Println(err)
	}

}

func (p *program) timer() {

	log.Println("Connection disconnected, retry in 5 seconds.")

	time.Sleep(5 * time.Second)
	p.run()

}
