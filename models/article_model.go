package models

type ArticleModel struct {
	Title        string   `gorm:"size:64" json:"title"`                         //标题
	Abstract     string   `gorm:"size:256" json:"abstract"`                     //描述
	Content      string   `json:"content"`                                      //正文
	CategoryID   uint     `json:"categoryID"`                                   //类别ID
	TagList      []string `gorm:"type:longtext;serializer:json" json:"tagList"` //标签列表
	Cover        string   `gorm:"size:256" json:"cover"`                        //封面
	UserId       uint     `json:"userId"`                                       //用户id
	LookCount    int      `json:"lookCount"`                                    //阅读量
	DiggCount    int      `json:"diggCount"`                                    //点赞数
	CommentCount int      `json:"commentCount"`                                 //评论数量
	CollectCount int      `json:"collectCount"`                                 //收藏量
	OpenComment  bool     `json:"openComment"`                                  //打开评论
	Status       int8     `json:"status"`                                       //状态 草稿 审核中 已发布
	IsLike       bool     `json:"isLike"`                                       //是否点赞
	IsStart      bool     `json:"isStart"`                                      //是否收藏
	Model
}

type ArticleRequestData struct {
	PageIndex  int      `json:"pageIndex" form:"pageIndex"`
	PageSize   int      `json:"pageSize" form:"pageSize"`
	Title      string   `json:"title" form:"title"`
	TagList    []string `json:"tagList" form:"tagList"`
	CategoryID *uint    `json:"categoryID" form:"categoryID"` //类别ID
	UserId     *uint    `json:"userId" form:"userId"`         //用户id查询
	Abstract   string   `json:"abstract" form:"abstract"`
}
type ResponseArticleModel struct {
	Title        string    `gorm:"size:32" json:"title"`                         //标题
	Abstract     string    `gorm:"size:256" json:"abstract"`                     //描述
	Content      string    `json:"content"`                                      //正文
	CategoryID   uint      `json:"categoryID"`                                   //类别ID
	TagList      []string  `gorm:"type:longtext;serializer:json" json:"tagList"` //标签列表
	Cover        string    `gorm:"size:256" json:"cover"`                        //封面
	UserId       uint      `json:"userId"`
	UserInfo     UserModel `json:"userInfo" gorm:"foreignKey:UserId"` //用户信息
	LookCount    int       `json:"lookCount"`                         //阅读量
	DiggCount    int       `json:"diggCount"`                         //
	CommentCount int       `json:"commentCount"`                      //评论数量
	CollectCount int       `json:"collectCount"`                      //收藏量
	OpenComment  bool      `json:"openComment"`                       //打开评论
	Status       int8      `json:"status"`                            //状态 草稿 审核中 已发布
	IsLike       bool      `json:"isLike"`                            //是否点赞
	IsStart      bool      `json:"isStart"`                           //是否收藏
	Model
}
