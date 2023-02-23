package certmagic

import (
	"github.com/caddyserver/certmagic"
	"github.com/libdns/alidns"
	"github.com/libdns/cloudflare"
	"github.com/spf13/viper"

	tencent "github.com/rehiy/libdns-tencentcloud"
)

type Params struct {
	Email     string
	Domain    string
	CaType    string
	Provider  string
	SecretId  string
	SecretKey string
}

func newIssuer(rp *Params) *certmagic.ACMEIssuer {

	issuer := &certmagic.ACMEIssuer{
		Agreed:                  true,
		DisableHTTPChallenge:    true,
		DisableTLSALPNChallenge: true,
		Email:                   rp.Email,
	}

	switch rp.CaType {
	case "zerossl":
		issuer.CA = certmagic.ZeroSSLProductionCA
	default:
		issuer.CA = certmagic.LetsEncryptProductionCA
	}

	if viper.GetBool("debug") { //调试模式强制重写
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
