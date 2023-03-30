package psutil

import (
	"strings"

	"tdp-cloud/helper/request"
)

// 云实例 Id

func CloudInstanceId() string {

	var url string
	var res string

	// alibaba
	url = "http://100.100.100.200/latest/meta-data/instance-id"
	res = request.TimingGet(url, request.H{}, 3)
	if res != "" {
		return strings.TrimSpace(res)
	}

	// tencent
	url = "http://metadata.tencentyun.com/latest/meta-data/instance-id"
	res = request.TimingGet(url, request.H{}, 3)
	if res != "" {
		return strings.TrimSpace(res)
	}

	// aws baidu huawei
	url = "http://169.254.169.254/latest/meta-data/instance-id"
	res = request.TimingGet(url, request.H{}, 3)
	if res != "" {
		return strings.TrimSpace(res)
	}

	// azure
	url = "http://169.254.169.254/metadata/instance/compute/vmId?api-version=2021-01-01"
	res = request.TimingGet(url, request.H{"Metadata": "true"}, 3)
	if res != "" {
		return strings.TrimSpace(res)
	}

	// google
	url = "http://metadata.google.internal/computeMetadata/v1/instance/id"
	res = request.TimingGet(url, request.H{"Metadata-Flavor": "Google"}, 3)
	if res != "" {
		return strings.TrimSpace(res)
	}

	// digitalocean
	url = "http://169.254.169.254/metadata/v1/id"
	res = request.TimingGet(url, request.H{}, 3)
	if res != "" {
		return strings.TrimSpace(res)
	}

	return res

}
