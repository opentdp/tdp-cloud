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

	userId := c.GetUint("UserId")
	workerId := c.Query("WorkerId")

	worker := &Worker{
		pod,
		userId,
		workerId,
		c.Query("OSType"),
		c.Query("HostName"),
		&psutil.SystemStat{},
	}

	nodePool[workerId] = worker
	defer delete(nodePool, workerId)

	// 绑定主机

	err = bindMachine(c.Param("id"), workerId)

	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	// 启动服务

	Daemon(worker)

}

func NodesOfUser(userId uint) *[]any {

	items := []any{}

	for k, v := range nodePool {
		if userId == v.UserId {
			items = append(items, map[string]any{
				"WorkerId":   k,
				"RemoteAddr": v.Conn.RemoteAddr().String(),
				"SystemStat": v.SystemStat,
			})
		}
	}

	return &items

}

func NewSender(workerId string) *SendPod {

	if node, ok := nodePool[workerId]; ok {
		return &SendPod{node}
	}

	return nil

}
