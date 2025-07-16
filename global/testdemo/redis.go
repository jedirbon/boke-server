package main

import (
	"boke-server/core"
	"boke-server/flags"
	"boke-server/global"
	"boke-server/service/redis_service"
	"boke-server/utils/jwt"
	"fmt"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.InitLogrus()
	global.Redis = core.InitRedis()

	token, err := jwt.CreateToken(jwt.UserInfo{
		UserId: 1,
		RoleId: 2,
	})
	fmt.Println(token, err)
	redis_service.TokenBlack(token)
	redis_service.FindTokenIsBlack(token)
}
