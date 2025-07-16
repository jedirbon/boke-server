package router

import (
	"boke-server/api"
	"github.com/gin-gonic/gin"
)

func CaptchaRouter(r *gin.Engine) {
	var app = api.App.CaptchaApi
	r.GET("/captcha/generate", app.CreateCapt)
	r.GET("/captcha/image/:id", app.CaptchaImage)
	r.GET("/captcha/verify", app.Verify)
}
