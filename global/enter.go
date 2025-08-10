package global

import (
	"boke-server/common/res"
	"boke-server/conf"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"mime/multipart"
	"sync"
)

const Version = "10.0.1"

var (
	Config           *conf.Config
	DB               *gorm.DB
	SlaveDB          *gorm.DB
	Redis            *redis.Client
	EmailVerifyStore = sync.Map{}
)

// 保存头像
func SaveAvatar(file *multipart.FileHeader, filename string, c *gin.Context) string {
	dst := "uploads/avatar/" + filename
	err := c.SaveUploadedFile(file, dst)
	if err != nil {
		res.FailedMsg(err.Error(), c)
		return ""
	}
	return dst
}
