package models

type ArticleDiggModel struct {
	UserId       uint         `gorm:"uniqueIndex:idx_name" json:"userId"`
	UserModel    UserModel    `gorm:"foreignKey:UserId" json:"-"`
	ArticleId    uint         `gorm:"uniqueIndex:idx_name" json:"articleId"`
	ArticleModel ArticleModel `gorm:"foreignKey:ArticleId" json:"-"`
	Model
}
