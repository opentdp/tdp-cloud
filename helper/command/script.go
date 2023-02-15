package command

import (
	"context"
	"errors"
	"io/ioutil"
	"os/exec"
	"runtime"
	"time"
)

func newScript(code string, ext string) (string, error) {

	tf, err := ioutil.TempFile("", "tdp-*."+ext)

	if err != nil {
		return "", errors.New("创建临时文件失败")
	}

	defer tf.Close()

	if _, err = tf.WriteString(code); err != nil {
		return "", errors.New("写入临时文件失败")
	}

	if runtime.GOOS != "windows" {
		tf.Chmod(0755)
	}

	return tf.Name(), nil

}

func execScript(bin string, arg []string, data *ExecPayload) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(data.Timeout)*time.Second)

	defer cancel()

	cmd := exec.CommandContext(ctx, bin, arg...)

	if data.WorkDirectory != "" {
		cmd.Dir = data.WorkDirectory
	}

	ret, err := cmd.CombinedOutput()

	return string(ret), err

}
