package jwt

import (
	"boke-server/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type UserInfo struct {
	RoleId int  `json:"roleId"`
	UserId uint `json:"userId"`
}
type MyClaims struct {
	UserInfo UserInfo `json:"userInfo"`
	jwt.RegisteredClaims
}

// 生成 JWT Token
func CreateToken(userInfo UserInfo) (string, error) {
	// 创建 JWT 负载
	claims := &MyClaims{
		UserInfo: userInfo,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.Config.Jwt.Time) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    global.Config.Jwt.Person,
		},
	}

	// 创建 Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 进行签名
	return token.SignedString([]byte(global.Config.Jwt.Key)) //使用密钥进行签名
}

// 验证token
func ParseToken(tokenString string) (*MyClaims, bool) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 确保使用的是正确的签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(global.Config.Jwt.Key), nil
	})

	if err != nil {
		return nil, false
	}

	// 验证 Token 是否有效
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, true
	} else {
		return nil, false
	}
}

func ParseTokenByGin(c *gin.Context) {
	header := c.GetHeader("Authorization")
	fmt.Println(header)
}
