package middleWare

import (
	"boke-server/common/res"
	"boke-server/service/redis_service"
	"boke-server/utils/jwt"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

var filterRoute = []string{
	"/api/article/list",
	"/api/category/list",
	"/api/article/details",
	"/api/comment/list",
	"/api/user/details",
}

func CheckToken(c *gin.Context) {
	//如果在白名单中则直接跳过
	for _, val := range filterRoute {
		if val == c.Request.URL.Path {
			c.Next()
			return
		}
	}
	fmt.Println(c.Request.URL.Path)
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		res.ExpireMsg("token过期", c)
		c.Abort()
		return
	}
	parts := strings.Split(authHeader, " ")
	if len(parts) == 2 && parts[0] == "Bearer" {
		token := parts[1]
		if redis_service.FindTokenIsBlack(token) {
			res.ExpireMsg("token过期", c)
			c.Abort()
			return
		}
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
			res.ExpireMsg("解析token失败", c)
			c.Abort() //中止
		}
	} else {
		res.FailedMsg("请携带token", c)
		c.Abort()
	}
}
