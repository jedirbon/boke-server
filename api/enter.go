package api

import (
	"boke-server/api/article_api"
	"boke-server/api/banner_api"
	"boke-server/api/captcha_api"
	"boke-server/api/category_api"
	"boke-server/api/comment_api"
	"boke-server/api/image_api"
	"boke-server/api/like_api"
	"boke-server/api/log_api"
	"boke-server/api/look_api"
	"boke-server/api/site_api"
	"boke-server/api/user_api"
)

type Api struct {
	SiteApi     site_api.SiteApi
	LogApi      log_api.LogApi
	UserApi     user_api.UserApi
	ImageApi    image_api.ImageApi
	BannerApi   banner_api.BannerApi
	CaptchaApi  captcha_api.CaptchaApi
	CateGoryApi category_api.CateGory
	ArticleApi  article_api.ArticleApi
	CommentApi  comment_api.CommentApi
	LikeApi     like_api.LikeApi
	LookApi     look_api.LookApi
}

var App = Api{}
