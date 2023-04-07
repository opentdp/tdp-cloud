package certbot

import (
	"tdp-cloud/cmd/args"
	"tdp-cloud/helper/certmagic"
	"tdp-cloud/helper/logman"
	"tdp-cloud/module/model"
	"tdp-cloud/module/model/certjob"
	"tdp-cloud/module/model/user"
	"tdp-cloud/module/model/vendor"
)

func Daemon() {

	certmagic.CertEvent = SetHistory

	go RunJobs()

}

func RunJobs() {

	jobs, err := certjob.FetchAll(&certjob.FetchAllParam{})

	if err != nil || len(jobs) == 0 {
		return
	}

	for _, job := range jobs {
		NewByJob(job)
	}

}

func NewByJob(job *model.Certjob) error {

	ur, err := user.Fetch(&user.FetchParam{
		Id:       job.UserId,
		StoreKey: args.Dataset.Secret,
	})

	if err != nil || ur.AppKey == "" {
		logman.Error("Get AppKey failed", "domain", job.Domain)
		return err
	}

	vd, err := vendor.Fetch(&vendor.FetchParam{
		Id:       job.VendorId,
		UserId:   job.UserId,
		StoreKey: ur.AppKey,
	})

	if err != nil || vd.SecretKey == "" {
		logman.Error("Get SecretKey failed", "domain", job.Domain)
		return err
	}

	if args.Debug {
		job.CaType = "debug" //调试模式强制重写
	}

	return certmagic.Manage(&certmagic.ReqeustParam{
		Email:     job.Email,
		Domain:    job.Domain,
		CaType:    job.CaType,
		Provider:  vd.Provider,
		SecretId:  vd.SecretId,
		SecretKey: vd.SecretKey,
		EabKeyId:  job.EabKeyId,
		EabMacKey: job.EabMacKey,
	})

}
