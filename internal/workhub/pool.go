package workhub

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/helper/psutil"
	"tdp-cloud/helper/socket"
)

var nodePool = map[string]*Worker{}

func Register(c *gin.Context) {

	pod, err := socket.NewJsonPod(c.Writer, c.Request)

	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	defer pod.Close()

	// 注册节点

	worker := &Worker{
		pod,
		c.GetUint("UserId"),
		c.Query("OSType"),
		c.Query("HostId"),
		c.Query("HostName"),
		&psutil.SystemStat{},
	}

	nodePool[worker.HostId] = worker
	defer delete(nodePool, worker.HostId)

	// 启动服务

	Daemon(worker)

}

func NodesOfUser(userId uint) *[]any {

	items := []any{}

	for _, v := range nodePool {
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

	if node, ok := nodePool[hostId]; ok {
		return &SendPod{node}
	}

	return nil

}
