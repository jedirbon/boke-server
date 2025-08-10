package router

import (
	"boke-server/api"
	"github.com/gin-gonic/gin"
)

func LoginRouter(r *gin.Engine) {
	app := api.App.UserApi
	r.POST("/login", app.Login)
	r.POST("/logout", app.Logout)
	r.POST("/register", app.Register)
	r.GET("/publicKey", app.GetPublicKey)
}
