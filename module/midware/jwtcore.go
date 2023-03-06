package midware

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
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

func CreateToken(userInfo *UserInfo) (string, error) {

	jwtkey := viper.GetString("server.jwtkey")

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

	return token.SignedString([]byte(jwtkey))

}

func ParserToken(signToken string) (*UserClaims, error) {

	var claims UserClaims

	keyFunc := func(token *jwt.Token) (any, error) {
		jwtkey := viper.GetString("server.jwtkey")
		return []byte(jwtkey), nil
	}

	token, err := jwt.ParseWithClaims(signToken, &claims, keyFunc)

	if token.Valid {
		return &claims, nil
	}

	return nil, err

}
