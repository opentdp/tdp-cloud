package alibaba

import (
	ac "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	au "github.com/alibabacloud-go/openapi-util/service"
	as "github.com/alibabacloud-go/tea-utils/v2/service"
	at "github.com/alibabacloud-go/tea/tea"
)

func Request(rp *Params) (any, error) {

	if ep, err := solveEndpoint(rp); ep != "" {
		rp.Endpoint = ep
	} else {
		return nil, err
	}

	config := &ac.Config{
		AccessKeyId:     &rp.SecretId,
		AccessKeySecret: &rp.SecretKey,
		RegionId:        &rp.RegionId,
		Endpoint:        &rp.Endpoint,
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

	if client, err := ac.NewClient(config); err == nil {
		return client.CallApi(params, request, runtime)
	} else {
		return nil, err
	}

}
