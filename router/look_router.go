package router

import (
	"boke-server/api"
	"github.com/gin-gonic/gin"
)

func LookRouter(r *gin.RouterGroup) {
	Api := api.App.LookApi
	r.POST("look/add", Api.Add)
}
