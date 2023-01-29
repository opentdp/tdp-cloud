package psutil

import (
	"log"
	"strings"

	"tdp-cloud/helper/request"
)

var ipAddress string

func getIpAddress(f bool) string {

	if f || ipAddress == "" {
		log.Println("fetch public ip from http://ipip.rpc.im")
		body, err := request.Get("http://ipip.rpc.im/ip", nil)
		if err == nil {
			ipAddress = strings.TrimSpace(body)
		}
	}

	return ipAddress

}
