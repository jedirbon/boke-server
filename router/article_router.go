package router

import (
	"boke-server/api"
	"github.com/gin-gonic/gin"
)

func ArticleRouter(r *gin.RouterGroup) {
	var app = api.App.ArticleApi
	r.GET("article/list", app.GetArticleList)
	r.POST("article/add", app.AddArticle)
	r.DELETE("article/del", app.DelArticle)
	r.PUT("article/update", app.UpdateArticle)
	r.GET("article/details", app.GetArticleDetails)
	r.GET("article/titles", app.GetTitleList)

	//文章操作
	r.POST("article/action", app.Action)
}
