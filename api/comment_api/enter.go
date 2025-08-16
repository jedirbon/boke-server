package comment_api

import (
	"boke-server/common/res"
	"boke-server/global"
	"boke-server/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
	global.DB.Model(&models.CommentModel{}).
		Where("id = ? or id = ?", data.RootParentId, data.ParentId).
		Update("son_count", gorm.Expr("son_count + ?", 1))
	res.Ok(data, c)
}

func (CommentApi) Del(c *gin.Context) {
	//在事务中执行操作
	type rq struct {
		ArticleId    uint  `json:"articleId"`
		Id           uint  `json:"id"`
		ParentId     *uint `json:"parentId"`
		RootParentId *uint `json:"rootParentId"`
	}
	var request rq
	err := c.ShouldBind(&request)
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

	var result *gorm.DB
	//判断是一级还是二级还是回复
	if request.ParentId == nil {
		//一级
		result = tx.Where("root_parent_id = ?", request.Id).Delete(&models.CommentModel{})
	} else {
		result = tx.Where("parent_id = ? or id = ?", request.Id, request.Id).Delete(&models.CommentModel{})
		tx.Model(&models.CommentModel{}).
			Where("id = ? or id = ?", request.ParentId, request.RootParentId).
			Update("son_count", gorm.Expr("son_count - ?", result.RowsAffected))
	}
	if result.Error != nil {
		res.FailedMsg(result.Error.Error(), c)
		tx.Rollback()
		return
	}
	//评论数减少
	tx.Model(&models.ArticleModel{}).
		Where("id = ?", request.ArticleId).
		Update("comment_count", gorm.Expr("comment_count - ?", result.RowsAffected))
	//返回值
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

func getCommentTree(model *models.CommentModel) *models.CommentModel {
	// 递归处理子评论，确保每个子评论都预加载了UserInfo
	for i := 0; i < len(model.SubCommentList); i++ {
		// 为每个子评论单独预加载UserInfo和子评论
		global.DB.Preload("UserInfo").Find(&model.SubCommentList[i])

		// 递归处理更深层的子评论
		if len(model.SubCommentList[i].SubCommentList) > 0 {
			getCommentTree(model.SubCommentList[i])
		}
	}
	return model
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
	global.DB.Model(&models.CommentModel{}).Where("parent_id is null").Count(&totalCount)
	fmt.Println(rq)
	result := global.DB.
		Model(&models.CommentModel{}).
		Preload("UserInfo").
		Preload("SubCommentList", func(db *gorm.DB) *gorm.DB {
			return db.Limit(1).Order("created_at")
		}).Find(&models.CommentModel{}).
		Where(&rq.RequestCommentModel).
		Where("parent_id is null").
		Limit(rq.PageSize).
		Offset((rq.PageIndex - 1) * rq.PageSize).
		Find(&commentList)
	if result.Error != nil {
		res.FailedMsg(result.Error.Error(), c)
		return
	}
	for _, model := range commentList {
		getCommentTree(&model)
	}
	res.Ok(global.ResponseData[models.CommentModel]{
		Data:       commentList,
		TotalCount: totalCount,
	}, c)
}

func (CommentApi) GetSonList(c *gin.Context) {
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
	global.DB.Model(&models.CommentModel{}).
		Where("parent_id is not null").
		Where("root_parent_id = ?", rq.RootParentId).Count(&totalCount)
	fmt.Println(rq)
	result := global.DB.
		Model(&models.CommentModel{}).
		Preload("UserInfo").
		Where(&rq.RequestCommentModel).
		Where("parent_id is not null").
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

func (CommentApi) Like(c *gin.Context) {
	var request models.CommentDiggModel
	err := c.ShouldBind(&request)
	if err != nil {
		res.FailedMsg(err.Error(), c)
		logrus.Error(err.Error())
	}
	userId, ok := c.Get("userId")
	if ok {
		request.UserId = userId.(uint)
		result := global.DB.Where("like_id = ? AND user_id = ?", request.LikeId, request.UserId).First(&request)
		if request.Id > 0 && result.Error == nil {
			global.DB.Delete(request)
			global.DB.Model(&models.CommentModel{}).
				Where("id = ?", request.LikeId).
				Update("digg_count", gorm.Expr("digg_count - ?", 1))
			res.OkMsg("取消点赞", c)
		} else {
			global.DB.Create(&request)
			global.DB.Model(&models.CommentModel{}).
				Where("id = ?", request.LikeId).
				Update("digg_count", gorm.Expr("digg_count + ?", 1))
			res.Ok(request, c)
		}
	}
}
