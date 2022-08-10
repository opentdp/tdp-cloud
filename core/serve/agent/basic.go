package agent

import (
	"errors"
	"tdp-cloud/core/socket"
)

type SocketData struct {
	Action  string
	Method  string
	Payload any
}

var AgentPool = map[string]*socket.JsonPod{}

func Register(wsp *socket.JsonPod) {

	ip := wsp.Conn.RemoteAddr().String()

	AgentPool[ip] = wsp

	defer delete(AgentPool, ip)

	// 保持客户端连接

	for {
		var rq SocketData

		if wsp.Read(&rq) != nil {
			break
		}

		if rq.Action == "ping" && rq.Method == "request" {
			rs := Ping(rq)
			if wsp.Write(&rs) != nil {
				break
			}
		}
	}

}

func SendAction(addr string, data SocketData) error {

	wsp := AgentPool[addr]

	if wsp != nil {
		wsp.Write(data)
	}

	return errors.New("客户端已断开")

}

func GetAgents() []string {

	items := make([]string, 0, len(AgentPool))

	for k := range AgentPool {
		items = append(items, k)
	}

	return items

}
