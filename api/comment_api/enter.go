package comment_api

import (
	"boke-server/common/res"
	"boke-server/global"
	"boke-server/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentApi struct{}

func (CommentApi) Add(c *gin.Context) {
	var data models.CommentModel
	err := c.ShouldBind(&data)
	if err != nil {
		res.FailedMsg(err.Error(), c)
		return
	}
	val, ok := c.Get("userId")
	if ok {
		data.UserId = val.(uint)
	} else {
		res.FailedMsg("请先登录", c)
		return
	}
	result := global.DB.Create(&data)
	if result.Error != nil {
		res.FailedMsg(result.Error.Error(), c)
		return
	}
	//评论数增加
	global.DB.Model(&models.ArticleModel{}).
		Where("id = ?", data.ArticleId).
		Update("comment_count", gorm.Expr("comment_count + ?", 1))
	res.Ok(data, c)
}

func (CommentApi) Del(c *gin.Context) {
	var ids []int
	err := c.ShouldBind(&ids)
	if err != nil {
		res.FailedMsg(err.Error(), c)
		return
	}
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	//在事务中执行操作
	result := tx.Where("id in ?").Delete(&models.CommentModel{})
	if result.Error != nil {
		res.FailedMsg(result.Error.Error(), c)
		tx.Rollback()
		return
	}
	res.Ok(fmt.Sprintf("删除了%d条数据", result.RowsAffected), c)
	tx.Commit()
}
func (CommentApi) Edit(c *gin.Context) {
	var data models.CommentModel
	err := c.ShouldBind(&data)
	if err != nil {
		res.FailedMsg(err.Error(), c)
		return
	}

	result := global.DB.Updates(&data)
	if result.Error != nil {
		res.FailedMsg(result.Error.Error(), c)
		return
	}

	res.Ok(data, c)
}

type requestData struct {
	global.PageResponse
	models.RequestCommentModel
}

func (CommentApi) GetList(c *gin.Context) {
	var rq requestData
	var commentList []models.CommentModel
	var totalCount int64
	rq.PageIndex = 1
	rq.PageSize = 10
	err := c.ShouldBind(&rq)
	if err != nil {
		res.FailedMsg(err.Error(), c)
		return
	}
	global.DB.Model(&models.CommentModel{}).Count(&totalCount)
	fmt.Println(rq)
	result := global.DB.
		Preload("UserInfo").
		Model(&models.CommentModel{}).
		Where(&rq.RequestCommentModel).
		Limit(rq.PageSize).
		Offset((rq.PageIndex - 1) * rq.PageSize).
		Find(&commentList)
	if result.Error != nil {
		res.FailedMsg(result.Error.Error(), c)
		return
	}
	res.Ok(global.ResponseData[models.CommentModel]{
		Data:       commentList,
		TotalCount: totalCount,
	}, c)
}
