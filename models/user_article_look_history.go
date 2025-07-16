package models

type UserArticleLookHistoryModel struct {
	UserId       uint         `json:"userId"`
	UserModel    UserModel    `gorm:"foreignKey:UserId" json:"-"`
	ArticleId    uint         `json:"articleId"`
	ArticleModel ArticleModel `gorm:"foreignKey:ArticleId" json:"_"`
	Model
}
