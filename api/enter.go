package api

import (
	"boke-server/api/banner_api"
	"boke-server/api/captcha_api"
	"boke-server/api/image_api"
	"boke-server/api/log_api"
	"boke-server/api/site_api"
	"boke-server/api/user_api"
)

type Api struct {
	SiteApi    site_api.SiteApi
	LogApi     log_api.LogApi
	UserApi    user_api.UserApi
	ImageApi   image_api.ImageApi
	BannerApi  banner_api.BannerApi
	CaptchaApi captcha_api.CaptchaApi
}

var App = Api{}
