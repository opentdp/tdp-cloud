package psutil

import (
	"log"

	"github.com/imroc/req/v3"
)

func getIpAddress() string {

	resp, err := req.Get("https://ipip.rpc.im/ip")
	log.Print(resp, err)

	return string(resp.Bytes())

}
