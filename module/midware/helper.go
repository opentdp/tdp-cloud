package midware

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// 获取错误代码

func errCode(c *gin.Context) int {

	if code := c.GetInt("ErrorCode"); code > 400 {
		return code
	}

	return 400

}

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

func NewErrorMessage(data any) gin.H {

	if err, ok := data.(error); ok {
		return gin.H{"Error": gin.H{"Message": err.Error()}}
	}

	if err, ok := data.(string); ok {
		return gin.H{"Error": gin.H{"Message": err}}
	}

	return gin.H{"Error": data}

}

// 构造结构数据

func NewPayloadMessage(data any, msg string) gin.H {

	if msg != "" {
		return gin.H{"Payload": data, "Message": msg}
	}

	if msg, ok := data.(string); ok {
		return gin.H{"Message": msg}
	}

	return gin.H{"Payload": data}

}
