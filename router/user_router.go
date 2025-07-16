package router

import (
	"boke-server/api"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	app := api.App.UserApi
	r.POST("user/send_email", app.SendEmailView)
	r.GET("user/details", app.ViewDetails)
	r.PUT("user/upload", app.UploadUser)
}
