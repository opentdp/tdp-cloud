package acme

import (
	"context"
	"os"

	"github.com/caddyserver/certmagic"
	"github.com/libdns/alidns"
	"github.com/libdns/cloudflare"
)

type CertParam struct {
	Email     string
	Domain    []string
	Provider  string
	SecretId  string
	SecretKey string
	StorePath string
}

func Manage(rp *CertParam) error {

	magic := newMagic(rp)

	magic.Issuers = []certmagic.Issuer{
		certmagic.NewACMEIssuer(magic, *newIssuer(rp)),
	}

	return magic.ManageSync(context.TODO(), rp.Domain)

}

func newMagic(rp *CertParam) *certmagic.Config {

	config := certmagic.Config{
		Storage: &certmagic.FileStorage{Path: rp.StorePath},
		OnEvent: magicEvent,
	}

	cache := certmagic.NewCache(certmagic.CacheOptions{
		GetConfigForCert: func(cert certmagic.Certificate) (*certmagic.Config, error) {
			return &config, nil
		},
	})

	return certmagic.New(cache, config)

}

func newIssuer(rp *CertParam) *certmagic.ACMEIssuer {

	issuer := &certmagic.ACMEIssuer{
		Email:  rp.Email,
		Agreed: true,
	}

	if os.Getenv("TDP_DEBUG") == "" {
		issuer.CA = certmagic.ZeroSSLProductionCA
	} else {
		issuer.CA = certmagic.LetsEncryptStagingCA
	}

	switch rp.Provider {
	case "alibaba":
		issuer.DNS01Solver = &certmagic.DNS01Solver{
			DNSProvider: &alidns.Provider{
				AccKeyID:     rp.SecretId,
				AccKeySecret: rp.SecretKey,
			},
		}
	case "cloudflare":
		issuer.DNS01Solver = &certmagic.DNS01Solver{
			DNSProvider: &cloudflare.Provider{
				APIToken: rp.SecretKey,
			},
		}
	case "tencent":
		issuer.DNS01Solver = &certmagic.DNS01Solver{}
	}

	return issuer

}
