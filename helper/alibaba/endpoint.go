package alibaba

import (
	"errors"
	"strings"

	"github.com/mitchellh/mapstructure"
)

var endpointData = map[string]string{}

func solveEndpoint(rp *Params) (string, error) {

	if rp.RegionId == "" {
		return rp.Service + ".aliyuncs.com", nil
	}

	// 从缓存返回

	key := rp.RegionId + rp.Service

	if endpointData[key] != "" {
		return endpointData[key], nil
	}

	// 从服务器获取

	if ep, err := requestEndpoint(rp); err == nil {
		endpointData[key] = ep.Endpoint
	} else {
		return "", err
	}

	// 校验缓存并返回

	if endpointData[key] != "" {
		return endpointData[key], nil
	}

	return "", errors.New("获取 Endpoint 失败")

}

func requestEndpoint(rp *Params) (*EndpointItem, error) {

	item := &EndpointItem{}

	// 从接口请求数据

	resp, err := newClient(&Params{
		SecretId:  rp.SecretId,
		SecretKey: rp.SecretKey,
		Version:   "2015-06-12",
		Action:    "DescribeEndpoints",
		Query: map[string]string{
			"Id":          rp.RegionId,
			"ServiceCode": strings.ToLower(rp.Service),
			"Type":        "openAPI",
		},
		Endpoint: "location-readonly.aliyuncs.com",
	})

	if err != nil {
		return item, err
	}

	// 尝试解析数据

	data := &EndpointBody{}
	err = mapstructure.Decode(resp["body"], data)

	if err != nil {
		return item, err
	}

	if len(data.Endpoints.Endpoint) > 0 {
		item = &data.Endpoints.Endpoint[0]
	} else {
		err = errors.New("获取 Endpoint 失败")
	}

	return item, err

}
