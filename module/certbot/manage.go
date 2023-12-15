package certbot

import (
	"github.com/opentdp/go-helper/certmagic"

	"tdp-cloud/model"
	"tdp-cloud/model/certjob"
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

func CertById(userId, id uint) (*model.Certjob, *certmagic.Certificate, error) {

	job, err := certjob.Fetch(&certjob.FetchParam{Id: id, UserId: userId})

	if err == nil && job.Id > 0 {
		cert, err := certmagic.CertDetail(job.Domain)
		if err == nil {
			return job, cert, nil
		}
		return job, nil, err
	}

	return nil, nil, err

}

func UpdateHistory(evt string, data map[string]any) {

	data["event"] = evt

	certjob.Update(&certjob.UpdateParam{
		Domain:  data["identifier"].(string),
		History: data,
	})

}
