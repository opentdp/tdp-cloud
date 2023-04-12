package migrator

import (
	"tdp-cloud/model/config"
)

func v100003() error {

	if isMigrated("v100003") {
		return nil
	}

	if err := v100003AddConfig(); err != nil {
		return err
	}

	return addMigration("v100003", "添加前端参数")

}

func v100003AddConfig() error {

	items := []config.CreateParam{
		{
			Name:        "SiteName",
			Value:       "",
			Type:        "string",
			Module:      "front",
			Description: "网站名称",
		},
		{
			Name:        "SiteIcon",
			Value:       "",
			Type:        "upload",
			Module:      "front",
			Description: "网站 Icon",
		},
		{
			Name:        "SiteLogo",
			Value:       "",
			Type:        "upload",
			Module:      "front",
			Description: "网站 Logo",
		},
		{
			Name:        "Copytext",
			Value:       "",
			Type:        "string",
			Module:      "front",
			Description: "版权单位",
		},
		{
			Name:        "Copylink",
			Value:       "",
			Type:        "string",
			Module:      "front",
			Description: "版权链接",
		},
		{
			Name:        "IcpCode",
			Value:       "",
			Type:        "string",
			Module:      "front",
			Description: "ICP 备案",
		},
		{
			Name:        "Analytics",
			Value:       "",
			Type:        "text",
			Module:      "front",
			Description: "统计代码",
		},
	}

	for _, item := range items {
		if _, err := config.Create(&item); err != nil {
			return err
		}
	}

	return nil

}
