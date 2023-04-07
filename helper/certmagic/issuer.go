package certmagic

import (
	"github.com/caddyserver/certmagic"
	"github.com/libdns/alidns"
	"github.com/libdns/cloudflare"
	"github.com/libdns/tencentcloud"
	"github.com/mholt/acmez/acme"
)

func newIssuer(rq *ReqeustParam) certmagic.ACMEIssuer {

	issuer := certmagic.ACMEIssuer{
		Agreed:                  true,
		DisableHTTPChallenge:    true,
		DisableTLSALPNChallenge: true,
		Email:                   rq.Email,
	}

	// Ref: https://github.com/acmesh-official/acme.sh/wiki/Server

	switch rq.CaType {
	case "letsencrypt":
		issuer.CA = certmagic.LetsEncryptProductionCA
	case "buypass":
		issuer.CA = "https://api.buypass.com/acme/directory"
	case "google":
		issuer.CA = "https://dv.acme-v02.api.pki.goog/directory"
	case "sslcom-ecc":
		issuer.CA = "https://acme.ssl.com/sslcom-dv-ecc"
	case "sslcom-rsa":
		issuer.CA = "https://acme.ssl.com/sslcom-dv-rsa"
	case "zerossl":
		issuer.CA = certmagic.ZeroSSLProductionCA
	default: //debug
		issuer.CA = certmagic.LetsEncryptStagingCA
	}

	if rq.EabKeyId != "" && rq.EabMacKey != "" {
		issuer.ExternalAccount = &acme.EAB{
			KeyID:  rq.EabKeyId,
			MACKey: rq.EabMacKey,
		}
	}

	// Ref: https://github.com/libdns/libdns

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
