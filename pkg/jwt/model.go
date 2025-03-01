package jwt

import "github.com/dgrijalva/jwt-go"

var mySecret = []byte("zrzring")

const UserIDKey = "userID"

type UserClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}
