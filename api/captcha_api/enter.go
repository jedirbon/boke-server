package captcha_api

import (
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CaptchaApi struct {
}

func (CaptchaApi) CreateCapt(c *gin.Context) {
	// 生成验证码（默认6位数字）
	captchaID := captcha.New()

	// 返回验证码ID和图片URL
	c.JSON(http.StatusOK, gin.H{
		"captcha_id": captchaID,
		"image_url":  "/uploads/image/" + captchaID + ".png",
	})
}

func (CaptchaApi) CaptchaImage(c *gin.Context) {
	captchaID := c.Param("id")
	if captchaID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Captcha ID is required"})
		return
	}

	// 设置响应头为PNG图片
	c.Header("Content-Type", "image/png")

	// 写入验证码图片到响应
	captcha.WriteImage(c.Writer, captchaID, 120, 80) // 宽度120px，高度80px
}

func (CaptchaApi) Verify(c *gin.Context) {
	type VerifyRequest struct {
		CaptchaID string `json:"captcha_id" binding:"required"`
		Input     string `json:"input" binding:"required"`
	}

	var req VerifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// 校验验证码
	if captcha.VerifyString(req.CaptchaID, req.Input) {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "验证码正确"})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "验证码错误"})
	}
}
