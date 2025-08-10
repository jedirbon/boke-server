package user_api

import (
	"boke-server/common/res"
	"boke-server/global"
	"boke-server/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (UserApi) ViewDetails(c *gin.Context) {
	val := c.Query("userId")
	if val == "" {
		logrus.Error("获取用户id失败")
		res.FailedMsg("获取用户id失败", c)
		return
	}
	var userDetails models.UserModel
	result := global.DB.Find(&userDetails, val)
	if result.Error != nil {
		res.FailedAny(result.Error, "查询失败", c)
		return
	}
	res.OkAny(userDetails, "查询成功", c)
}
