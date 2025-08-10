package user_api

import (
	"boke-server/common/res"
	"boke-server/global"
	"boke-server/models"
	"boke-server/utils/jwt"
	"fmt"
	"strings"

	"boke-server/service/redis_service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
func (UserApi) Logout(c *gin.Context) {
	// 获取Authorization header中的token
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		res.FailedMsg("未找到token", c)
		return
	}

	// 解析Bearer token
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		res.FailedMsg("token格式错误", c)
		return
	}

	token := parts[1]

	// 验证token是否有效
	_, ok := jwt.ParseToken(token)
	if !ok {
		res.FailedMsg("token无效", c)
		return
	}

	// 将token加入黑名单，使其失效
	redis_service.TokenBlack(token)

	// 记录退出日志（可选，这里注释掉避免未使用变量）
	// claims, _ := jwt.ParseToken(token)
	// userId := claims.UserInfo.UserId
	// roleId := claims.UserInfo.RoleId
	// log_service.RecordUserAction(userId, "logout", "用户退出登录", c.ClientIP())
	logrus.Info("用户退出登录")

	// 返回成功响应
	res.OkMsg("退出登录成功", c)
}

// 修改用户信息
func (UserApi) UploadUser(c *gin.Context) {
	var userInfo models.UserModel
	err := c.ShouldBind(&userInfo)
	if err != nil {
		res.FailedMsg(err.Error(), c)
		return
	}
	file, err := c.FormFile("file")
	if file != nil {
		if err != nil {
			res.FailedMsg(err.Error(), c)
		}
		//保存头像
		userInfo.Avatar = global.SaveAvatar(file, file.Filename, c)
	}
	result := global.DB.Model(&userInfo).Updates(userInfo)
	if result.RowsAffected > 0 && result.Error == nil {
		res.OkAny(userInfo, "更新成功", c)
		return
	} else {
		res.FailedMsg(err.Error(), c)
	}
}
