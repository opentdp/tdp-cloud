package crontab

import (
	"log/slog"

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
var logger *slog.Logger

func Daemon() {

	logger = logman.NewLogger("crontab")
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
		logger.Error("run jobs", "error", err)
	}

}

func NewByScriptJob(job *model.Cronjob) error {

	sepc := job.Second + " " + job.Minute + " " + job.Hour + " " + job.DayofMonth + " " + job.Month + " " + job.DayofWeek
	entryId, err := crontab.AddFunc(sepc, func() {
		// 找不到脚本
		she, err := script.Fetch(&script.FetchParam{
			Id:     strutil.ToUint(job.Content),
			UserId: job.UserId,
		})
		if err != nil {
			logger.Error("计划任务执行失败，找不到脚本", "error", err)
			return
		}
		// 找不到目标
		mac, err := machine.Fetch(&machine.FetchParam{
			Id:     strutil.ToUint(job.Target),
			UserId: job.UserId,
		})
		if err != nil {
			logger.Error("计划任务执行失败，找不到目标", "error", err)
			return
		}
		// 节点已断开
		send := workhub.GetSendPod(mac.WorkerId)
		if send == nil {
			logger.Error("计划任务执行失败，节点已断开", "workerId", mac.WorkerId)
			return
		}
		// 执行计划任务
		send.Exec(&command.ExecPayload{
			Name:          "Cron: " + she.Name,
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
