package image_api

import (
	"boke-server/common/res"
	"boke-server/global"
	"boke-server/service/qiniu_service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type QiNiuGenTokenResponse struct {
	Token  string `json:"token"`
	Key    string `json:"key"`
	Region string `json:"region"`
	Url    string `json:"url"`
	Size   int64  `json:"size"`
}

func (ImageApi) QiNiuToken(c *gin.Context) {
	q := global.Config.QiNiu
	if !q.Enable {
		res.FailedMsg("未启用七牛云配置！", c)
		return
	}
	token, err := qiniu_service.GetQiNiuToken()
	if err != nil {
		res.FailedMsg("获取七牛token失败"+err.Error(), c)
		return
	}
	uid := uuid.New().String()
	key := fmt.Sprintf("%s,%s.png", q.Prefix, uid)
	url := fmt.Sprintf("%s,%s", q.Uri, key)
	res.Ok(QiNiuGenTokenResponse{
		Token:  token,
		Key:    key,
		Region: q.Region,
		Url:    url,
		Size:   q.Size,
	}, c)
}
