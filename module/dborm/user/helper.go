package user

import (
	"golang.org/x/crypto/bcrypt"
)

// 生成密码

func HashPassword(p1 string) string {

	hash, _ := bcrypt.GenerateFromPassword([]byte(p1), bcrypt.DefaultCost)

	return string(hash)

}

// 验证密码

func CheckPassword(p1, p2 string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(p1), []byte(p2))

	return err == nil

}
