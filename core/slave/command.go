package slave

import (
	"bufio"
	"context"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"sync"
	"time"

	"tdp-cloud/core/helper"
	"tdp-cloud/core/socket"
)

type CommandPayload struct {
	Content          string `binding:"required"`
	Username         string `binding:"required"`
	CommandType      string `binding:"required"`
	WorkingDirectory string `binding:"required"`
	Timeout          uint   `binding:"required"`
}

func RunCommand(pod *socket.JsonPod, data *CommandPayload) error {

	var err error
	var ret string

	switch data.CommandType {
	case "CMD":
		ret, err = cmdScript(data)
	case "POWERSHELL":
		ret, err = ps1Script(data)
	case "SHELL":
		ret, err = shellScript(data)
	}

	v := &SocketData{
		Action:  "runCommand",
		Method:  "response",
		Payload: ret,
		Error:   err,
	}

	if err := pod.Write(v); err != nil {
		log.Println("[RunCommand] ", err)
		return err
	}

	return err

}

/////

func cmdScript(data *CommandPayload) (string, error) {

	tf, err := ioutil.TempFile(os.TempDir(), "tat-*.bat")

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

func ps1Script(data *CommandPayload) (string, error) {

	tf, err := ioutil.TempFile(os.TempDir(), "tat-*.ps1")

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

func shellScript(data *CommandPayload) (string, error) {

	tf, err := ioutil.TempFile(os.TempDir(), "tat-*")

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

	// 超时时间

	otime := time.Duration(timeout) * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), otime)

	defer cancel()

	// 执行命令

	cmd := exec.CommandContext(ctx, name, params...)

	// 捕获输出

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return ret, err
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		reader := bufio.NewReader(stdout)
		for {
			rs, err := reader.ReadString('\n')
			if err != nil || err == io.EOF {
				return
			}
			ret += helper.Byte2String([]byte(rs), "GB18030")
		}
	}()

	err = cmd.Start()
	wg.Wait()

	return ret, err

}
