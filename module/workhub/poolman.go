package workhub

import (
	"time"

	"github.com/gin-gonic/gin"

	"tdp-cloud/helper/psutil"
	"tdp-cloud/helper/socket"
)

var workerResp = map[uint]any{}
var workerPool = map[string]*Worker{}

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
		&psutil.SummaryStat{},
	}

	worker.WorkerMeta.From(c.GetHeader("TDP-Worker-Meta"))

	// 注册主机

	workerPool[worker.WorkerId] = worker
	defer delete(workerPool, worker.WorkerId)

	if err = updateMachine(worker); err != nil {
		return err
	}

	// 启动服务

	return Daemon(worker)

}

func WorkerOfUser(userId uint) []*Worker {

	items := []*Worker{}

	for _, v := range workerPool {
		if userId == v.UserId {
			items = append(items, v)
		}
	}

	return items

}

func NewSender(id string) *SendPod {

	if worker, ok := workerPool[id]; ok {
		return &SendPod{worker}
	}

	return nil

}

func WaitResponse(id uint, wait int) any {

	for i := 0; i < wait; i++ {
		if res, ok := workerResp[id]; ok {
			delete(workerResp, id)
			return res
		}
		time.Sleep(300 * time.Millisecond)
	}

	return ""

}
