package dnspod

type DescribeDomainListRequest struct {
	Type    *string
	Offset  *int64
	Limit   *int64
	GroupId *int64
	Keyword *string
}
