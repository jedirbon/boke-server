package redis_service

import (
	"boke-server/global"
	"boke-server/utils/jwt"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

// token失效 添加到黑名单
func TokenBlack(token string) {
	key := fmt.Sprintf("token_black_%s", token)
	claims, ok := jwt.ParseToken(token)
	if !ok {
		logrus.Errorf("token解析失败")
		return
	}
	second := claims.Exp - time.Now().Unix()
	res, err := global.Redis.Set(key, "1", time.Duration(second)*time.Second).Result()
	if err != nil {
		logrus.Errorf(err.Error())
		return
	}
	fmt.Println(res)
}

// 查找是否在黑名单
func FindTokenIsBlack(token string) {
	key := fmt.Sprintf("token_black_%s", token)
	res, err := global.Redis.Get(key).Result()
	if err != nil {
		logrus.Errorf(err.Error())
		return
	}
	fmt.Println(res)

}
