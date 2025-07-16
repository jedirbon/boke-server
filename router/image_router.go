package router

import (
	"boke-server/api"
	"github.com/gin-gonic/gin"
)

func ImageRouter(r *gin.RouterGroup) {
	app := api.App.ImageApi

	r.POST("/images", app.ImageUploadView)
	r.DELETE("/deleteImages", app.OnUploadImage)
	r.POST("/images/qiniu", app.QiNiuToken)
}
