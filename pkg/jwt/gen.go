package jwt

import (
	"github.com/spf13/viper"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenToken 生成JWT
func GenToken(userID int64, username string) (string, error) {
	tokenExpireDuration := time.Duration(viper.GetInt("auth.jwt_expire")) * time.Hour
	// 创建一个我们自己的声明的数据
	c := UserClaims{
		userID,
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenExpireDuration).Unix(),
			Issuer:    "ztalk",
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(mySecret)
}
