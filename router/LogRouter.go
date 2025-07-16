package router

import (
	"boke-server/api"
	"github.com/gin-gonic/gin"
)

func LogRouter(r *gin.RouterGroup) {
	app := api.App.LogApi
	r.GET("getloglist", app.LogList)
	r.DELETE("deletelog", app.DeleteLog)
}
