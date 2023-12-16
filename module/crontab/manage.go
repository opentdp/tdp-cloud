package crontab

import (
	"tdp-cloud/model"
	"tdp-cloud/model/cronjob"

	"github.com/robfig/cron/v3"
)

func NewById(userId, id uint) {

	job, err := cronjob.Fetch(&cronjob.FetchParam{Id: id})

	if err == nil && job.Id > 0 {
		NewByScriptJob(job)
	}

}

func UndoById(userId, id uint) {

	job, err := cronjob.Fetch(&cronjob.FetchParam{Id: id})

	if err == nil && job.Id > 0 {
		crontab.Remove(cron.EntryID(job.EntryId))
	}

}

func RedoById(userId, id uint) {

	job, err := cronjob.Fetch(&cronjob.FetchParam{Id: id})

	if err == nil && job.Id > 0 {
		crontab.Remove(cron.EntryID(job.EntryId))
		NewByScriptJob(job)
	}

}

func GetEntries(jobs []*model.Cronjob) map[uint]any {

	list := map[uint]any{}

	for _, job := range jobs {
		entry := crontab.Entry(cron.EntryID(job.EntryId))
		list[job.Id] = map[string]any{
			"EntryId":  entry.ID,
			"NextTime": entry.Next.Unix(),
			"PrevTime": entry.Prev.Unix(),
		}
	}

	return list

}
