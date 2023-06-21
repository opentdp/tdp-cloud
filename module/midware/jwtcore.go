package midware

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"tdp-cloud/cmd/args"
)

type UserInfo struct {
	Id     uint
	Level  uint
	AppKey string
}

type UserClaims struct {
	jwt.RegisteredClaims
	*UserInfo
}

func CreateToken(userInfo *UserInfo) (string, error) {

	claims := UserClaims{
		jwt.RegisteredClaims{
			Issuer: args.AppName,
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(3 * time.Hour),
			},
		},
		userInfo,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtkey := args.Server.JwtKey
	return token.SignedString([]byte(jwtkey))

}

func UpdateToken(signToken string) (string, error) {

	claims, err := ParserToken(signToken)
	if err != nil {
		return "", err
	}

	claims.ExpiresAt.Time = time.Now().Add(3 * time.Hour)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtkey := args.Server.JwtKey
	return token.SignedString([]byte(jwtkey))

}

func ParserToken(signToken string) (*UserClaims, error) {

	claims := &UserClaims{}

	keyFunc := func(token *jwt.Token) (any, error) {
		jwtkey := args.Server.JwtKey
		return []byte(jwtkey), nil
	}

	token, err := jwt.ParseWithClaims(signToken, claims, keyFunc)

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil

}
