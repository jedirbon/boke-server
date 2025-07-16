package models

type CategoryModel struct {
	Title     string    `json:"title"`
	UserId    uint      `json:"userId"`
	UserModel UserModel `gorm:"foreignKey:UserId" json:"-"`
	Model
}
