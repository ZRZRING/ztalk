package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

// ParseToken 解析JWT
func ParseToken(tokenString string) (*UserClaims, error) {
	var mc = new(UserClaims)
	token, err := jwt.ParseWithClaims(
		tokenString, mc, func(token *jwt.Token) (interface{}, error) {
			return mySecret, nil
		})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
