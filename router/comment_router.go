package router

import (
	"boke-server/api"
	"github.com/gin-gonic/gin"
)

func CommentRouter(r *gin.RouterGroup) {
	app := api.App.CommentApi
	r.POST("/comment/add", app.Add)
	r.DELETE("/comment/delete", app.Del)
	r.PUT("/comment/edit", app.Edit)
	r.GET("/comment/list", app.GetList)
}
