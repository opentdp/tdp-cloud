package webssh

import (
	"io"
	"net/http"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/term"

	"tdp-cloud/helper/socket"
)

type ConnectParam struct {
	Request *http.Request
	Writer  http.ResponseWriter
	Option  *SSHClientOption
}

func Connect(p *ConnectParam) error {

	pod, err := socket.NewIOPod(p.Writer, p.Request)

	if err != nil {
		return err
	}

	defer pod.Close()

	// 创建 SSH 连接

	client, err := NewSSHClient(p.Option)

	if err != nil {
		pod.Write([]byte("> " + err.Error() + "\r\n"))
		return err
	}

	defer client.Close()

	// 转发 SSH 会话

	quit := make(chan bool, 1)
	go sshProxy(client, pod, quit)
	<-quit

	return nil

}

func sshProxy(client *ssh.Client, pod *socket.IOPod, quit chan bool) {

	defer func() {
		quit <- true
	}()

	rw := io.ReadWriter(pod)

	session, err := client.NewSession()

	if err != nil {
		rw.Write([]byte(err.Error() + "\r\n"))
		return
	}

	defer session.Close()

	// 客户端断开时清理资源
	pod.OnClose(session.Close)

	session.Stdin = rw
	session.Stdout = rw
	session.Stderr = rw

	fd := int(os.Stdin.Fd())
	width, height, _ := term.GetSize(fd)

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	if err := session.RequestPty("xterm", width, height, modes); err != nil {
		rw.Write([]byte(err.Error() + "\r\n"))
	}

	if err := session.Shell(); err != nil {
		rw.Write([]byte(err.Error() + "\r\n"))
	}

	if err := session.Wait(); err != nil {
		rw.Write([]byte(err.Error() + "\r\n"))
	}

}
