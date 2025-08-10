package models

import "gorm.io/gorm"

type LookModel struct {
	UserId uint `json:"userId" form:"userId"`
	LookId uint `json:"lookId" form:"lookId" binding:"required"`
	gorm.Model
}
