package workhub

import (
	"tdp-cloud/helper/json"
	"tdp-cloud/internal/dborm/machine"
	history "tdp-cloud/internal/dborm/task_history"
)

// 主机

func createMachine(node *Worker) {

	item := &machine.CreateParam{
		UserId:      node.UserId,
		VendorId:    0,
		HostName:    node.HostName,
		IpAddress:   node.Conn.RemoteAddr().String(),
		OSType:      node.OSType,
		Region:      "",
		Model:       "worker",
		CloudId:     node.HostId,
		CloudMeta:   json.ToString(node.SystemStat),
		Description: "",
		Status:      "{}",
	}

	machine.Create(item)

}

func deleteMachine(node *Worker) {
}

// 任务历史

func createHistory(pod *SendPod, data *ExecPayload) uint {

	item := &history.CreateParam{
		UserId:   pod.UserId,
		HostId:   pod.SystemStat.HostId,
		Subject:  "Exec: " + data.Name,
		HostName: pod.SystemStat.HostName,
		Request:  json.ToString(data),
		Response: "",
		Status:   "Doing",
	}

	id, _ := history.Create(item)

	return id

}

func updateHistory(pod *RespPod, rq *SocketData) error {

	status := "Failed"
	if rq.Success {
		status = "Success"
	}

	item := &history.UpdateParam{
		Id:       rq.TaskId,
		UserId:   pod.UserId,
		Response: json.ToString(rq.Payload),
		Status:   status,
	}

	return history.Update(item)

}
