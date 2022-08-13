package migrator

import (
	"github.com/google/uuid"

	"tdp-cloud/core/dborm/config"
)

func v100002() error {

	if isMigrated("100002") {
		return nil
	}

	if err := newMigration(); err != nil {
		return err
	}

	if err := newAgentKey(); err != nil {
		return err
	}

	addMigration("100002")

	return nil

}

func newAgentKey() error {

	item, err := config.Fetch("AgentKey")

	if err == nil && item.Id > 0 {
		return nil
	}

	return config.Create(&config.CreateParam{
		Key:         "AgentKey",
		Value:       uuid.NewString(),
		Module:      "System",
		Description: "客户端注册密钥",
	})

}

func newMigration() error {

	item, err := config.Fetch("Migration")

	if err == nil && item.Id > 0 {
		return nil
	}

	Versions = "100001"

	return config.Create(&config.CreateParam{
		Key:         "Migration",
		Value:       Versions,
		Module:      "System",
		Description: "自动迁移记录",
	})

}
