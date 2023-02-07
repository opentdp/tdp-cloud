package alibaba

import (
	"encoding/json"
	"errors"
	"strings"

	ac "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	as "github.com/alibabacloud-go/tea-utils/v2/service"
	at "github.com/alibabacloud-go/tea/tea"
)

var endpointData = map[string]string{}

type endpointBody struct {
	Endpoints struct {
		Endpoint []struct {
			Id        string
			Endpoint  string
			Namespace string
			Protocols struct {
				Protocols []string
			}
			SerivceCode string
			Type        string
		}
	}
	RequestId string
	Success   bool
}

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

	resp, err := requestEndpoint(rp)

	if err != nil {
		return "", err
	}

	data := &endpointBody{}
	body := []byte(resp["body"].(string))

	if err := json.Unmarshal(body, data); err != nil {
		return "", err
	}

	// 将结果写入缓存

	if len(data.Endpoints.Endpoint) > 0 {
		endpointData[key] = data.Endpoints.Endpoint[0].Endpoint
	}

	// 校验缓存并返回

	if endpointData[key] != "" {
		return endpointData[key], nil
	}
	return "", errors.New("获取 Endpoint 失败")

}

func requestEndpoint(rp *Params) (map[string]interface{}, error) {

	config := &ac.Config{
		AccessKeyId:     &rp.SecretId,
		AccessKeySecret: &rp.SecretKey,
		Endpoint:        at.String("location-readonly.aliyuncs.com"),
	}

	params := &ac.Params{
		Action:      at.String("DescribeEndpoints"),
		Version:     at.String("2015-06-12"),
		Protocol:    at.String("HTTPS"),
		Pathname:    at.String("/"),
		Method:      at.String("GET"),
		AuthType:    at.String("AK"),
		Style:       at.String("RPC"),
		ReqBodyType: at.String("json"),
		BodyType:    at.String("string"),
	}

	request := &ac.OpenApiRequest{
		Query: map[string]*string{
			"Id":          &rp.RegionId,
			"ServiceCode": at.String(strings.ToLower(rp.Service)),
			"Type":        at.String("openAPI"),
		},
	}

	runtime := &as.RuntimeOptions{}

	if client, err := ac.NewClient(config); err == nil {
		return client.CallApi(params, request, runtime)
	} else {
		return nil, err
	}

}
