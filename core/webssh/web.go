package webssh

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebSSHConfig struct {
	RemoteAddr string
	User       string
	Password   string
	AuthModel  AuthModel
	PkPath     string
}

type WebSSH struct {
	*WebSSHConfig
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024 * 10,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func NewWebSSH(conf *WebSSHConfig) *WebSSH {
	return &WebSSH{
		WebSSHConfig: conf,
	}
}

func (w WebSSH) ServeConn(c *gin.Context) {

	wsConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	fmt.Println(err)
	if err != nil {
		c.AbortWithStatusJSON(200, gin.H{"ok": false, "msg": err.Error()})
		return
	}

	defer wsConn.Close()

	var config *SSHClientConfig

	switch w.AuthModel {
	case PASSWORD:
		config = SSHClientConfigPassword(
			w.RemoteAddr,
			w.User,
			w.Password,
		)
	case PUBLICKEY:
		config = SSHClientConfigPulicKey(
			w.RemoteAddr,
			w.User,
			w.PkPath,
		)
	}

	client, err := NewSSHClient(config)
	if err != nil {
		wsConn.WriteControl(websocket.CloseMessage, []byte(err.Error()), time.Now().Add(time.Second))
		return
	}

	defer client.Close()

	turn, err := NewTurn(wsConn, client)
	if err != nil {
		wsConn.WriteControl(websocket.CloseMessage, []byte(err.Error()), time.Now().Add(time.Second))
		return
	}

	defer turn.Close()

	logBuff := bufPool.Get().(*bytes.Buffer)
	logBuff.Reset()

	defer bufPool.Put(logBuff)

	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		err := turn.LoopRead(logBuff, ctx)
		if err != nil {
			log.Printf("%#v", err)
		}
	}()

	go func() {
		defer wg.Done()
		err := turn.SessionWait()
		if err != nil {
			log.Printf("%#v", err)
		}
		cancel()
	}()

	wg.Wait()

}
