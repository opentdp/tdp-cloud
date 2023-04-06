package webssh

import (
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/net/websocket"
	"golang.org/x/term"
)

func Connect(ws *websocket.Conn, opt *SSHClientOption) error {

	defer ws.Close()

	// 创建客户端

	client, err := NewSSHClient(opt)

	if err != nil {
		ws.Write([]byte("> " + err.Error() + "\r\n"))
		return err
	}

	defer client.Close()

	// 打开新会话

	session, err := client.NewSession()

	if err != nil {
		ws.Write([]byte(err.Error() + "\r\n"))
		return err
	}

	defer session.Close()

	// 绑定输入输出

	session.Stdin = ws
	session.Stdout = ws
	session.Stderr = ws

	// 创建模拟终端

	fd := int(os.Stdin.Fd())
	width, height, _ := term.GetSize(fd)

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	if err := session.RequestPty("xterm", height, width, modes); err != nil {
		ws.Write([]byte(err.Error() + "\r\n"))
		return err
	}

	if err := session.Shell(); err != nil {
		ws.Write([]byte(err.Error() + "\r\n"))
		return err
	}

	if err := session.Wait(); err != nil {
		ws.Write([]byte(err.Error() + "\r\n"))
		return err
	}

	return nil

}
