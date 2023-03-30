package psutil

import (
	"net"
	"strings"

	"tdp-cloud/helper/request"
)

// 公网 IP

var publicIpv4 string
var publicIpv6 string

func PublicAddress(force bool) (string, string) {

	if force || (publicIpv4 == "" && publicIpv6 == "") {
		v4 := request.TimingGet("http://ipv4.rehi.org/ip", request.H{}, 10)
		v6 := request.TimingGet("http://ipv6.rehi.org/ip", request.H{}, 10)
		publicIpv4 = strings.TrimSpace(v4)
		publicIpv6 = strings.TrimSpace(v6)
	}

	return publicIpv4, publicIpv6

}

// 设备 IP

func InterfaceAddrs(name string) ([]string, []string) {

	ipv4 := []string{}
	ipv6 := []string{}

	addrs := []net.Addr{}

	if name == "" {
		if list, _ := net.InterfaceAddrs(); list != nil {
			addrs = list
		}
	} else {
		if ift, _ := net.InterfaceByName(name); ift != nil {
			if list, _ := ift.Addrs(); list != nil {
				addrs = list
			}
		}
	}

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
