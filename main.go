package main

import (
	"boke-server/api/user_api"
	"boke-server/core"
	"boke-server/flags"
	"boke-server/global"
	"boke-server/router"
)

func main() {
	flags.Parse()
	//读取配置信息
	global.Config = core.ReadConf()
	//初始话日志
	core.InitLogrus()
	//初始化数据库
	global.DB = core.InitDB()
	//global.SlaveDB = core.InitSlave()
	global.Redis = core.InitRedis()
	//迁移数据库
	flags.Run()
	//初始化Redis
	//初始化ip地址数据库
	core.InitIPDB()
	//生成密钥
	user_api.CreateKeys()
	//启动Web程序
	router.Run()
}
