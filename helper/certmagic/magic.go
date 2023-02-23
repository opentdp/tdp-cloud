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

type Certificate struct {
	Names       []string
	OCSPStaple  []byte
	Certificate [][]byte
	PrivateKey  any
}

func Manage(rq *Params) error {

	magic, ok := sMagic[rq.SecretKey]

	if !ok {
		magic = CreateMagic()
		sMagic[rq.SecretKey] = magic
		// 创建发行人信息
		magic.Issuers = []certmagic.Issuer{
			certmagic.NewACMEIssuer(magic, *newIssuer(rq)),
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

func CertData(domain string) (*Certificate, error) {

	cert := &Certificate{}

	magic, ok := dMagic[domain]

	if !ok {
		return cert, errors.New("域名不存在")
	}

	crt, err := magic.CacheManagedCertificate(context.Background(), domain)

	if err != nil {
		return cert, err
	}

	cert.Names = crt.Names
	cert.OCSPStaple = crt.Certificate.OCSPStaple
	cert.PrivateKey = crt.Certificate.PrivateKey
	cert.Certificate = crt.Certificate.Certificate

	return cert, nil

}
