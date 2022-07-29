package webssh

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/ssh"
	"golang.org/x/term"

	"tdp-cloud/core/utils"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024 * 1024 * 10,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Handle(c *gin.Context, option *SSHClientOption) {

	wsConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.AbortWithStatusJSON(500, utils.NewMessage(err))
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

	quitChan := make(chan bool, 1)
	go wsHandle(wsConn, sshClient, quitChan)
	<-quitChan

}

func wsHandle(wsConn *websocket.Conn, sshClient *ssh.Client, quitChan chan bool) {

	defer setQuit(quitChan)

	rw := io.ReadWriter(&wsWrapper{wsConn})

	webPrintln := func(data string) {
		rw.Write([]byte(data + "\r\n"))
	}

	wsConn.SetCloseHandler(func(code int, text string) error {
		wsConn.Close()
		return nil
	})

	sshHandle(rw, sshClient, webPrintln)

}

func sshHandle(rw io.ReadWriter, sshClient *ssh.Client, errhandle func(string)) {

	session, err := sshClient.NewSession()
	if err != nil {
		errhandle(err.Error())
		return
	}

	defer session.Close()

	session.Stdout = rw
	session.Stderr = rw
	session.Stdin = rw

	fd := int(os.Stdin.Fd())
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
	termWidth, termHeight, _ := term.GetSize(fd)

	err = session.RequestPty("xterm", termHeight, termWidth, modes)
	if err != nil {
		errhandle(err.Error())
	}

	err = session.Shell()
	if err != nil {
		errhandle(err.Error())
	}

	err = session.Wait()
	if err != nil {
		errhandle(err.Error())
	}

}

func setQuit(ch chan bool) {
	ch <- true
}
