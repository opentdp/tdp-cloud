package command

import (
	"bufio"
	"context"
	"errors"
	"io"
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

func execScript(name string, arg []string, timeout uint) (string, error) {

	var ret string

	// 设置上下文

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(timeout))

	cmd := exec.CommandContext(ctx, name, arg...)

	defer cancel()

	// 持续读取输出

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		return "", err
	}

	defer stdout.Close()

	go func() {
		reader := bufio.NewReader(stdout)

		for {
			str, err := reader.ReadString('\n')
			if err != nil || err == io.EOF {
				break
			}
			ret += str
		}
	}()

	// 开始执行命令

	if err := cmd.Start(); err != nil {
		return ret, err
	}

	// 等待命令结束

	if err := cmd.Wait(); err != nil {
		return ret, err
	}

	return ret, err

}
