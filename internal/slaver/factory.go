package slaver

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/internal/helper"
	"tdp-cloud/internal/socket"
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

	hostId := c.Query("HostId")
	userId := c.GetUint("UserId")

	NodePool[hostId] = &SlaveNode{
		pod, userId, hostId, &helper.SystemStat{},
	}

	defer delete(NodePool, hostId)

	// 监听数据

	Receiver(NodePool[hostId])

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

func NewSender(hostId string) *SendPod {

	if node, ok := NodePool[hostId]; ok {
		return &SendPod{node}
	}

	return nil

}
