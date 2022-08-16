package slaver

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/helper"
	"tdp-cloud/core/socket"
)

type SlaveNode struct {
	*socket.JsonPod
	UserId     uint
	HostId     string
	SystemStat *helper.SystemStat
}

var NodePool = map[string]*SlaveNode{}

func Upgrader(c *gin.Context) {

	pod, err := socket.NewJsonPod(c.Writer, c.Request)

	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	defer pod.Close()

	// 注册节点

	node := &SlaveNode{
		pod, c.GetUint("UserId"), "", &helper.SystemStat{},
	}

	Receiver(node)

	// 清理资源

	if node.HostId != "" {
		delete(NodePool, node.HostId)
	}

}

func NodesOfUser(userId uint) *[]any {

	items := []any{}

	for _, v := range NodePool {
		if userId == v.UserId {
			items = append(items, map[string]any{
				"HostId":     v.HostId,
				"RemoteAddr": v.Conn.RemoteAddr().String(),
				"SystemStat": v.SystemStat,
			})
		}
	}

	return &items

}
