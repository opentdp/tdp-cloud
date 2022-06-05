package user

import (
	"tdp-cloud/core/dborm"
	"tdp-cloud/core/utils"

	"golang.org/x/crypto/bcrypt"
)

// 登录账号

func Login(username string, password string) (string, string) {

	var user dborm.User

	// 验证账号

	dborm.Db.First(&user, "username = ?", username)

	if user.ID == 0 {
		return "", "账号错误"
	}

	// 验证密码

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return "", "密码错误"
	}

	// 创建令牌

	token := utils.RandString(32)
	dborm.Db.Create(&dborm.Session{UserID: user.ID, Token: token})

	return token, ""
}

// 注册账号

func Register(username string, password string) (string, string) {

	var user dborm.User

	// 验证账号

	dborm.Db.First(&user, "username = ?", username)

	if user.ID > 0 {
		return "", "账号已被使用"
	}

	// 创建账号

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	dborm.Db.Create(&dborm.User{Username: username, Password: string(hash)})

	return "账号注册成功", ""

}

// 获取会话

func FetchSession(token string) dborm.Session {

	var session dborm.Session

	dborm.Db.First(&session, "token = ?", token)

	return session

}

// 获取密钥

func FetchSecret(keyId string, userId uint) dborm.Secret {

	var secret dborm.Secret

	dborm.Db.First(&secret, "id = ? AND user_id = ?", keyId, userId)

	return secret

}

// 获取密钥列表

func FetchSecrets(userId uint) ([]*dborm.Secret, string) {

	var secret []*dborm.Secret

	dborm.Db.Find(&secret, "user_id = ?", userId)

	return secret, ""

}

// 添加密钥

func CreateSecret(post *SecretInput) (string, string) {

	dborm.Db.Create(&dborm.Secret{
		UserID:    post.UserID,
		Describe:  post.Describe,
		SecretId:  post.SecretId,
		SecretKey: post.SecretKey,
	})

	return "添加成功", ""

}

// 删除密钥

func DeleteSecret(id uint) (string, string) {

	dborm.Db.Delete(&dborm.Secret{}, id)

	return "删除成功", ""

}
