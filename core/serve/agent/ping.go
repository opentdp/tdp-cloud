package agent

func Ping(data SocketData) SocketData {
	return SocketData{
		Action:  "ping",
		Method:  "response",
		Payload: data.Payload,
	}
}
