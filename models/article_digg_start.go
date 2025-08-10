package models

import (
	"database/sql"
	"time"
)

type ActionType string

const (
	Like  ActionType = "like"
	Start ActionType = "start"
)

type ArticleDiggStartModel struct {
	ID        uint         `gorm:"primarykey" json:"id"`
	CreatedAt time.Time    `json:"createAt"`
	UpdatedAt time.Time    `json:"updateAt"`
	DeletedAt sql.NullTime `gorm:"index"`
	ArticleId uint         `json:"articleId"`
	UserId    uint         `json:"userId"`
	Status    ActionType   `gorm:"type:ENUM('like','start')" json:"status" binding:"required"`
}
