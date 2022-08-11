package agent

import (
	"log"

	"tdp-cloud/core/socket"
)

type SocketData struct {
	TaskId  string
	Method  string
	Success bool
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

	for {
		var rq *SocketData

		if pod.Read(&rq) != nil {
			break
		}

		switch rq.Method {
		case "Ping":
			if RecvPing(addr, rq) != nil {
				return
			}
		default:
			log.Println("recv:", rq)
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
