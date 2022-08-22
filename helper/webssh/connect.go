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

	session, err := client.NewSession()

	if err != nil {
		pod.Write([]byte(err.Error() + "\r\n"))
		return err
	}

	defer session.Close()

	pod.OnClose(session.Close)

	// 代理 SSH 会话

	return sshProxy(session, pod)

}

func sshProxy(session *ssh.Session, rw io.ReadWriter) error {

	// 绑定输入输出

	session.Stdin = rw
	session.Stdout = rw
	session.Stderr = rw

	// 创建模拟终端

	fd := int(os.Stdin.Fd())
	width, height, _ := term.GetSize(fd)

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	if err := session.RequestPty("xterm", height, width, modes); err != nil {
		rw.Write([]byte(err.Error() + "\r\n"))
		return err
	}

	if err := session.Shell(); err != nil {
		rw.Write([]byte(err.Error() + "\r\n"))
		return err
	}

	if err := session.Wait(); err != nil {
		rw.Write([]byte(err.Error() + "\r\n"))
		return err
	}

	return nil

}
