package agent

import (
	"tdp-cloud/core/socket"
)

type SocketData struct {
	Action  string
	Method  string
	Payload any
}

type AgentNode struct {
	Pod  *socket.JsonPod
	Addr string
}

var AgentPool = map[string]AgentNode{}

func Register(pod *socket.JsonPod) {

	addr := pod.Conn.RemoteAddr().String()

	AgentPool[addr] = AgentNode{
		Addr: addr,
		Pod:  pod,
	}

	defer delete(AgentPool, addr)

	// 保持客户端连接

	for {
		var rq SocketData

		if pod.Read(&rq) != nil {
			break
		}

		switch rq.Action {
		case "ping":
			if Pong(addr, &rq.Payload) != nil {
				return
			}
		}
	}

}

func GetNodeList() []AgentNode {

	items := make([]AgentNode, 0, len(AgentPool))

	for _, v := range AgentPool {
		items = append(items, AgentNode{
			Addr: v.Addr,
		})
	}

	return items

}
