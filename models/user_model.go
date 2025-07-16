package models

import "time"

type UserModel struct {
	Username       string `gorm:"size:32" json:"username" form:"username" binding:"required"`
	Nickname       string `gorm:"size:32" json:"nickname" form:"nickname" binding:"required"`
	Avatar         string `gorm:"size:256" json:"avatar" form:"avatar"`
	Abstract       string `gorm:"size:256" json:"abstract" form:"abstract"`
	RegisterSource string `json:"registerSource" form:"registerSource"` //注册来源
	CodeAge        int    `json:"codeAge" form:"codeAge"`               //马岭
	Password       string `gorm:"size:64" json:"-" binding:"required" form:"password"`
	Email          string `gorm:"size:256" json:"email" form:"email"`
	OpenID         string `gorm:"size:64" json:"openID" form:"openID"` //第三方登录的唯一ID
	RoleId         int    `json:"roleId" form:"roleId"`                //1管理员 2用户 3访客
	Model
}

type UserRegister struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Email    string `json:"email"`
}

type UserDetails struct {
	Username       string    `gorm:"size:32" json:"username" form:"username" binding:"required"`
	Nickname       string    `gorm:"size:32" json:"nickname" form:"nickname" binding:"required"`
	Avatar         string    `gorm:"size:256" json:"avatar" form:"avatar"`
	Abstract       string    `gorm:"size:256" json:"abstract" form:"abstract"`
	RegisterSource string    `json:"registerSource" form:"registerSource"` //注册来源
	CreatedAt      time.Time `json:"create"`
	UserConfModel  UserConfModel
}

type UserConfModel struct {
	UserId             uint       `gorm:"unique" json:"userId"`
	UserModel          UserModel  `gorm:"foreignKey:UserId" json:"-"`
	LikeTags           []string   `gorm:"type:longtext;serializer:json" json:"likeTags"`
	UpdateUsernameDate *time.Time `json:"UpdateUsernameDate"` //上次修改用户名时间
	OpenCollect        bool       `json:"openCollect"`        //公开我的收藏
	OpenFans           bool       `json:"OpenFans"`           //公开我的粉丝
	OpenFollow         bool       `json:"OpenFollow"`         //公开我的关注
	HomeStyleId        uint       `json:"homeStyleId"`        //样式风格Id
	Model
}
