package worker

import (
	"bufio"
	"context"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/mitchellh/mapstructure"

	"tdp-cloud/helper/strings"
	"tdp-cloud/module/workhub"
)

type ExecPayload = workhub.ExecPayload

func (pod *RecvPod) Exec(rs *SocketData) error {

	var err error
	var ret string

	var data *ExecPayload

	err = mapstructure.Decode(rs.Payload, &data)

	log.Println("Exec:recv", data.Name)

	if err == nil {
		switch data.CommandType {
		case "CMD":
			ret, err = cmdScript(data)
		case "POWERSHELL":
			ret, err = ps1Script(data)
		case "SHELL":
			ret, err = shellScript(data)
		}
	}

	rq := &SocketData{
		Method:  "Exec:resp",
		TaskId:  rs.TaskId,
		Success: err == nil,
		Payload: map[string]string{
			"Output": ret,
		},
	}

	if err := pod.Write(rq); err != nil {
		log.Println("Exec:error ", err)
		return err
	}

	return err

}

/////

func cmdScript(data *ExecPayload) (string, error) {

	tf, err := ioutil.TempFile(os.TempDir(), "tdp-*.bat")

	if err != nil {
		return "", errors.New("创建临时文件失败")
	}

	defer os.Remove(tf.Name())

	_, err = tf.WriteString(data.Content)

	if err != nil {
		return "", errors.New("写入临时文件失败")
	}

	name := "cmd.exe"
	params := []string{"/c", "CALL", tf.Name()}

	return execCommand(name, params, data.Timeout)

}

func ps1Script(data *ExecPayload) (string, error) {

	tf, err := ioutil.TempFile(os.TempDir(), "tdp-*.ps1")

	if err != nil {
		return "", errors.New("创建临时文件失败")
	}

	defer os.Remove(tf.Name())

	_, err = tf.WriteString(data.Content)

	if err != nil {
		return "", errors.New("写入临时文件失败")
	}

	name := "powershell.exe"
	params := []string{"-File", tf.Name()}

	return execCommand(name, params, data.Timeout)

}

func shellScript(data *ExecPayload) (string, error) {

	tf, err := ioutil.TempFile(os.TempDir(), "tdp-*")

	if err != nil {
		return "", errors.New("创建临时文件失败")
	}

	defer os.Remove(tf.Name())

	_, err = tf.WriteString(data.Content)

	if err != nil {
		return "", errors.New("写入临时文件失败")
	}

	tf.Chmod(0755)

	name := tf.Name()
	params := []string{}

	return execCommand(name, params, data.Timeout)

}

func execCommand(name string, params []string, timeout uint) (string, error) {

	var ret string

	// 设置上下文

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(timeout))

	cmd := exec.CommandContext(ctx, name, params...)

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

		if runtime.GOOS == "windows" {
			ret = strings.Gb18030ToUtf8(ret)
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
