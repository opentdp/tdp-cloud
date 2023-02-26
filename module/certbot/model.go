package certbot

import (
	"tdp-cloud/helper/certmagic"
	"tdp-cloud/module/dborm"
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

type certinfo struct {
	*dborm.Certjob
	Cert *certmagic.Certificate
}

func CertById(userId, id uint) (*certinfo, error) {

	job, err := certjob.Fetch(&certjob.FetchParam{Id: id, UserId: userId})

	if err == nil && job.Id > 0 {
		cert, err := certmagic.CertDetail(job.Domain)
		if err != nil {
			return nil, err
		}
		return &certinfo{job, cert}, nil
	}

	return nil, err

}
