package certbot

import (
	"github.com/open-tdp/go-helper/certmagic"
	"github.com/open-tdp/go-helper/logman"

	"tdp-cloud/cmd/args"
	"tdp-cloud/model"
	"tdp-cloud/model/certjob"
	"tdp-cloud/model/user"
	"tdp-cloud/model/vendor"
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
		logman.Error("get AppKey failed", "domain", job.Domain)
		return err
	}

	vd, err := vendor.Fetch(&vendor.FetchParam{
		Id:       job.VendorId,
		UserId:   job.UserId,
		StoreKey: ur.AppKey,
	})

	if err != nil || vd.SecretKey == "" {
		logman.Error("get SecretKey failed", "domain", job.Domain)
		return err
	}

	if args.Debug {
		job.CaType = "debug" //调试模式强制重写
	}

	return certmagic.Manage(&certmagic.ReqeustParam{
		Email:       job.Email,
		Domain:      job.Domain,
		CaType:      job.CaType,
		Provider:    vd.Provider,
		SecretId:    vd.SecretId,
		SecretKey:   vd.SecretKey,
		EabKeyId:    job.EabKeyId,
		EabMacKey:   job.EabMacKey,
		StoragePath: args.Dataset.Dir + "/certmagic",
	})

}
