package certbot

import (
	"github.com/opentdp/go-helper/certmagic"
	"github.com/opentdp/go-helper/logman"

	"tdp-cloud/cmd/args"
	"tdp-cloud/model"
	"tdp-cloud/model/certjob"
	"tdp-cloud/model/user"
	"tdp-cloud/model/vendor"
)

func Daemon() {

	certmagic.CertEvent = UpdateHistory

	RunJobs()

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

	usr, err := user.Fetch(&user.FetchParam{
		Id:       job.UserId,
		StoreKey: args.Assets.Secret,
	})

	if err != nil || usr.AppKey == "" {
		logman.Error("get AppKey failed", "domain", job.Domain)
		return err
	}

	vdr, err := vendor.Fetch(&vendor.FetchParam{
		Id:       job.VendorId,
		UserId:   job.UserId,
		StoreKey: usr.AppKey,
	})

	if err != nil || vdr.SecretKey == "" {
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
		Provider:    vdr.Provider,
		SecretId:    vdr.SecretId,
		SecretKey:   vdr.SecretKey,
		EabKeyId:    job.EabKeyId,
		EabMacKey:   job.EabMacKey,
		StoragePath: args.Assets.Dir + "/certmagic",
	})

}
