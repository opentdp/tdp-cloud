package psutil

import (
	"strings"

	"tdp-cloud/helper/request"
)

var ipAddress string

func getIpAddress(f bool) string {

	if f || ipAddress == "" {
		body, err := request.TextGet("http://ipip.rpc.im/ip", nil)
		if err == nil {
			ipAddress = strings.TrimSpace(body)
		}
	}

	return ipAddress

}
