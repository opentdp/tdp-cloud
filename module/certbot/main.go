package certbot

import (
	"log"
	"strings"

	"tdp-cloud/helper/certmagic"
	"tdp-cloud/module/dborm"
	"tdp-cloud/module/dborm/certbot"
	"tdp-cloud/module/dborm/vendor"
)

func Daemon() {

	certbots, err := certbot.FetchAll(&certbot.FetchAllParam{})

	if err != nil || len(certbots) == 0 {
		return
	}

	for _, bot := range certbots {
		NewTask(bot)
	}

}

func NewTask(bot *dborm.Certbot) error {

	vendor, err := vendor.Fetch(&vendor.FetchParam{
		Id: bot.VendorId, UserId: bot.UserId,
	})

	if err != nil || vendor.Id == 0 {
		log.Println("Certbot Ignore Domain:", bot.Domain)
		return err
	}

	return certmagic.Async(&certmagic.Params{
		Domain:    strings.Split(bot.Domain, ","),
		Email:     bot.Email,
		Provider:  vendor.Provider,
		SecretId:  vendor.SecretId,
		SecretKey: vendor.SecretKey,
		StorePath: "./data/certs/" + bot.Email,
	})

}
