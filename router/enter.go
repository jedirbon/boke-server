package router

import (
	"boke-server/core"
	"boke-server/global"
	"boke-server/middleWare"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	//配置静态文件路劲
	r.Static("/static", "static")
	r.Static("/uploads", "uploads")
	//配置跨域问题
	r.Use(core.InitCors())
	LoginRouter(r)
	CaptchaRouter(r)
	api := r.Group("/api")
	//api.Use(middleWare.LogMiddleWare)
	api.Use(middleWare.CheckToken)
	UserRouter(api)
	SiteRouter(api)
	LogRouter(api)
	ImageRouter(api)
	BannerRouter(api)
	CateGoryRouter(api)
	ArticleRouter(api)
	CommentRouter(api)
	LikeRouter(api)
	LookRouter(api)
	addr := global.Config.System.Addr()
	r.Run(addr)
}
