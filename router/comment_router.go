package router

import (
	"boke-server/api"
	"github.com/gin-gonic/gin"
)

func CommentRouter(r *gin.RouterGroup) {
	app := api.App.CommentApi
	r.POST("/comment/add", app.Add)
	r.POST("/comment/delete", app.Del)
	r.PUT("/comment/edit", app.Edit)
	r.GET("/comment/list", app.GetList)
	r.POST("/comment/like", app.Like)
	r.GET("/comment/son/list", app.GetSonList)
}
