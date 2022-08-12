package migrator

import (
	"strconv"

	"tdp-cloud/core/dborm/config"
)

var Version = 0

func Migrate() {

	getVersion()
	defer setVersion()

	if v00001() != nil {
		panic("Migration to v100001 failed")
	}

	if Version < 100002 {
		if v00002() != nil {
			panic("Migration to v100002 failed")
		}
		Version = 100002
	}

}

func getVersion() {

	item, err := config.Fetch("migrate")

	if err == nil && item.Value != "" {
		v, _ := strconv.Atoi(item.Value)
		Version = v
	}

}

func setVersion() {

	v := strconv.Itoa(Version)

	config.Update(&config.UpdateParam{
		Key: "migrate", Value: v,
	})

}
