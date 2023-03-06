package midware

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserInfo struct {
	UserId    uint
	UserLevel uint
	SecretKey string
}

type UserClaims struct {
	jwt.RegisteredClaims
	*UserInfo
}

var signKey = []byte("sdfsdfsdfsdfsdf")

func CreateToken(userInfo *UserInfo) (string, error) {

	claims := UserClaims{
		jwt.RegisteredClaims{
			Issuer: "TDP Cloud",
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(7 * time.Hour),
			},
		},
		userInfo,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(signKey)

}

func ParserToken(signToken string) (*UserClaims, error) {

	var claims UserClaims

	keyFunc := func(token *jwt.Token) (any, error) {
		return signKey, nil
	}

	token, err := jwt.ParseWithClaims(signToken, &claims, keyFunc)

	if token.Valid {
		return &claims, nil
	}

	return nil, err

}
