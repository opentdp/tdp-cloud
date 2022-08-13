package agent

import (
	"log"

	"tdp-cloud/core/helper"
	"tdp-cloud/core/socket"
)

type RecvPod struct {
	*socket.JsonPod
}

type SendPod struct {
	*socket.JsonPod
}

type AgentNode struct {
	Addr string
	Pod  *socket.JsonPod
	Stat *helper.SystemStat
}

type SocketData struct {
	TaskId  string
	Method  string
	Success bool
	Payload any
}

var AgentPool = map[string]AgentNode{}

func AddNode(pod *socket.JsonPod) {

	addr := pod.Conn.RemoteAddr().String()

	AgentPool[addr] = AgentNode{
		Addr: addr,
		Pod:  pod,
		Stat: &helper.SystemStat{},
	}

	defer delete(AgentPool, addr)

	// 接收数据

	recv := NewRecvPod(pod)

	for {
		var rq *SocketData

		if pod.Read(&rq) != nil {
			break
		}

		switch rq.Method {
		case "Ping":
			if recv.Ping(rq) != nil {
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
		items = append(items, v)
	}

	return items

}

func NewRecvPod(pod *socket.JsonPod) *RecvPod {

	return &RecvPod{pod}

}

func NewSendPod(addr string) *SendPod {

	if node, ok := AgentPool[addr]; ok {
		return &SendPod{node.Pod}
	}

	return nil

}
