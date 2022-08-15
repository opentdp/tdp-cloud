package agent

import (
	"log"

	"tdp-cloud/core/helper"
	"tdp-cloud/core/socket"
)

type RecvPod struct {
	*AgentNode
}
type RespPod struct {
	*AgentNode
}
type SendPod struct {
	*AgentNode
}

type AgentNode struct {
	*socket.JsonPod
	UserId     uint
	RemoteAddr string
	SystemStat *helper.SystemStat
}

type SocketData struct {
	TaskId  string
	Method  string
	Success bool
	Payload any
}

var NodePool = map[string]AgentNode{}

func AddNode(pod *socket.JsonPod, userId uint) {

	addr := pod.Conn.RemoteAddr().String()
	node := AgentNode{
		pod, userId, addr, &helper.SystemStat{},
	}

	NodePool[addr] = node
	defer delete(NodePool, addr)

	// 接收数据

	recv := &RecvPod{&node}
	resp := &RespPod{&node}

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
		case "Exec:resp":
			resp.Exec(rq)
		default:
			log.Println("recv:", rq)
		}
	}

}

func NodesOfUser(userId uint) []AgentNode {

	items := make([]AgentNode, 0, len(NodePool))

	for _, v := range NodePool {
		if userId == v.UserId {
			items = append(items, v)
		}
	}

	return items

}

func NewSendPod(addr string) *SendPod {

	if node, ok := NodePool[addr]; ok {
		return &SendPod{&node}
	}

	return nil

}
