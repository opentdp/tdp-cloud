package webssh

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

func Handle(c *gin.Context, option *SSHClientOption) {

	log.Println("Webssh - Connecting")

	wsConn, err := wsUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.Set("Error", err)
		c.Abort()
		return
	}

	defer wsConn.Close()

	wsw := &wsWrapper{wsConn}

	sshClient, err := NewSSHClient(option)
	if err != nil {
		wsw.Write([]byte("> " + err.Error() + "\r\n"))
		return
	}

	defer sshClient.Close()

	quit := make(chan bool, 1)
	go sshHandle(sshClient, wsw, quit)
	<-quit

	log.Println("Webssh - Disconnected")

}

func sshHandle(sshClient *ssh.Client, wsw *wsWrapper, quit chan bool) {

	defer func() {
		quit <- true
	}()

	rw := io.ReadWriter(wsw)

	session, err := sshClient.NewSession()
	if err != nil {
		rw.Write([]byte(err.Error() + "\r\n"))
		return
	}

	defer session.Close()

	// 客户端断开连接时清理资源
	wsw.SetCloseHandler(func() error {
		return session.Close()
	})

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
