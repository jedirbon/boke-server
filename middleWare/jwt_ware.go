package middleWare

import (
	"boke-server/common/res"
	"boke-server/utils/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func CheckToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		res.ExpireMsg("token过期", c)
		c.Abort()
		return
	}
	parts := strings.Split(authHeader, " ")
	if len(parts) == 2 && parts[0] == "Bearer" {
		token := parts[1]
		claims, ok := jwt.ParseToken(token)
		if ok {
			fmt.Println("**********")
			fmt.Println(claims)
			fmt.Println("**********")
			// 将用户信息存储到gin上下文中，供后续API使用
			c.Set("userId", claims.UserInfo.UserId)
			c.Set("roleId", claims.UserInfo.RoleId)
			c.Next() //放行
			return
		} else {
			res.FailedMsg("解析token失败", c)
			c.Abort() //中止
		}
	} else {
		res.FailedMsg("请携带token", c)
		c.Abort()
	}
}
