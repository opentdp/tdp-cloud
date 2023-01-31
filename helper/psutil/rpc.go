package psutil

import (
	"strings"

	"tdp-cloud/helper/request"
)

var ipAddress string

func getIpAddress(f bool) string {

	if f || ipAddress == "" {
		body, err := request.TextGet("https://ipip.rehi.org/ip", nil)
		if err == nil {
			ipAddress = strings.TrimSpace(body)
		}
	}

	return ipAddress

}
