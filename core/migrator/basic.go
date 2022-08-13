package migrator

import (
	"log"
	"strings"

	"tdp-cloud/core/dborm/config"
)

var Versions = ""

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

func getMigration() {

	item, err := config.Fetch("Migration")

	if err == nil && item.Value != "" {
		Versions = item.Value
	}

}

func addMigration(v string) {

	if isMigrated(v) {
		return
	}

	Versions += ":" + v

	config.Update(&config.UpdateParam{
		Key:         "Migration",
		Value:       Versions,
		Module:      "System",
		Description: "自动迁移记录",
	})

}
