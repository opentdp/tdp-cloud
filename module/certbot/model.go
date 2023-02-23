package certbot

import (
	"tdp-cloud/helper/certmagic"
	"tdp-cloud/module/dborm/certjob"
)

func NewById(userId, id uint) {

	job, err := certjob.Fetch(&certjob.FetchParam{Id: id, UserId: userId})

	if err == nil && job.Id > 0 {
		NewByJob(job)
	}

}

func UndoById(userId, id uint) {

	job, err := certjob.Fetch(&certjob.FetchParam{Id: id, UserId: userId})

	if err == nil && job.Id > 0 {
		certmagic.Unmanage(job.Domain)
	}

}

func RedoById(userId, id uint) {

	job, err := certjob.Fetch(&certjob.FetchParam{Id: id, UserId: userId})

	if err == nil && job.Id > 0 {
		certmagic.Unmanage(job.Domain)
		NewByJob(job)
	}

}
