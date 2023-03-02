package user

import (
	"errors"
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// 验证用户

func CheckUser(u, p, e string) error {

	ul, pl, el := len(u), len(p), len(e)

	// 校验用户名

	if ul > 0 {
		if ul < 4 || ul > 32 {
			return errors.New("用户名长度不符合要求")
		}
		usernameExpr := "^[0-9a-zA-Z\u3040-\u309F\u30A0-\u30FF\u4E00-\u9FA5\uF900-\uFA2D]+$"
		if !regexp.MustCompile(usernameExpr).MatchString(u) {
			return errors.New("用户名禁止使用特殊字符")
		}
	}

	// 校验密码

	if pl > 0 {
		if pl < 6 || pl > 32 {
			return errors.New("密码长度不符合要求")
		}
		if strings.Contains(u, p) || strings.Contains(p, u) {
			return errors.New("密码与用户名不可包含")
		}
	}

	// 校验邮箱

	if el > 0 {
		if el < 6 || el > 128 {
			return errors.New("邮箱长度不符合要求")
		}
		emailExpr := `^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`
		if !regexp.MustCompile(emailExpr).MatchString(e) {
			return errors.New("邮箱格式不正确")
		}
	}

	return nil

}

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
