package router

import (
	"boke-server/api"
	"github.com/gin-gonic/gin"
)

func BannerRouter(r *gin.RouterGroup) {
	var app = api.App.BannerApi
	r.POST("/banner", app.BannerCreateView)
	r.DELETE("/banner", app.BannerRemoveView)
	r.PUT("/banner", app.BannerUpdateView)
	r.GET("/banner", app.BannerList)
}
