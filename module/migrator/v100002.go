package migrator

import (
	"tdp-cloud/module/model/config"
)

func v100002() error {

	if isMigrated("v100002") {
		return nil
	}

	if err := v100002AddConfig(); err != nil {
		return err
	}

	return addMigration("v100002", "添加默认参数")

}

func v100002AddConfig() error {

	list := []config.CreateParam{
		{
			Name:        "Registrable",
			Value:       "true",
			Type:        "bool",
			Module:      "system",
			Description: "允许注册",
		},
		{
			Name:        "Copytext",
			Value:       "",
			Type:        "string",
			Module:      "system",
			Description: "版权单位",
		},
		{
			Name:        "Copylink",
			Value:       "",
			Type:        "string",
			Module:      "system",
			Description: "版权链接",
		},
		{
			Name:        "Analytics",
			Value:       "",
			Type:        "text",
			Module:      "system",
			Description: "统计代码",
		},
		{
			Name:        "IcpCode",
			Value:       "",
			Type:        "string",
			Module:      "system",
			Description: "ICP 备案",
		},
	}

	for _, item := range list {
		_, err := config.Create(&item)
		if err != nil {
			return err
		}
	}

	return nil

}
