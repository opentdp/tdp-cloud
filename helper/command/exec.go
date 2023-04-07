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

	var (
		err error
		tmp string
		bin string
		arg []string
	)

	switch data.CommandType {
	case "BAT":
		tmp, err = newScript(data.Content, "bat")
		arg = []string{"/c", "CALL", tmp}
		bin = "cmd.exe"
	case "POWERSHELL":
		tmp, err = newScript(data.Content, "ps1")
		arg = []string{"-File", tmp}
		bin = "powershell.exe"
	case "SHELL":
		tmp, err = newScript(data.Content, "sh")
		arg = []string{}
		bin = tmp
	default:
		err = errors.New("不支持此类脚本")
	}

	if err != nil {
		return "", err
	}

	defer os.Remove(tmp)

	ret, err := execScript(bin, arg, data)

	if runtime.GOOS == "windows" {
		ret = strutil.Gb18030ToUtf8(ret)
	}

	return ret, err

}
