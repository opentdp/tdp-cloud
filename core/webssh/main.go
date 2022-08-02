package webssh

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

func Handle(c *gin.Context, option *SSHClientOption) {

	wsConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.Set("Error", err)
		c.Abort()
		return
	}

	defer wsConn.Close()

	sshClient, err := NewSSHClient(option)
	if err != nil {
		msg := "> " + err.Error() + "\r\n"
		wsConn.WriteMessage(websocket.TextMessage, []byte(msg))
		return
	}

	defer sshClient.Close()

	ch := make(chan bool, 1)
	go sshHandle(wsConn, sshClient, ch)
	<-ch

}

func sshHandle(wsConn *websocket.Conn, sshClient *ssh.Client, ch chan bool) {

	defer func() {
		ch <- true
	}()

	rw := io.ReadWriter(&readWriter{wsConn})

	session, err := sshClient.NewSession()
	if err != nil {
		rw.Write([]byte(err.Error() + "\r\n"))
		return
	}

	// 客户端关闭连接时清理会话
	wsConn.SetCloseHandler(func(code int, text string) error {
		session.Close()
		wsConn.Close()
		return nil
	})

	defer session.Close()

	session.Stdout = rw
	session.Stderr = rw
	session.Stdin = rw

	fd := int(os.Stdin.Fd())
	width, height, _ := term.GetSize(fd)

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	err = session.RequestPty("xterm", width, height, modes)
	if err != nil {
		rw.Write([]byte(err.Error() + "\r\n"))
	}

	err = session.Shell()
	if err != nil {
		rw.Write([]byte(err.Error() + "\r\n"))
	}

	err = session.Wait()
	if err != nil {
		rw.Write([]byte(err.Error() + "\r\n"))
	}

}
