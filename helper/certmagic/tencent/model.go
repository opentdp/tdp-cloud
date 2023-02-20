package tencent

type Provider struct {
	SecretId  string
	SecretKey string
}

type DescribeRecordListResponse struct {
	RecordList []struct {
		RecordId int
		Value    string
		Name     string
		Type     string
		TTL      int
		MX       int
	}
}

type CreateRecordResponse struct {
	RecordId int
}
