package alibaba

import (
	"errors"
	"strings"

	"github.com/mitchellh/mapstructure"
)

var endpointData = map[string]string{}

func solveEndpoint(rq *ReqeustParam) (string, error) {

	if rq.RegionId == "" {
		return rq.Service + ".aliyuncs.com", nil
	}

	// 从缓存返回

	key := rq.RegionId + rq.Service

	if endpointData[key] != "" {
		return endpointData[key], nil
	}

	// 从服务器获取

	if ep, err := requestEndpoint(rq); err == nil {
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

func requestEndpoint(rq *ReqeustParam) (*EndpointItem, error) {

	item := &EndpointItem{}

	// 从接口请求数据

	resp, err := newClient(&ReqeustParam{
		SecretId:  rq.SecretId,
		SecretKey: rq.SecretKey,
		Version:   "2015-06-12",
		Action:    "DescribeEndpoints",
		Query: map[string]string{
			"Id":          rq.RegionId,
			"ServiceCode": strings.ToLower(rq.Service),
			"Type":        "openAPI",
		},
		Endpoint: "location-readonly.aliyuncs.com",
	})

	if err != nil {
		return item, getSDKError(err)
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
