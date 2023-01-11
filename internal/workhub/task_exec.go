package workhub

type ExecPayload struct {
	Name          string
	CommandType   string
	Content       string
	Username      string
	WorkDirectory string
	Timeout       uint
}

func (pod *SendPod) Exec(data *ExecPayload) (uint, error) {

	taskId := createHistory(pod, data)

	rq := &SocketData{
		Method:  "Exec",
		TaskId:  taskId,
		Payload: data,
	}

	return rq.TaskId, pod.Write(rq)

}

func (pod *RespPod) Exec(rq *SocketData) {

	updateHistory(pod, rq)

}
