package like_api

import (
	"boke-server/common/res"
	"boke-server/global"
	"boke-server/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type LikeApi struct{}

func (LikeApi) Add(c *gin.Context) {
	var data models.LikeModel
	err := c.ShouldBind(&data)
	if err != nil {
		res.FailedMsg("请传入参数", c)
		return
	}
	userId, ok := c.Get("userId")
	if ok {
		data.UserId = userId.(uint)
		if data.LikeId == data.UserId {
			res.FailedMsg("禁止耍小聪明", c)
			return
		}
		result := global.DB.Create(&data)
		if result.Error != nil {
			res.FailedMsg(result.Error.Error(), c)
			return
		}
		//增加用户表的like数
		global.DB.Model(&models.UserModel{}).
			Where("id = ?", userId).
			Update("like_count", gorm.Expr("like_count + ?", 1))
		//增加被关注的用户的粉丝数
		global.DB.Model(&models.UserModel{}).
			Where("id = ?", data.LikeId).
			Update("fans_count", gorm.Expr("fans_count + ?", 1))
		res.Ok(data, c)
	}
}

func (LikeApi) Del(c *gin.Context) {
	var data models.LikeModel
	likeId := c.Query("likeId")
	if likeId == "" {
		res.FailedMsg("请传入参数", c)
		return
	}
	userId, ok := c.Get("userId")
	num, _ := strconv.ParseUint(likeId, 10, 64)
	if ok {
		data.LikeId = uint(num)
		data.UserId = userId.(uint)
		if data.LikeId == data.UserId {
			res.FailedMsg("禁止耍小聪明", c)
			return
		}
		result := global.DB.Delete(&data)
		if result.Error != nil {
			res.FailedMsg(result.Error.Error(), c)
			return
		}
		//减少用户表的like数
		global.DB.Model(&models.UserModel{}).
			Where("id = ?", userId).
			Update("like_count", gorm.Expr("like_count - ?", 1))
		//减少被关注的用户的粉丝数
		global.DB.Model(&models.UserModel{}).
			Where("id = ?", likeId).
			Update("fans_count", gorm.Expr("fans_count - ?", 1))

		res.Ok(data, c)
	}
}
