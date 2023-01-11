package workhub

type ExecPayload struct {
	Name             string
	CommandType      string
	Content          string
	Username         string
	WorkingDirectory string
	Timeout          uint
}

func (pod *SendPod) Exec(data *ExecPayload) (uint, error) {

	taskId := createHistory(pod, data)

	v := &SocketData{
		Method:  "Exec",
		TaskId:  taskId,
		Payload: data,
	}

	return v.TaskId, pod.Write(v)

}

func (pod *RespPod) Exec(rq *SocketData) {

	updateHistory(pod, rq)

}
