package migrator

import (
	"log"
	"strings"

	"tdp-cloud/core/dborm/config"
)

var Versions = "100000"

func Start() {

	if err := doMigrate(); err != nil {
		log.Panic(err)
	}

}

func doMigrate() error {

	getMigration()

	if err := v100001(); err != nil {
		return err
	}

	if err := v100002(); err != nil {
		return err
	}

	Versions = "" // 释放资源

	return nil

}

func isMigrated(v string) bool {

	return strings.Contains(Versions, v)

}

func getMigration() string {

	item, err := config.Fetch("Migration")

	if err == nil && item.Value != "" {
		Versions = item.Value
	}

	return Versions

}

func addMigration(v string) error {

	if isMigrated(v) {
		return nil
	}

	Versions += ":" + v

	return config.Update(&config.UpdateParam{
		Name:        "Migration",
		Value:       Versions,
		Module:      "System",
		Description: "自动迁移记录",
	})

}
