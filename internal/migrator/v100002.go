package migrator

import (
	"strings"
	"tdp-cloud/internal/dborm/task_script"
)

func v100002() error {

	if isMigrated("v100002") {
		return nil
	}

	if err := v100002AddScript(); err != nil {
		return err
	}

	return addMigration("v100002", "添加任务脚本")

}

func v100002AddScript() error {

	content := `
#!/bin/sh
TDP_EXEC_ARGS="--remote {{TDP_WSURL}}"
wget -qO- http://tdp.icu/worker-linux | sh -
`
	_, err := task_script.Create(&task_script.CreateParam{
		UserId:        0,
		Name:          "安装 TDP Cloud Worker 服务",
		Username:      "root",
		Description:   `变量表\n 服务端地址 {{TDP_WSURL}}`,
		Content:       strings.TrimLeft(content, "\n"),
		CommandType:   "shell",
		WorkDirectory: "/root",
		Timeout:       300,
	})

	return err

}
