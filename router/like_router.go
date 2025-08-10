package router

import (
	"boke-server/api"
	"github.com/gin-gonic/gin"
)

func LikeRouter(r *gin.RouterGroup) {
	Api := api.App.LikeApi
	r.POST("like/add", Api.Add)
	r.DELETE("like/del", Api.Del)
}
