package migrator

import (
	"github.com/google/uuid"

	"tdp-cloud/core/dborm/config"
)

func v00002() error {

	return config.Create(&config.CreateParam{
		Key:   "AgentKey",
		Value: uuid.NewString(),
	})

}
