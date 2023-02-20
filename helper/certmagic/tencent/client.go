package tencent

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/libdns/libdns"
	"github.com/mitchellh/mapstructure"

	"tdp-cloud/helper/tencent"
)

func (p *Provider) describeRecordList(ctx context.Context, zone string) ([]libdns.Record, error) {

	list := []libdns.Record{}

	payload := map[string]any{
		"Domain": strings.Trim(zone, "."),
	}

	res, err := p.doRequest("DescribeRecordList", payload)
	if err != nil {
		return list, err
	}

	data := DescribeRecordListResponse{}
	err = mapstructure.Decode(res, &data)
	if err != nil {
		return list, err
	}

	for _, record := range data.RecordList {
		list = append(list, libdns.Record{
			ID:    strconv.Itoa(record.RecordId),
			Type:  record.Type,
			Name:  record.Name,
			Value: record.Value,
			TTL:   time.Duration(record.TTL) * time.Second,
		})
	}

	return list, err

}

func (p *Provider) createRecord(ctx context.Context, zone string, record libdns.Record) (string, error) {

	payload := map[string]any{
		"Domain":     strings.Trim(zone, "."),
		"SubDomain":  record.Name,
		"RecordType": record.Type,
		"RecordLine": "默认",
		"Value":      record.Value,
	}

	res, err := p.doRequest("CreateRecord", payload)
	if err != nil {
		return "", err
	}

	data := CreateRecordResponse{}
	err = mapstructure.Decode(res, &data)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(data.RecordId), nil

}

func (p *Provider) modifyRecord(ctx context.Context, zone string, record libdns.Record) error {

	recordId, _ := strconv.Atoi(record.ID)

	payload := map[string]any{
		"Domain":     strings.Trim(zone, "."),
		"SubDomain":  record.Name,
		"RecordType": record.Type,
		"RecordLine": "默认",
		"Value":      record.Value,
		"RecordId":   recordId,
	}

	_, err := p.doRequest("ModifyRecord", payload)

	return err

}

func (p *Provider) deleteRecord(ctx context.Context, zone string, record libdns.Record) error {

	recordId, _ := strconv.Atoi(record.ID)

	payload := map[string]any{
		"Domain":   strings.Trim(zone, "."),
		"RecordId": recordId,
	}

	_, err := p.doRequest("DeleteRecord", payload)

	return err

}

func (p *Provider) doRequest(action string, payload any) (any, error) {

	params := &tencent.Params{
		SecretId:  p.SecretId,
		SecretKey: p.SecretKey,
		Service:   "dnspod",
		Version:   "2021-03-23",
		Action:    action,
		Payload:   payload,
	}

	return tencent.Request(params)

}
