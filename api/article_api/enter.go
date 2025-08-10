package article_api

import (
	"boke-server/common/res"
	"boke-server/global"
	"boke-server/models"
	"boke-server/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ArticleApi struct{}

// 获取文章列表
func (ArticleApi) GetArticleList(c *gin.Context) {
	var rq models.ArticleRequestData
	var articleList []models.ResponseArticleModel
	var totalCount int64
	rq.PageIndex = 1
	rq.PageSize = 10
	err := c.ShouldBindQuery(&rq)
	if err != nil {
		res.FailedMsg(err.Error(), c)
		return
	}
	fmt.Println("*********")
	fmt.Println(rq)
	fmt.Println("*********")
	global.DB.Model(&models.ArticleModel{}).Count(&totalCount)
	result := global.SlaveDB.
		Model(&models.ArticleModel{}).
		Preload("UserInfo").
		Limit(rq.PageSize).
		Offset((rq.PageIndex-1)*rq.PageSize).
		Where(struct {
			CategoryID *uint
			UserId     *uint
		}{
			CategoryID: rq.CategoryID,
			UserId:     rq.UserId,
		}).
		Where("title like ?", utils.FormatLike(rq.Title)).
		Find(&articleList)
	if result.Error != nil {
		res.FailedMsg(result.Error.Error(), c)
		return
	}
	res.Ok(global.ResponseData[models.ResponseArticleModel]{
		Data:       articleList,
		TotalCount: totalCount,
	}, c)
}
func checkArticleAction(articleId string, userId string, status string) (bool, error) {
	var count int64
	err := global.DB.Model(models.ArticleDiggStartModel{}).
		Where("article_id = ? AND user_id = ? AND status = ?", articleId, userId, status).
		Count(&count).Error
	return count > 0, err
}
func (ArticleApi) GetArticleDetails(c *gin.Context) {
	id := c.Query("id")
	userId := c.Query("userId")
	var data models.ResponseArticleModel
	if id == "" {
		res.FailedMsg("无效的参数", c)
		return
	}
	result := global.DB.
		Preload("UserInfo").
		Model(&models.ArticleModel{}).Where("id = ?", id).First(&data)
	if result.Error != nil {
		res.FailedMsg(result.Error.Error(), c)
		return
	}
	//获取用户操作
	data.IsLike, _ = checkArticleAction(id, userId, "like")
	data.IsStart, _ = checkArticleAction(id, userId, "start")
	global.DB.Model(&models.ArticleModel{}).
		Where("id = ?", id).
		Update("look_count", gorm.Expr("look_count + ?", 1))
	res.Ok(data, c)
}

// 添加文章
func (ArticleApi) AddArticle(c *gin.Context) {
	var addData models.ArticleModel
	err := c.ShouldBind(&addData)
	if err != nil {
		res.FailedMsg(err.Error(), c)
		return
	}
	//读取用户Id
	val, ok := c.Get("userId")
	if ok {
		addData.UserId = val.(uint)
	}
	result := global.DB.Create(&addData)
	if result.Error != nil {
		res.FailedMsg(result.Error.Error(), c)
		return
	}
	global.DB.Model(&models.UserModel{}).
		Where("id = ?", addData.UserId).
		Update("article_count", gorm.Expr("article_count + ?", 1))
	res.Ok(addData, c)
}

// 删除文章
func (ArticleApi) DelArticle(c *gin.Context) {
	var dels []int
	err := c.ShouldBindJSON(dels)
	if err != nil {
		res.FailedMsg(err.Error(), c)
		return
	}
	result := global.DB.Where("id in ?", dels).Delete(&models.ArticleModel{})
	if result.Error != nil {
		res.FailedMsg(result.Error.Error(), c)
		return
	}
	if result.RowsAffected > 0 {
		val, ok := c.Get("userId")
		if ok {
			global.DB.Model(&models.UserModel{}).
				Where("id = ?", val).
				Update("article_count", gorm.Expr("article_count - ?", 1))
		}
		res.Ok(fmt.Sprintf("删除了%d条数据", result.RowsAffected), c)
	} else {
		res.FailedMsg("未找到该数据", c)
	}
}

// 修改
func (ArticleApi) UpdateArticle(c *gin.Context) {
	var editData models.ArticleModel
	err := c.ShouldBind(&editData)
	if err != nil {
		res.FailedMsg(err.Error(), c)
		return
	}
	result := global.DB.Updates(&editData)
	if result.Error != nil {
		res.FailedMsg(result.Error.Error(), c)
		return
	}
	if result.RowsAffected > 0 {
		res.Ok(editData, c)
	} else {
		res.FailedMsg("未找到该数据", c)
	}
}

// 文章操作
func (ArticleApi) Action(c *gin.Context) {
	var data models.ArticleDiggStartModel
	err := c.ShouldBind(&data)
	if err != nil {
		res.FailedMsg(err.Error(), c)
		return
	}
	val, ok := c.Get("userId")
	if ok {
		data.UserId = val.(uint)
	}
	tx := global.DB.Begin()
	var count int64
	tx.Where(&data).Find(&data).Count(&count)
	fmt.Println(count)
	if count > 0 {
		tx.Delete(&data)
		switch data.Status {
		case "like":
			{
				tx.Model(&models.ArticleModel{}).Where("id = ?", data.ArticleId).Update("DiggCount", gorm.Expr("digg_count - ?", 1))
			}
		case "start":
			{
				tx.Model(&models.ArticleModel{}).Where("id = ?", data.ArticleId).Update("CollectCount", gorm.Expr("collect_count - ?", 1))
			}
		}

	} else {
		switch data.Status {
		case "like":
			{
				tx.Model(&models.ArticleModel{}).Where("id = ?", data.ArticleId).Update("DiggCount", gorm.Expr("digg_count + ?", 1))
			}
		case "start":
			{
				tx.Model(&models.ArticleModel{}).Where("id = ?", data.ArticleId).Update("CollectCount", gorm.Expr("collect_count + ?", 1))
			}
		}
		result := tx.Save(&data)
		if result.Error != nil {
			res.FailedMsg(result.Error.Error(), c)
			tx.Rollback()
			return
		}
	}
	tx.Commit()
	res.OkMsg("操作成功", c)
}

func (ArticleApi) GetTitleList(c *gin.Context) {
	title := c.Query("title")
	if title == "" {
		res.FailedMsg("无效字符", c)
		return
	}
	var titles []string
	result := global.DB.Model(models.ArticleModel{}).
		Select("title").
		Where("title LIKE ?", utils.FormatLike(title)).Find(&titles)
	if result.Error != nil {
		res.FailedMsg(result.Error.Error(), c)
		logrus.Error(result.Error.Error())
		return
	}
	res.Ok(titles, c)
}
