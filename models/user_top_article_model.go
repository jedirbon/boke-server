package models

type UserTopArticleModel struct {
	UserId       uint         `gorm:"uniqueIndex:idx_name" json:"userId"`
	ArticleId    uint         `gorm:"uniqueIndex:idx_name" json:"articleId"`
	UserModel    UserModel    `gorm:"foreignKey:UserId" json:"-"`
	ArticleModel ArticleModel `gorm:"foreignKey:ArticleId" json:"-"`
	Model
}
