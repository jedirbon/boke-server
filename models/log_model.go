package models

import "boke-server/models/enum"

type LogModel struct {
	LogType     enum.LogType      `json:"logType"` //日志类型 1 2 3
	Title       string            `gorm:"size:32" json:"title"`
	Content     string            `gorm:"size:32" json:"content"`
	Level       enum.LogLevelType `form:"level" json:"level"`
	UserId      uint              `json:"userId"`
	UserModel   UserModel         `gorm:"foreignKey:UserId" json:"-" binding:"-"` //用户信息
	Ip          string            `gorm:"size:32" json:"ip"`
	Addr        string            `gorm:"size:64" json:"addr"`
	IsRead      bool              `json:"isRead"`      //是否读取
	LoginStatus bool              `json:"loginStatus"` //登录的状态
	UserName    string            `json:"userName"`    //登录日志的用户名
	Pwd         string            `json:"pwd"`         //登录的日志密码
	LoginType   enum.LoginType    `json:"loginType"`   //登录的类型
	Model
}
