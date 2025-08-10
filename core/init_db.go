package core

import (
	"boke-server/global"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() (db *gorm.DB) {
	dc := global.Config.DB
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open(mysql.Open(dc.DSN()), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, //不生成实体外键
		Logger:                                   logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		logrus.Fatalf("数据库连接失败:", err)
	}
	logrus.Infof("数据库链接成功！")
	return
}

func InitSlave() (db *gorm.DB) {
	dc := global.Config.SlaveDB
	//链接mysql
	db, err := gorm.Open(mysql.Open(dc.SlaveDSN()), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, //不生成实体外键
		Logger:                                   logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err)
		logrus.Fatalf("从库链接失败！", err)
	}
	logrus.Infof("从库链接成功！")
	return
}
