package certmagic

import (
	"context"
	"errors"
	"strings"
	"tdp-cloud/helper/strutil"

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

	skey := strutil.Md5(rq.Email + rq.SecretKey + rq.CaType)

	magic, ok := sMagic[skey]

	if !ok {
		magic = CreateMagic()
		magic.Issuers = []certmagic.Issuer{
			certmagic.NewACMEIssuer(magic, *newIssuer(rq)),
		}
		// 写入缓存
		sMagic[skey] = magic
	}

	dMagic[rq.Domain] = magic

	domains := strings.Split(rq.Domain, ",")
	return magic.ManageAsync(context.Background(), domains)

}

func Unmanage(domain string) {

	magic, ok := dMagic[domain]
	domains := strings.Split(domain, ",")

	if ok {
		magic.Unmanage(domains)
		delete(dMagic, domain)
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
