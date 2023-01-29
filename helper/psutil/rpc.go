package psutil

import (
	"strings"

	"github.com/imroc/req/v3"
)

var ipAddress string

func getIpAddress(f bool) string {

	if f || ipAddress == "" {
		resp, err := req.Get("http://ipip.rpc.im/ip")
		if err == nil {
			ipAddress = strings.TrimSpace(resp.String())
		}
	}

	return ipAddress

}
