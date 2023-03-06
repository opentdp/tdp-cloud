package user

import (
	"errors"
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"tdp-cloud/helper/strutil"
	"tdp-cloud/module/dborm"
)

// 新建密码密钥

func CreateSecret(p, k string) (string, string, error) {

	if k == "" {
		k = strutil.Rand(32)
	}

	ak, err := strutil.Des3Encrypt(k, p)

	if err != nil {
		return "", "", err // 创建密钥失败
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	return string(hash), ak, err

}

// 修改密码密钥

func UpdateSecret(p, k string) (string, string, error) {

	if p == "" {
		return "", "", nil // 未设置密码时忽略
	}

	if k == "" {
		return "", "", errors.New("密钥不能为空")
	}

	return CreateSecret(p, k)

}

// 验证用户密钥

func CheckSecret(u *dborm.User, p, k string) bool {

	rs := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))
	if rs != nil {
		return false
	}

	if k != "" {
		ak, err := strutil.Des3Decrypt(u.AppKey, p)
		if err != nil || ak != k {
			return false
		}
	}

	return true

}

// 验证用户信息

func CheckUserinfo(u, p, e string) error {

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
