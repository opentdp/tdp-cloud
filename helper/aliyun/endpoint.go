package aliyun

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/jmespath/go-jmespath"
)

var _ *embed.FS

var initOnce sync.Once
var endpointConfig interface{}

//go:embed endpoint.json
var endpointsJson []byte

func getEndpoint(product, regionId string) (string, error) {

	regionalExpression := fmt.Sprintf("products[?code=='%s'].regional_endpoints", strings.ToLower(product))
	regionalData, err := jmespath.Search(regionalExpression, getEndpointConfig())

	if err != nil || regionalData == nil || len(regionalData.([]interface{})) == 0 {
		return "", err
	}

	endpointExpression := fmt.Sprintf("[0][?region=='%s'].endpoint", strings.ToLower(regionId))
	endpointData, err := jmespath.Search(endpointExpression, regionalData)

	if err != nil || endpointData == nil || len(endpointData.([]interface{})) == 0 {
		return "", err
	}

	return endpointData.([]interface{})[0].(string), nil

}

func getEndpointConfig() interface{} {

	initOnce.Do(func() {
		if err := json.Unmarshal(endpointsJson, &endpointConfig); err != nil {
			log.Fatalln("init endpoint config data failed.", err)
		}
	})

	return endpointConfig

}
