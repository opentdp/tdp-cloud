package certmagic

import (
	"context"
	"strings"

	"github.com/caddyserver/certmagic"
	"github.com/spf13/viper"
)

var (
	sMagic = map[string]*certmagic.Config{}
	dMagic = map[string]*certmagic.Config{}
)

func Manage(rp *Params) error {

	magic, ok := sMagic[rp.SecretKey]

	if !ok {
		magic = CreateMagic()
		sMagic[rp.SecretKey] = magic
		// 创建发行人信息
		magic.Issuers = []certmagic.Issuer{
			certmagic.NewACMEIssuer(magic, *newIssuer(rp)),
		}
	}

	dMagic[rp.Domain] = magic

	domains := strings.Split(rp.Domain, ",")
	return magic.ManageAsync(context.TODO(), domains)

}

func Unmanage(domain string) {

	magic, ok := dMagic[domain]
	domains := strings.Split(domain, ",")

	if ok {
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
