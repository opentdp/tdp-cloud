package tencent

import (
	"context"

	"github.com/libdns/libdns"
)

func (p *Provider) GetRecords(ctx context.Context, zone string) ([]libdns.Record, error) {

	return p.describeRecordList(ctx, zone)

}

func (p *Provider) AppendRecords(ctx context.Context, zone string, records []libdns.Record) ([]libdns.Record, error) {

	for k, record := range records {
		if id, err := p.createRecord(ctx, zone, record); err != nil {
			return records, err
		} else {
			records[k].ID = id
		}
	}

	return records, nil

}

func (p *Provider) SetRecords(ctx context.Context, zone string, records []libdns.Record) ([]libdns.Record, error) {

	for _, record := range records {
		if err := p.modifyRecord(ctx, zone, record); err != nil {
			return nil, err
		}
	}

	return records, nil

}

func (p *Provider) DeleteRecords(ctx context.Context, zone string, records []libdns.Record) ([]libdns.Record, error) {

	for _, record := range records {
		if err := p.deleteRecord(ctx, zone, record); err != nil {
			return nil, err
		}
	}

	return records, nil

}

// Interface guards

var (
	_ libdns.RecordGetter   = (*Provider)(nil)
	_ libdns.RecordAppender = (*Provider)(nil)
	_ libdns.RecordSetter   = (*Provider)(nil)
	_ libdns.RecordDeleter  = (*Provider)(nil)
)
