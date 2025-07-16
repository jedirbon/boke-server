package main

import (
	"boke-server/core"
	"boke-server/flags"
	"boke-server/global"
	"boke-server/utils/jwt"
	"fmt"
)

func main() {
	flags.Parse()
	//读取配置信息
	global.Config = core.ReadConf()
	var userInfo = jwt.UserInfo{
		UserId: 1,
		RoleId: 2,
	}
	token, err := jwt.CreateToken(userInfo)
	if err != nil {
		fmt.Println("创建jwt失败")
		return
	}
	fmt.Println(token)
	fmt.Println("\n")
	claims, ok := jwt.ParseToken(token)
	if !ok {
		fmt.Println("解析token失败")
		return
	}
	fmt.Println("\n")
	fmt.Println(claims)
}
