package certmagic

import (
	"context"
	"os"

	"github.com/caddyserver/certmagic"
	"github.com/libdns/alidns"
	"github.com/libdns/cloudflare"

	tencent "github.com/rehiy/libdns-tencentcloud"
)

type Params struct {
	Email     string
	Domain    []string
	Provider  string
	SecretId  string
	SecretKey string
	StorePath string
}

func Async(rp *Params) error {

	magic := newMagic(rp)

	magic.Issuers = []certmagic.Issuer{
		certmagic.NewACMEIssuer(magic, *newIssuer(rp)),
	}

	return magic.ManageAsync(context.TODO(), rp.Domain)

}

func newMagic(rp *Params) *certmagic.Config {

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

func newIssuer(rp *Params) *certmagic.ACMEIssuer {

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
		issuer.DNS01Solver = &certmagic.DNS01Solver{
			DNSProvider: &tencent.Provider{
				SecretId:  rp.SecretId,
				SecretKey: rp.SecretKey,
			},
		}
	}

	return issuer

}
