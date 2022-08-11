package agent

func (pod *RecvPod) Ping(rq *SocketData) error {

	v := &SocketData{
		TaskId:  rq.TaskId,
		Method:  "Ping:end",
		Payload: rq.Payload,
	}

	return pod.Write(v)

}
