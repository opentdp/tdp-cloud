package certbot

import (
	"log"
	"strings"

	"github.com/spf13/cast"
	"github.com/spf13/viper"

	"tdp-cloud/helper/certmagic"
	"tdp-cloud/module/dborm"
	"tdp-cloud/module/dborm/certjob"
	"tdp-cloud/module/dborm/vendor"
)

func Daemon() {

	jobs, err := certjob.FetchAll(&certjob.FetchAllParam{})

	if err != nil || len(jobs) == 0 {
		return
	}

	for _, job := range jobs {
		NewTask(job)
	}

}

func NewTask(job *dborm.Certjob) error {

	vendor, err := vendor.Fetch(&vendor.FetchParam{
		Id: job.VendorId, UserId: job.UserId,
	})

	if err != nil || vendor.Id == 0 {
		log.Println("Certjob Ignore Domain:", job.Domain)
		return err
	}

	dir := viper.GetString("dataset.dir") + "/certbot-" + cast.ToString(job.UserId)

	return certmagic.Async(&certmagic.Params{
		Domain:    strings.Split(job.Domain, ","),
		Email:     job.Email,
		Provider:  vendor.Provider,
		SecretId:  vendor.SecretId,
		SecretKey: vendor.SecretKey,
		StorePath: dir,
	})

}
