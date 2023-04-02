package crontab

import (
	cron "github.com/robfig/cron/v3"

	"tdp-cloud/module/model"
	"tdp-cloud/module/model/cronjob"
)

func Daemon() {

	cron.New(cron.WithSeconds())

}

func NewById(id uint) {

	job, err := cronjob.Fetch(&cronjob.FetchParam{Id: id})

	if err == nil && job.Id > 0 {
		NewByJob(job)
	}

}

func UndoById(id uint) {

	job, err := cronjob.Fetch(&cronjob.FetchParam{Id: id})

	if err == nil && job.Id > 0 {
		// TODO
		return
	}

}

func RedoById(id uint) {

	job, err := cronjob.Fetch(&cronjob.FetchParam{Id: id})

	if err == nil && job.Id > 0 {
		// TODO
		return
	}

}

func NewByJob(job *model.Cronjob) error {

	// TODO
	return nil

}
