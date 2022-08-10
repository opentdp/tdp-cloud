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

type AgentNode struct {
	Pod *socket.JsonPod
	Ip  string
}

var AgentPool = map[string]AgentNode{}

func Register(pod *socket.JsonPod) {

	ip := pod.Conn.RemoteAddr().String()

	AgentPool[ip] = AgentNode{
		Pod: pod,
		Ip:  ip,
	}

	defer delete(AgentPool, ip)

	// 保持客户端连接

	for {
		var rq SocketData

		if pod.Read(&rq) != nil {
			break
		}

		if rq.Action == "ping" && rq.Method == "request" {
			rs := Ping(rq)
			if pod.Write(&rs) != nil {
				break
			}
		}
	}

}

func SendAction(addr string, data SocketData) error {

	node, ok := AgentPool[addr]

	if !ok {
		return errors.New("客户端已断开")
	}

	return node.Pod.Write(data)

}

func GetNodeList() []AgentNode {

	items := make([]AgentNode, 0, len(AgentPool))

	for _, v := range AgentPool {
		items = append(items, AgentNode{
			Ip: v.Ip,
		})
	}

	return items

}
