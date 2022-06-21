package midware

import (
	"github.com/gin-gonic/gin"
)

func NewError(message string) gin.H {

	return gin.H{"Error": gin.H{"Message": message}}

}

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
