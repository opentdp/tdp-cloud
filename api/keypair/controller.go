package keypair

import (
	"github.com/gin-gonic/gin"
	"github.com/opentdp/go-helper/secure"

	"tdp-cloud/model/keypair"
)

type Controller struct{}

// 密钥列表

func (*Controller) list(c *gin.Context) {

	var rq *keypair.FetchAllParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	rq.UserId = c.GetUint("UserId")

	if lst, err := keypair.FetchAll(rq); err == nil {
		c.Set("Payload", gin.H{"Items": lst})
	} else {
		c.Set("Error", err)
	}

}

// 获取密钥

func (*Controller) detail(c *gin.Context) {

	var rq *keypair.FetchParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if res, err := keypair.Fetch(rq); err == nil {
		c.Set("Payload", gin.H{"Item": res})
	} else {
		c.Set("Error", err)
	}

}

// 添加密钥

func (*Controller) create(c *gin.Context) {

	var rq *keypair.CreateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	rq.UserId = c.GetUint("UserId")
	rq.StoreKey = c.GetString("AppKey")

	if id, err := keypair.Create(rq); err == nil {
		c.Set("Payload", gin.H{"Id": id})
		c.Set("Message", "添加成功")
	} else {
		c.Set("Error", err)
	}

}

// 修改密钥

func (*Controller) update(c *gin.Context) {

	var rq *keypair.UpdateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")
	rq.StoreKey = c.GetString("AppKey")

	if err := keypair.Update(rq); err == nil {
		c.Set("Message", "修改成功")
	} else {
		c.Set("Error", err)
	}

}

// 删除密钥

func (*Controller) delete(c *gin.Context) {

	var rq *keypair.DeleteParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if err := keypair.Delete(rq); err == nil {
		c.Set("Message", "删除成功")
	} else {
		c.Set("Error", err)
	}

}

// 生成密钥

func (*Controller) keygen(c *gin.Context) {

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
