package webssh

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
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

	quitChan := make(chan bool, 1)
	go wsHandle(wsConn, sshClient, quitChan)
	<-quitChan

}

func wsHandle(wsConn *websocket.Conn, sshClient *ssh.Client, quitChan chan bool) {

	defer setQuit(quitChan)

	wsConn.SetCloseHandler(func(code int, text string) error {
		wsConn.Close()
		return nil
	})

	rw := io.ReadWriter(&wsWrapper{wsConn})

	sshHandle(rw, sshClient)

}

func sshHandle(rw io.ReadWriter, sshClient *ssh.Client) {

	session, err := sshClient.NewSession()
	if err != nil {
		rw.Write([]byte(err.Error() + "\r\n"))
		return
	}

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

func setQuit(ch chan bool) {
	ch <- true
}
