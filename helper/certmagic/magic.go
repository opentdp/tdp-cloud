package certmagic

import (
	"context"
	"strings"

	"github.com/caddyserver/certmagic"
	"github.com/spf13/viper"
)

var magic *certmagic.Config

func Manage(rp *Params) error {

	if magic == nil {
		magic = CreateMagic()
	}

	issuer := certmagic.NewACMEIssuer(magic, *newIssuer(rp))
	magic.Issuers = append(magic.Issuers, issuer)

	domains := strings.Split(rp.Domain, ",")
	return magic.ManageAsync(context.TODO(), domains)

}

func Unmanage(domain string) {

	domains := strings.Split(domain, ",")

	if magic != nil {
		magic.Unmanage(domains)
	}

}

func CreateMagic() *certmagic.Config {

	config := certmagic.Config{
		Storage: &certmagic.FileStorage{
			Path: viper.GetString("dataset.dir") + "/certmagic",
		},
		OnEvent: magicEvent,
	}

	cache := certmagic.NewCache(certmagic.CacheOptions{
		GetConfigForCert: func(cert certmagic.Certificate) (*certmagic.Config, error) {
			return &config, nil
		},
	})

	return certmagic.New(cache, config)

}
