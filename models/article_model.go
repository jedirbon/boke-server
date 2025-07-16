package models

type ArticleModel struct {
	Title        string    `gorm:"size:32" json:"title"`
	Abstract     string    `gorm:"size:256" json:"abstract"`
	Content      string    `json:"content"`
	CategoryID   uint      `json:"categoryID"`
	TagList      []string  `gorm:"type:longtext;serializer:json" json:"tagList"`
	Cover        string    `gorm:"size:256" json:"cover"`
	UserId       uint      `json:"userId"`
	UserModel    UserModel `gorm:"foreignKey:UserId" json:"-"`
	LookCount    int       `json:"lookCount"`
	DiggCount    int       `json:"diggCount"`
	CommentCount int       `json:"commentCount"`
	CollectCount int       `json:"collectCount"`
	OpenComment  bool      `json:"openComment"` //打开评论
	Status       int8      `json:"status"`      //状态 草稿 审核中 已发布
	Model
}
