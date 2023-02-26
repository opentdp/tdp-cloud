package certmagic

import (
	"github.com/caddyserver/certmagic"
	"github.com/libdns/alidns"
	"github.com/libdns/cloudflare"
	"github.com/libdns/tencentcloud"
	"github.com/spf13/viper"
)

func newIssuer(rq *Params) *certmagic.ACMEIssuer {

	issuer := &certmagic.ACMEIssuer{
		Agreed:                  true,
		DisableHTTPChallenge:    true,
		DisableTLSALPNChallenge: true,
		Email:                   rq.Email,
	}

	if viper.GetBool("debug") {
		rq.CaType = "debug" //调试模式强制重写
	}

	switch rq.CaType {
	case "zerossl":
		issuer.CA = certmagic.ZeroSSLProductionCA
	case "letsencrypt":
		issuer.CA = certmagic.LetsEncryptProductionCA
	default:
		issuer.CA = certmagic.LetsEncryptStagingCA
	}

	switch rq.Provider {
	case "alibaba":
		issuer.DNS01Solver = &certmagic.DNS01Solver{
			DNSProvider: &alidns.Provider{
				AccKeyID:     rq.SecretId,
				AccKeySecret: rq.SecretKey,
			},
		}
	case "cloudflare":
		issuer.DNS01Solver = &certmagic.DNS01Solver{
			DNSProvider: &cloudflare.Provider{
				APIToken: rq.SecretKey,
			},
		}
	case "tencent":
		issuer.DNS01Solver = &certmagic.DNS01Solver{
			DNSProvider: &tencentcloud.Provider{
				SecretId:  rq.SecretId,
				SecretKey: rq.SecretKey,
			},
		}
	}

	return issuer

}
