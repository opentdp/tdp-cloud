package certmagic

import (
	"context"
	"strings"

	"github.com/caddyserver/certmagic"

	"tdp-cloud/cmd/args"
	"tdp-cloud/helper/logman"
)

var CertEvent func(evt string, data map[string]any)

func newMagic(iss certmagic.ACMEIssuer) *certmagic.Config {

	config := certmagic.Config{
		Storage: &certmagic.FileStorage{
			Path: args.Dataset.Dir + "/certmagic",
		},
	}

	config.Issuers = []certmagic.Issuer{
		certmagic.NewACMEIssuer(&config, iss),
	}

	config.OnEvent = func(ctx context.Context, evt string, data map[string]any) error {
		logman.Named("cert.event").Warn(evt, "data", data)
		if CertEvent != nil {
			switch evt {
			case "cert_obtaining", "cert_failed", "cert_obtained":
				CertEvent(strings.Split(evt, "_")[1], data)
			case "cached_managed_cert", "cached_unmanaged_cert":
				data["identifier"] = strings.Join(data["sans"].([]string), ",")
				CertEvent("cached", data)
			}
		}
		return nil
	}

	cache := certmagic.NewCache(certmagic.CacheOptions{
		GetConfigForCert: func(cert certmagic.Certificate) (*certmagic.Config, error) {
			return &config, nil
		},
	})

	return certmagic.New(cache, config)

}
