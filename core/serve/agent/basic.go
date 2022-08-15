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
	HostId     string
	SystemStat *helper.SystemStat
}

type SocketData struct {
	Method  string
	TaskId  uint
	Success bool
	Payload any
}

var NodePool = map[string]*AgentNode{}

func AddNode(pod *socket.JsonPod, userId uint) {

	node := &AgentNode{
		pod, userId, "", &helper.SystemStat{},
	}

	// 接收数据

	recv := &RecvPod{node}
	resp := &RespPod{node}

	for {
		var rq *SocketData

		if pod.Read(&rq) != nil {
			break
		}

		switch rq.Method {
		case "Exec:resp":
			resp.Exec(rq)
		case "Register":
			recv.Register(rq, node)
		case "Ping":
			recv.Ping(rq)
		default:
			log.Println("recv:", rq)
		}
	}

	// 清理资源

	if node.HostId != "" {
		delete(NodePool, node.HostId)
	}

}

func NodesOfUser(userId uint) []any {

	items := make([]any, 0, len(NodePool))

	for _, v := range NodePool {
		if userId == v.UserId {
			items = append(items, map[string]any{
				"HostId":     v.HostId,
				"RemoteAddr": v.Conn.RemoteAddr().String(),
				"SystemStat": v.SystemStat,
			})
		}
	}

	return items

}

func NewSendPod(hostId string) *SendPod {

	if node, ok := NodePool[hostId]; ok {
		return &SendPod{node}
	}

	return nil

}
