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

	// 节点信息

	worker := &Worker{
		pod,
		c.GetUint("UserId"),
		c.GetUint("MachineId"),
		c.GetHeader("TDP-Worker-Id"),
		&psutil.SystemInfo{},
	}

	worker.WorkerMeta.From(c.GetHeader("TDP-Worker-Meta"))

	// 注册主机

	nodePool[worker.WorkerId] = worker
	defer delete(nodePool, worker.WorkerId)

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
