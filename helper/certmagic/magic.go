package acme

import (
	"context"
	"log"
	"os"

	"github.com/caddyserver/certmagic"
	"github.com/libdns/alidns"
	"github.com/libdns/cloudflare"
)

type CertParam struct {
	Username    string
	StoragePath string
	Email       string
	Domain      []string
	Provider    string
	SecretId    string
	SecretKey   string
}

func MagicManage(rp *CertParam) error {

	magic := newConfig(rp)

	magic.Issuers = []certmagic.Issuer{
		certmagic.NewACMEIssuer(magic, *newIssuer(rp)),
	}

	return magic.ManageSync(context.TODO(), rp.Domain)

}

func newConfig(rp *CertParam) *certmagic.Config {

	config := certmagic.Config{
		Storage: &certmagic.FileStorage{
			Path: rp.StoragePath + "/" + rp.Username,
		},
		KeySource: certmagic.StandardKeyGenerator{
			KeyType: certmagic.RSA2048,
		},
		OnEvent: func(_ context.Context, event string, data map[string]any) error {
			log.Printf("Event: %s with data: %v\n", event, data)
			return nil
		},
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
