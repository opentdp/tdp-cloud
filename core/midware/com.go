package midware

import (
	"github.com/gin-gonic/gin"
)

func NewError(message string) gin.H {

	return gin.H{"Error": gin.H{"message": message}}

}

func GetUserdata(c *gin.Context) Userdata {

	ud := Userdata{
		c.GetInt("KeyId"),
		c.GetInt("UserId"),
		c.GetString("Region"),
		c.GetString("SecretId"),
		c.GetString("SecretKey"),
	}

	return ud

}

type Userdata struct {
	KeyId     int
	UserId    int
	Region    string
	SecretId  string
	SecretKey string
}
