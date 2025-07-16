package models

type CommentModel struct {
	Content        string          `gorm:"size256" json:"content"`
	UserId         uint            `json:"userId"`
	UserModel      UserModel       `gorm:"foreignKey:UserId" json:"-"`
	ArticleId      uint            `json:"articleId"`
	ArticleModel   ArticleModel    `gorm:"foreignKey:ArticleId" json:"-"`
	ParentId       *uint           `json:"parentId"` //父id
	ParentModel    *CommentModel   `gorm:"foreignKey:ParentId" json:"-"`
	SubCommentList []*CommentModel `gorm:"foreignKey:ParentId" json:"-"`
	RootParentId   *int            `json:"rootParentId"` //根评论id
	DiggCount      int             `json:"diggCount"`    //评论点赞数
	Model
}
