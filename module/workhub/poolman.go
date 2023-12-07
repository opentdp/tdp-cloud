package workhub

import (
	"time"

	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/socket"
)

var workerPool = map[string]*Worker{}
var workerResp = map[string]*socket.PlainData{}

func DeleteWorker(Worker *Worker) {

	logman.Info("Worker disconnect", "Id", Worker.WorkerId)

	if Worker.WorkerId != "" {
		delete(workerPool, Worker.WorkerId)
	}

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

func GetSendPod(id string) *SendPod {

	if worker, ok := workerPool[id]; ok {
		return &SendPod{worker}
	}

	return nil

}

func WaitResponse(id string, wait int) *socket.PlainData {

	for i := 0; i < wait; i++ {
		if res, ok := workerResp[id]; ok {
			delete(workerResp, id)
			return res
		}
		time.Sleep(300 * time.Millisecond)
	}

	return &socket.PlainData{
		Success: false,
		Message: "请求超时",
	}

}
