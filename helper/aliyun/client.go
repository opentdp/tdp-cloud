package aliyun

import (
	ac "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	au "github.com/alibabacloud-go/openapi-util/service"
	as "github.com/alibabacloud-go/tea-utils/v2/service"
	at "github.com/alibabacloud-go/tea/tea"
)

func Request(rp *Params) (any, error) {

	// 分地域接入
	endpoint, err := getEndpoint(rp.Service, rp.Region)

	if err != nil {
		return nil, err
	}

	config := &ac.Config{
		AccessKeyId:     &rp.SecretId,
		AccessKeySecret: &rp.SecretKey,
		RegionId:        &rp.Region,
		Endpoint:        &endpoint,
	}

	client, err := ac.NewClient(config)
	if err != nil {
		return nil, err
	}

	params := &ac.Params{
		Action:      at.String(rp.Action),
		Version:     at.String(rp.Version),
		Protocol:    at.String("HTTPS"),
		Pathname:    at.String("/"),
		Method:      at.String("POST"),
		AuthType:    at.String("AK"),
		Style:       at.String("RPC"),
		ReqBodyType: at.String("json"),
		BodyType:    at.String("json"),
	}

	request := &ac.OpenApiRequest{
		Query: au.Query(rp.Query),
		Body:  au.Query(rp.Payload),
	}

	runtime := &as.RuntimeOptions{}

	return client.CallApi(params, request, runtime)

}

// rs, er := aliyun.Request(&aliyun.Params{
// 	SecretId:  "LTAI5tEmFdkxkudYqBSZZqnf",
// 	SecretKey: "os30YWmM2wfC2pnOazSZ87NaXt6NpM",
// 	Service:   "ecs",
// 	Version:   "2014-05-26",
// 	Region:    "cn-hangzhou",
// 	Action:    "DescribeAvailableResource",
// 	Query:     map[string]string{
// 		"RegionId":            "cn-hangzhou",
// 		"DestinationResource": "Zone",
// 	},
// })

// fmt.Println(rs, er)
