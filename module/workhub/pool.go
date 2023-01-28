package workhub

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/helper/psutil"
	"tdp-cloud/helper/socket"
)

var nodePool = map[string]*Worker{}

func Register(c *gin.Context) error {

	pod, err := socket.NewJsonPod(c.Writer, c.Request)

	if err != nil {
		return err
	}

	defer pod.Close()

	// 注册节点

	systemStat := &psutil.SystemStat{}
	systemStat.From(c.GetHeader("TDP-HostStat"))

	worker := &Worker{
		pod,
		c.GetUint("UserId"),
		c.GetUint("MachineId"),
		c.GetHeader("TDP-WorkerId"),
		systemStat,
	}

	nodePool[worker.WorkerId] = worker
	defer delete(nodePool, worker.WorkerId)

	// 绑定主机

	if err = bindMachine(worker); err != nil {
		return err
	}

	// 启动服务

	return Daemon(worker)

}

func NodesOfUser(userId uint) []*Worker {

	items := []*Worker{}

	for _, v := range nodePool {
		if userId == v.UserId {
			items = append(items, v)
		}
	}

	return items

}

func NewSender(workerId string) *SendPod {

	if node, ok := nodePool[workerId]; ok {
		return &SendPod{node}
	}

	return nil

}
