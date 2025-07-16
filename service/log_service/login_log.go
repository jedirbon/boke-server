package log_service

import (
	"boke-server/core"
	"boke-server/global"
	"boke-server/models"
	"boke-server/models/enum"
	"fmt"
	"github.com/gin-gonic/gin"
)

func NewLoginSuccess(c *gin.Context, loginType enum.LoginType) {
	ip := c.ClientIP()
	addr := core.GetIpAddr(ip)

	token := c.GetHeader("token")
	fmt.Println(token)
	//TODO：通过jwt获取
	userId := uint(1)
	userName := "张三"
	global.DB.Create(&models.LogModel{
		LogType:     enum.LoginLogType,
		Title:       "用户登录",
		Content:     "",
		UserId:      userId,
		Ip:          ip,
		Addr:        addr,
		LoginStatus: true,
		UserName:    userName,
		Pwd:         "-",
		LoginType:   loginType,
	})
}

func NewLoginFailed(c *gin.Context, loginType enum.LoginType, msg string, username string, pwd string) {
	ip := c.ClientIP()
	addr := core.GetIpAddr(ip)

	global.DB.Create(&models.LogModel{
		LogType:     enum.LoginLogType,
		Title:       "用户登录失败",
		Content:     msg,
		Ip:          ip,
		Addr:        addr,
		LoginStatus: true,
		UserName:    username,
		Pwd:         pwd,
		LoginType:   loginType,
	})
}
