package crontab

import (
	"github.com/opentdp/go-helper/command"
	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/strutil"
	cron "github.com/robfig/cron/v3"

	"tdp-cloud/model"
	"tdp-cloud/model/cronjob"
	"tdp-cloud/model/machine"
	"tdp-cloud/model/script"
	"tdp-cloud/module/workhub"
)

var crontab *cron.Cron

func Daemon() {

	crontab = cron.New(cron.WithSeconds())

	RunJobs()

}

func RunJobs() {

	jobs, err := cronjob.FetchAll(&cronjob.FetchAllParam{})

	if err != nil || len(jobs) == 0 {
		return
	}

	for _, job := range jobs {
		NewByJob(job)
	}

	crontab.Start()

}

func NewByJob(job *model.Cronjob) {

	var err error

	switch job.Type {
	case "script":
		err = NewByScriptJob(job)
	}

	if err != nil {
		logman.Error("run jobs", "error", err)
	}

}

func NewByScriptJob(job *model.Cronjob) error {

	she, err := script.Fetch(&script.FetchParam{
		Id:     strutil.ToUint(job.Content),
		UserId: job.UserId,
	})
	if err != nil {
		return err
	}

	mac, err := machine.Fetch(&machine.FetchParam{
		Id:     strutil.ToUint(job.Location),
		UserId: job.UserId,
	})
	if err != nil {
		return err
	}

	spec := job.Second + " " + job.Minute + " " + job.Hour + " " + job.DayofMonth + " " + job.Month + " " + job.DayofWeek

	entryId, err := crontab.AddFunc(spec, func() {
		workhub.GetSendPod(mac.WorkerId).Exec(&command.ExecPayload{
			Name:          she.Name,
			CommandType:   she.CommandType,
			Username:      she.Username,
			WorkDirectory: she.WorkDirectory,
			Content:       she.Content,
			Timeout:       she.Timeout,
		})
	})
	if err != nil {
		return err
	}

	err = cronjob.Update(&cronjob.UpdateParam{
		Id:      job.Id,
		EntryId: int64(entryId),
	})

	return err

}
