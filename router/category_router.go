package router

import (
	"boke-server/api"
	"github.com/gin-gonic/gin"
)

func CateGoryRouter(r *gin.RouterGroup) {
	var app = api.App.CateGoryApi
	r.GET("category/list", app.GetGoryList)
	r.POST("category/add", app.AddCateGory)
	r.DELETE("category/del", app.DelCateGory)
	r.POST("upload/cover", app.UploadCover)
}
