package router

import (
	"boke-server/api"
	"github.com/gin-gonic/gin"
)

func SiteRouter(r *gin.RouterGroup) {
	app := api.App.SiteApi

	r.GET("/site/qq_url", app.SiteInfoQQView)
	r.GET("/site/:name", app.SiteInfoView)
	r.PUT("/site/:name", app.SiteUpdateView)
}
