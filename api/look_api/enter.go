package look_api

import (
	"boke-server/common/res"
	"boke-server/global"
	"boke-server/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LookApi struct{}

func (LookApi) Add(c *gin.Context) {
	var data models.LookModel
	err := c.ShouldBind(&data)
	if err != nil {
		res.FailedMsg("请传入参数", c)
		return
	}
	userId, ok := c.Get("userId")
	if ok {
		data.UserId = userId.(uint)
		//自己看自己
		if data.LookId == data.UserId {
			return
		}
		result := global.DB.Create(&data)
		if result.Error != nil {
			res.FailedMsg(result.Error.Error(), c)
			return
		}
		//增加用户表的look数
		global.DB.Model(&models.UserModel{}).
			Where("id = ?", data.LookId).
			Update("look_count", gorm.Expr("look_count + ?", 1))
		res.Ok(data, c)
	}
}
