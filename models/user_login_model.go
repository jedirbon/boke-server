package models

type UserLoginModel struct {
	UserId    uint      `json:"userId"`
	UserModel UserModel `gorm:"foreignKey:UserId" json:"-"`
	Ip        string    `gorm:"size:32" json:"ip"`
	Addr      string    `gorm:"size:64" json:"addr"`
	Ua        string    `gorm:"size:128" json:"ua"`
	Model
}
