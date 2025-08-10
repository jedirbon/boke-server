package models

import "gorm.io/gorm"

type LikeModel struct {
	UserId uint `json:"userId" form:"userId"`
	LikeId uint `json:"likeId" form:"likeId" binding:"required"`
	gorm.Model
}
