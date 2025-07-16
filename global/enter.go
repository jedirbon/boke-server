package global

import (
	"boke-server/conf"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"sync"
)

const Version = "10.0.1"

var (
	Config           *conf.Config
	DB               *gorm.DB
	Redis            *redis.Client
	EmailVerifyStore = sync.Map{}
)
