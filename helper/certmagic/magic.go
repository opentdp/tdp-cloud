package certmagic

import (
	"context"
	"errors"
	"strings"

	"github.com/caddyserver/certmagic"
	"github.com/spf13/viper"
)

var (
	sMagic = map[string]*certmagic.Config{}
	dMagic = map[string]*certmagic.Config{}
)

func Manage(rq *Params) error {

	magic, ok := sMagic[rq.SecretKey]

	if !ok {
		magic = CreateMagic()
		sMagic[rq.SecretKey] = magic
		// 创建发行人信息
		magic.Issuers = []certmagic.Issuer{
			certmagic.NewACMEIssuer(magic, *newIssuer(rp)),
		}
	}

	dMagic[rq.Domain] = magic

	domains := strings.Split(rq.Domain, ",")
	return magic.ManageAsync(context.Background(), domains)

}

func Unmanage(domain string) {

	magic, ok := dMagic[domain]
	domains := strings.Split(domain, ",")

	if ok {
		delete(sMagic, domain)
		magic.Unmanage(domains)
	}

}

func Certificate(domain string) (certmagic.Certificate, error) {

	if magic, ok := dMagic[domain]; ok {
		return magic.CacheManagedCertificate(context.Background(), domain)
	}

	return certmagic.Certificate{}, errors.New("域名不存在")

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
