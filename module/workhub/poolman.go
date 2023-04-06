package workhub

import (
	"time"
)

var workerResp = map[uint]any{}
var workerPool = map[string]*Worker{}

func DeleteWorker(Worker *Worker) {

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
