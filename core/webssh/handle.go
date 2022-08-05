package webssh

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

func Handle(c *gin.Context) {

	wsp, err := NewSocketPod(c.Writer, c.Request)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	defer wsp.Close()

	// 获取 SSH 参数

	var option SSHClientOption

	if err := c.ShouldBindQuery(&option); err != nil {
		wsp.Write([]byte("> " + err.Error() + "\r\n"))
		return
	}

	// 创建 SSH 连接

	client, err := NewSSHClient(&option)
	if err != nil {
		wsp.Write([]byte("> " + err.Error() + "\r\n"))
		return
	}

	defer client.Close()

	// 转发 SSH 会话

	quit := make(chan bool, 1)
	go sshBridge(client, wsp, quit)
	<-quit

}

func sshBridge(client *ssh.Client, wsp *SocketPod, quit chan bool) {

	defer func() {
		quit <- true
	}()

	rw := io.ReadWriter(wsp)

	session, err := client.NewSession()
	if err != nil {
		rw.Write([]byte(err.Error() + "\r\n"))
		return
	}

	defer session.Close()

	// 客户端断开时清理资源
	wsp.OnClose(session.Close)

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
