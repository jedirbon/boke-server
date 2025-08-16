package models

type CommentDiggModel struct {
	Id     uint `json:"id" gorm:"id primaryKey;autoIncrement"`
	LikeId uint `json:"likeId"`
	UserId uint `json:"userId"`
}
