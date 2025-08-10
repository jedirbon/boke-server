package user_api

import (
	"boke-server/common/res"
	"boke-server/global"
	"boke-server/models"
	"boke-server/requestModel"
	"boke-server/utils"
	"github.com/gin-gonic/gin"
)

func (UserApi) GetDiggArticle(c *gin.Context) {
	var params requestModel.RequestParams
	var articleList []models.ArticleModel
	var data global.ResponseData[models.ArticleModel]
	var totalCount int64
	err := c.ShouldBind(&params)
	if err != nil {
		res.FailedMsg(err.Error(), c)
		return
	}
	userId, ok := c.Get("userId")
	if !ok {
		res.ExpireMsg("请登录", c)
		return
	}
	result := global.DB.Model(&models.ArticleModel{}).
		Joins("LEFT JOIN article_digg_start_models as digg on digg.article_id = article_models.id").
		Where("digg.user_id = ? AND digg.status = ? AND title LIKE ?", userId, params.Status, utils.FormatLike(params.Title)).
		Count(&totalCount).
		Scopes(utils.Paginate(params.PageIndex, params.PageSize)).
		Find(&articleList)
	if result.Error != nil {
		res.FailedMsg(result.Error.Error(), c)
		return
	}
	data.Data = articleList
	data.TotalCount = totalCount
	res.Ok(data, c)
}
