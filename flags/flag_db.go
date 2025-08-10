package flags

import (
	"boke-server/global"
	"boke-server/models"
	"github.com/sirupsen/logrus"
)

func FlagDB() {
	err := global.DB.AutoMigrate(
		&models.ArticleDiggModel{},
		&models.ArticleModel{},
		&models.CategoryModel{},
		&models.ImageModel{},
		&models.UserArticleLookHistoryModel{}, //用户浏览历史表
		&models.UserModel{},
		&models.UserConfModel{},
		&models.UserTopArticleModel{},
		&models.CommentModel{},
		&models.BannerModel{},
		&models.LogModel{},
		&models.GlobalNotificationModel{},
		&models.ArticleDiggStartModel{},
		&models.LikeModel{},
		&models.LookModel{},
	)
	if err != nil {
		logrus.Errorf("数据库迁移失败")
		return
	}
	logrus.Infof("数据库迁移成功！")
}
