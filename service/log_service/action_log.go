package log_service

import (
	"boke-server/core"
	"boke-server/global"
	"boke-server/models"
	"boke-server/models/enum"
	"github.com/sirupsen/logrus"
)

type ActionLog struct {
	Level   enum.LogLevelType
	Title   string
	Ip      string
	Content string
}

func (ac *ActionLog) SetTitle(t string) {
	ac.Title = t
}
func (ac *ActionLog) SetLevel(l enum.LogLevelType) {
	ac.Level = l
}
func Save(ac *ActionLog) {
	addr := core.GetIpAddr(ac.Ip)
	result := global.DB.Create(&models.LogModel{
		LogType: enum.ActionLogType,
		Title:   ac.Title,
		Content: ac.Content,
		Level:   ac.Level,
		UserId:  uint(1),
		Ip:      ac.Ip,
		Addr:    addr,
	})
	if result.Error != nil {
		logrus.Errorf("日志创建失败")
	}
}
