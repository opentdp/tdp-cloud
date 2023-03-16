package psutil

import (
	"net"
	"strings"

	"tdp-cloud/helper/request"
)

// 设备 IP

func InterfaceAddrs() ([]string, []string) {

	ipv4 := []string{}
	ipv6 := []string{}

	addrs, _ := net.InterfaceAddrs()

	for _, ip := range addrs {
		if ipnet, ok := ip.(*net.IPNet); ok && ipnet.IP.IsGlobalUnicast() {
			if ipnet.IP.To4() != nil {
				ipv4 = append(ipv4, ipnet.IP.String())
			} else {
				ipv6 = append(ipv6, ipnet.IP.String())
			}
		}
	}

	return ipv4, ipv6

}

// 公网 IP

func PublicAddress() ([]string, []string) {

	v4 := request.SimpleGet("http://ipv4.rehi.org/ip", request.H{}, 10)
	v6 := request.SimpleGet("http://ipv6.rehi.org/ip", request.H{}, 10)

	ipv4 := strings.Split(strings.TrimSpace(v4), ",")
	ipv6 := strings.Split(strings.TrimSpace(v6), ",")

	return ipv4, ipv6

}

// 云实例 Id

func CloudInstanceId() string {

	var url string
	var res string

	// alibaba
	url = "http://100.100.100.200/latest/meta-data/instance-id"
	res = request.SimpleGet(url, request.H{}, 3)
	if res != "" {
		return strings.TrimSpace(res)
	}

	// tencent
	url = "http://metadata.tencentyun.com/latest/meta-data/instance-id"
	res = request.SimpleGet(url, request.H{}, 3)
	if res != "" {
		return strings.TrimSpace(res)
	}

	// aws baidu huawei
	url = "http://169.254.169.254/latest/meta-data/instance-id"
	res = request.SimpleGet(url, request.H{}, 3)
	if res != "" {
		return strings.TrimSpace(res)
	}

	// azure
	url = "http://169.254.169.254/metadata/instance/compute/vmId?api-version=2021-01-01"
	res = request.SimpleGet(url, request.H{"Metadata": "true"}, 3)
	if res != "" {
		return strings.TrimSpace(res)
	}

	// google
	url = "http://metadata.google.internal/computeMetadata/v1/instance/id"
	res = request.SimpleGet(url, request.H{"Metadata-Flavor": "Google"}, 3)
	if res != "" {
		return strings.TrimSpace(res)
	}

	// digitalocean
	url = "http://169.254.169.254/metadata/v1/id"
	res = request.SimpleGet(url, request.H{}, 3)
	if res != "" {
		return strings.TrimSpace(res)
	}

	return res

}
