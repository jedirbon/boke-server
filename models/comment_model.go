package models

type CommentModel struct {
	Content        string          `gorm:"size:256" json:"content"`
	UserId         uint            `json:"userId"`
	UserInfo       *UserModel      `gorm:"foreignKey:UserId" json:"userInfo"`
	ArticleId      uint            `json:"articleId"`
	ArticleModel   *ArticleModel   `gorm:"foreignKey:ArticleId" json:"-"`
	ParentId       *uint           `json:"parentId"` //父id
	ParentModel    *CommentModel   `gorm:"foreignKey:ParentId" json:"-"`
	SubCommentList []*CommentModel `gorm:"foreignKey:ParentId" json:"-"`
	RootParentId   *int            `json:"rootParentId"` //根评论id
	DiggCount      int             `json:"diggCount"`    //评论点赞数
	Model
}

type RequestCommentModel struct {
	UserId    int `json:"userId" form:"userId"`
	ArticleId int `json:"articleId" form:"articleId"`
	ParentId  int `json:"parentId" form:"parentId"`
}
