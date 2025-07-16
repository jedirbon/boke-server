package user_api

import (
	"boke-server/common/res"
	"boke-server/global"
	"boke-server/utils/email_stroe"
	"github.com/gin-gonic/gin"
)

type RegisterEmailRequest struct {
	EmailId   string `json:"emailId" binding:"required"`
	EmailCode string `json:"emailCode" binding:"required"`
	Pwd       string `json:"pwd" binding:"required"`
}

func (UserApi) RegisterEmailView(c *gin.Context) {
	var cr RegisterEmailRequest
	err := c.ShouldBind(&cr)
	if err != nil {
		res.FailedMsg("参数错误", c)
		return
	}

	value, ok := global.EmailVerifyStore.Load(cr.EmailId)
	if !ok {
		res.FailedMsg("邮箱验证失败", c)
		return
	}
	info, ok := value.(email_stroe.EmailStoreInfo)
	if !ok {
		res.FailedMsg("邮箱验证失败", c)
		return
	}
	if info.Code != cr.EmailCode {
		global.EmailVerifyStore.Delete(cr.EmailId)
		res.FailedMsg("邮箱验证失败", c)
		return
	}
}
