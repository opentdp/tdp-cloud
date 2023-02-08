package command

import (
	"errors"
	"os"
	"runtime"
	"tdp-cloud/helper/strutil"
)

type ExecPayload struct {
	Name          string
	CommandType   string
	Username      string
	WorkDirectory string
	Content       string
	Timeout       uint
}

func Exec(data *ExecPayload) (string, error) {

	var err error
	var tmp string
	var cmd string
	var arg []string

	switch data.CommandType {
	case "CMD":
		tmp, err = newScript(data.Content, "bat")
		arg = []string{"/c", "CALL", tmp}
		cmd = "cmd.exe"
	case "POWERSHELL":
		tmp, err = newScript(data.Content, "ps1")
		arg = []string{"-File", tmp}
		cmd = "powershell.exe"
	case "SHELL":
		tmp, err = newScript(data.Content, "sh")
		arg = []string{}
		cmd = tmp
	default:
		return "", errors.New("不支持此类脚本")
	}

	if err != nil {
		return "", err
	}

	defer os.Remove(tmp)

	ret, err := execScript(cmd, arg, data.Timeout)

	if runtime.GOOS == "windows" {
		ret = strutil.Gb18030ToUtf8(ret)
	}

	return ret, err

}
