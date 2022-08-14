package agent

import (
	"log"

	"tdp-cloud/core/helper"
	"tdp-cloud/core/socket"
)

type RecvPod struct {
	*socket.JsonPod
}

type RespPod RecvPod
type SendPod RecvPod

type AgentNode struct {
	Pod    *socket.JsonPod
	Stat   *helper.SystemStat
	Addr   string
	UserId uint
}

type SocketData struct {
	TaskId  string
	Method  string
	Success bool
	Payload any
}

var AgentPool = map[string]AgentNode{}

func AddNode(pod *socket.JsonPod, userId uint) {

	addr := pod.Conn.RemoteAddr().String()

	AgentPool[addr] = AgentNode{
		Pod:    pod,
		Stat:   &helper.SystemStat{},
		Addr:   addr,
		UserId: userId,
	}

	defer delete(AgentPool, addr)

	// 接收数据

	recv := &RecvPod{pod}
	resp := &RespPod{pod}

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

func GetNodeList(userId uint) []AgentNode {

	items := make([]AgentNode, 0, len(AgentPool))

	for _, v := range AgentPool {
		if userId == v.UserId {
			items = append(items, v)
		}
	}

	return items

}

func NewSendPod(addr string) *SendPod {

	if node, ok := AgentPool[addr]; ok {
		return &SendPod{node.Pod}
	}

	return nil

}
