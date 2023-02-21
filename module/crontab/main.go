package crontab

import (
	cron "github.com/robfig/cron/v3"
)

func Create() {

	cron.New(cron.WithSeconds())

}
