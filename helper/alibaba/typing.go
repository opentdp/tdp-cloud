package alibaba

type ReqeustParam struct {
	Service   string `note:"产品名称"`
	Version   string `note:"接口版本"`
	Action    string `note:"接口名称"`
	Query     any    `note:"结构化数据"`
	Payload   any    `note:"结构化数据"`
	RegionId  string `note:"资源所在区域"`
	Endpoint  string `note:"指定接口域名"`
	SecretId  string `note:"访问密钥 Id"`
	SecretKey string `note:"访问密钥 Key"`
}

type ResponseData struct {
	Response any
}

//// Endpoint

type EndpointItem struct {
	Id        string
	Endpoint  string
	Namespace string
	Protocols struct {
		Protocols []string
	}
	SerivceCode string
	Type        string
}

type EndpointBody struct {
	Endpoints struct {
		Endpoint []EndpointItem
	}
	RequestId string
	Success   bool
}
