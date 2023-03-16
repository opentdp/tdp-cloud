package keypair

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/helper/secure"
)

func keygen(c *gin.Context) {

	var rq struct {
		KeyType string
	}

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	var (
		err                   error
		privateKey, publicKey string
	)

	switch rq.KeyType {
	case "ssh":
		privateKey, publicKey, err = secure.NewSSHKeypair()
	}

	if err == nil {
		c.Set("Payload", gin.H{"PrivateKey": privateKey, "PublicKey": publicKey})
	} else {
		c.Set("Error", err)
	}

}
