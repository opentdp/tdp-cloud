package certmagic

import (
	"context"
	"errors"
	"strings"

	"github.com/caddyserver/certmagic"
	"github.com/forgoer/openssl"
)

var magicPool = map[string]*certmagic.Config{}

func Manage(rq *ReqeustParam) error {

	skey := openssl.Md5ToString(rq.Email + rq.SecretKey + rq.CaType)

	magic, ok := magicPool[skey]

	if !ok {
		magic = newMagic(newIssuer(rq))
		magicPool[skey] = magic
	}

	magicPool[rq.Domain] = magic

	domains := strings.Split(rq.Domain, ",")
	return magic.ManageAsync(context.Background(), domains)

}

func Unmanage(domain string) {

	magic, ok := magicPool[domain]
	domains := strings.Split(domain, ",")

	if ok {
		magic.Unmanage(domains)
		delete(magicPool, domain)
	}

}

func CertDetail(domain string) (*Certificate, error) {

	cert := &Certificate{}

	magic, ok := magicPool[domain]

	if !ok {
		return cert, errors.New("任务不存在或已被删除")
	}

	crt, err := magic.CacheManagedCertificate(context.Background(), domain)

	if err != nil {
		return cert, err
	}

	pk, err := certmagic.PEMEncodePrivateKey(crt.Certificate.PrivateKey)

	cert.Names = crt.Names
	cert.NotAfter = crt.Leaf.NotAfter.Unix()
	cert.NotBefore = crt.Leaf.NotBefore.Unix()
	cert.Certificate = crt.Certificate.Certificate
	cert.PrivateKey = pk

	cert.Issuer = map[string]any{
		"CommonName":   crt.Leaf.Issuer.CommonName,
		"Organization": crt.Leaf.Issuer.Organization[0],
		"Country":      crt.Leaf.Issuer.Country[0],
	}

	return cert, err

}
