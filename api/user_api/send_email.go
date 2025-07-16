package user_api

import (
	"boke-server/common/res"
	"boke-server/global"
	"boke-server/models"
	"boke-server/service/email_service"
	"boke-server/utils/email_stroe"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type SendEmailRequest struct {
	Type  int8   `json:"type" binding:"oneof=1 2"` //1注册，2重置
	Email string `json:"email" binding:"required"`
}

func (UserApi) SendEmailView(c *gin.Context) {
	var cr SendEmailRequest
	err := c.ShouldBind(&cr)
	if err != nil {
		res.FailedMsg("参数错误", c)
		return
	}
	captcha.NewLen(4)
	code := captcha.New()
	id := uuid.New().String()
	switch cr.Type {
	case 1:
		result := global.DB.Take(&models.UserModel{}, "email = ?", cr.Email)
		if result.Error == nil {
			res.FailedMsg("该邮箱已经使用", c)
			return
		}
		fmt.Println(code)
		err = email_service.SendRegisterCode(cr.Email, code)
	case 2:
		err = email_service.SendResetPwdCode(cr.Email, code)
	}
	if err != nil {
		logrus.Errorf("邮件发送失败%s", err)
		res.FailedMsg("邮件发送失败", c)
		return
	}
	global.EmailVerifyStore.Store(id, email_stroe.EmailStoreInfo{
		Email: cr.Email,
		Code:  code,
	})
	res.Ok(id, c)
}
