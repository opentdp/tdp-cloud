package psutil

import (
	"net"
	"strings"

	"tdp-cloud/helper/request"
)

// 内网 IP

func PrivateIpAddress() string {

	res := []string{}
	addrs, _ := net.InterfaceAddrs()

	for _, ip := range addrs {
		if ipnet, ok := ip.(*net.IPNet); ok && ipnet.IP.IsPrivate() {
			res = append(res, ipnet.IP.String())
		}
	}

	return strings.Join(res, ", ")

}

// 公网 IP

var publicIpAddr string

func PublicIpAddress(f bool) string {

	if f || publicIpAddr == "" {
		ip := request.SimpleGet("https://ipip.rehi.org/ip", 10)
		publicIpAddr = strings.TrimSpace(ip)
	}

	return publicIpAddr

}

// 云实例 Id

const alibabaUrl = "http://100.100.100.200/latest/meta-data"
const tencentUrl = "http://metadata.tencentyun.com/latest/meta-data"

func CloudInstanceId() string {

	var id string

	id = request.SimpleGet(alibabaUrl+`/instance-id`, 3)
	if id != "" {
		return strings.TrimSpace(id)
	}

	id = request.SimpleGet(tencentUrl+`/instance-id`, 3)
	if id != "" {
		return strings.TrimSpace(id)
	}

	return id

}
