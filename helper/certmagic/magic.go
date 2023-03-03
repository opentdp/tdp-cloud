package certmagic

import (
	"context"

	"github.com/caddyserver/certmagic"

	"tdp-cloud/cmd/args"
	"tdp-cloud/helper/logman"
)

func newMagic(iss certmagic.ACMEIssuer) *certmagic.Config {

	config := certmagic.Config{
		Logger: logman.Named("cert.magic"),
		Storage: &certmagic.FileStorage{
			Path: args.Dataset.Dir + "/certmagic",
		},
		OnEvent: OnEvent,
	}

	config.Issuers = []certmagic.Issuer{
		certmagic.NewACMEIssuer(&config, iss),
	}

	cache := certmagic.NewCache(certmagic.CacheOptions{
		GetConfigForCert: func(cert certmagic.Certificate) (*certmagic.Config, error) {
			return &config, nil
		},
		Logger: logman.Named("cert.cache"),
	})

	return certmagic.New(cache, config)

}

func OnEvent(ctx context.Context, evt string, data map[string]any) error {

	evtlog := logman.Named("cert.event").Sugar()

	evtlog.Warnf("Certmagic Event: %s with data: %v\n", evt, data)
	return nil

}
