package midware

import (
	"github.com/gin-gonic/gin"
)

// 获取会话数据

type Userdata struct {
	KeyId     uint
	UserId    uint
	Region    string
	SecretId  string
	SecretKey string
}

func GetUserdata(c *gin.Context) *Userdata {

	ud := &Userdata{
		KeyId:     c.GetUint("KeyId"),
		UserId:    c.GetUint("UserId"),
		Region:    c.GetString("Region"),
		SecretId:  c.GetString("SecretId"),
		SecretKey: c.GetString("SecretKey"),
	}

	return ud

}

// 构造结构数据

func NewPayload(data any) gin.H {

	if msg, ok := data.(string); ok {
		return gin.H{"Payload": gin.H{"Message": msg}}
	}

	return gin.H{"Payload": data}

}

// 构造错误信息

func NewError(data any) gin.H {

	if err, ok := data.(error); ok {
		return gin.H{"Error": gin.H{"Message": err.Error()}}
	}

	if err, ok := data.(string); ok {
		return gin.H{"Error": gin.H{"Message": err}}
	}

	return gin.H{"Error": data}

}
