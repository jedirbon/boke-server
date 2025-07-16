package user_api

import (
	"boke-server/common/res"
	"boke-server/global"
	"boke-server/models"
	"boke-server/utils/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserApi struct {
}
type LoginForm struct {
	UserName string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func (UserApi) Register(c *gin.Context) {
	var registerForm models.UserRegister
	err := c.ShouldBind(&registerForm)
	if err != nil {
		res.FailedMsg(err.Error(), c)
		return
	}
	password, err := DecryptPassword(registerForm.Password)
	if err != nil {
		res.FailedMsg("密码解密失败", c)
		return
	}
	registerForm.Password = password
	result := global.DB.Where("username = ?", registerForm.Username).Find(&models.UserModel{})
	if result.RowsAffected > 0 && result.Error == nil {
		res.FailedMsg("该账号已存在", c)
		return
	}
	result = global.DB.Create(&models.UserModel{
		Username: registerForm.Username,
		Password: password,
		Nickname: registerForm.Nickname,
		Email:    registerForm.Email,
	})
	if result.RowsAffected > 0 && result.Error == nil {
		res.OkMsg("创建成功", c)
		return
	} else {
		res.FailedAny(result.Error, "创建失败", c)
	}
}

func (UserApi) Login(c *gin.Context) {
	var loginForm LoginForm
	var userInfo models.UserModel
	err := c.ShouldBind(&loginForm)
	if err != nil {
		res.FailedMsg("参数错误", c)
		return
	}

	realPassword, err := DecryptPassword(loginForm.Password)
	if err != nil {
		res.FailedMsg("密码解密失败", c)
		return
	}
	result := global.DB.
		Where("username = ? and password = ?", loginForm.UserName, realPassword).
		Find(&userInfo)
	if result.RowsAffected > 0 && result.Error == nil {
		fmt.Println(userInfo.ID, userInfo.RoleId)
		token, err := jwt.CreateToken(jwt.UserInfo{
			UserId: userInfo.ID,
			RoleId: userInfo.RoleId,
		})
		if err == nil {
			res.OkAny(struct {
				UserInfo models.UserModel `json:"userInfo"`
				Token    string           `json:"token"`
			}{
				UserInfo: userInfo,
				Token:    token,
			}, "登录成功", c)
		} else {
			res.FailedAny(err, "创建token失败", c)
		}
	} else {
		res.FailedAny(result.Error, "账号密码错误", c)
	}
}

func (UserApi) UploadUser(c *gin.Context) {
	var userInfo models.UserModel
	err := c.ShouldBind(&userInfo)
	if err != nil {
		res.FailedMsg("参数错误", c)
		return
	}
	result := global.DB.Model(&models.UserModel{}).Take(userInfo)
	if result.RowsAffected > 0 {

	}
}
