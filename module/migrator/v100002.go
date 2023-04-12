package migrator

import (
	"tdp-cloud/model/config"
)

func v100002() error {

	if isMigrated("v100002") {
		return nil
	}

	if err := v100002AddConfig(); err != nil {
		return err
	}

	return addMigration("v100002", "添加系统参数")

}

func v100002AddConfig() error {

	items := []config.CreateParam{
		{
			Name:        "Registrable",
			Value:       "true",
			Type:        "bool",
			Module:      "system",
			Description: "允许注册",
		},
	}

	for _, item := range items {
		if _, err := config.Create(&item); err != nil {
			return err
		}
	}

	return nil

}
