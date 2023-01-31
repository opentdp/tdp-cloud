package midware

import (
	"errors"
)

type H map[string]any

// 创建错误实例

func NewError(data any) error {

	if err, ok := data.(error); ok {
		return err
	}

	if err, ok := data.(string); ok {
		return errors.New(err)
	}

	return errors.New("未知错误")

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

// 构造结构数据

func NewPayload(data any) H {

	if msg, ok := data.(string); ok {
		return H{"Payload": H{"Message": msg}}
	}

	return H{"Payload": data}

}
