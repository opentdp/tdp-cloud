package helper

type H map[string]any

// 构造结构数据

func NewPayload(data any) H {

	if msg, ok := data.(string); ok {
		return H{"Payload": H{"Message": msg}}
	}

	return H{"Payload": data}

}

// 构造错误信息

func NewMessage(data any) H {

	if err, ok := data.(error); ok {
		return H{"Error": H{"Message": err.Error()}}
	}

	if err, ok := data.(string); ok {
		return H{"Error": H{"Message": err}}
	}

	return H{"Error": data}

}
