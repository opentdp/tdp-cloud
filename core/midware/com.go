package midware

import (
	"github.com/gin-gonic/gin"
)

func NewError(message string) gin.H {

	return gin.H{"Error": gin.H{"Message": message}}

}

// 获取会话数据

type Userdata struct {
	KeyId     int
	UserId    int
	Region    string
	SecretId  string
	SecretKey string
}

func GetUserdata(c *gin.Context) *Userdata {

	ud := &Userdata{
		KeyId:     c.GetInt("KeyId"),
		UserId:    c.GetInt("UserId"),
		Region:    c.GetString("Region"),
		SecretId:  c.GetString("SecretId"),
		SecretKey: c.GetString("SecretKey"),
	}

	return ud

}
