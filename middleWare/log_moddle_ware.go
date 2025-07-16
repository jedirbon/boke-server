package middleWare

import (
	"boke-server/service/log_service"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
)

type ResponseWrite struct {
	gin.ResponseWriter
}

func LogMiddleWare(c *gin.Context) {
	//请求中间件
	byteData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logrus.Errorf(err.Error())
	}
	fmt.Println("body:", string(byteData))
	c.Request.Body = io.NopCloser(bytes.NewReader(byteData))
	var ac = log_service.ActionLog{
		Level:   1,
		Title:   "你好",
		Ip:      c.ClientIP(),
		Content: string(byteData),
	}
	log_service.Save(&ac)
	c.Next()
	//响应中间件
}
