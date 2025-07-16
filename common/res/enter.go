package res

import "github.com/gin-gonic/gin"

type Code int

const (
	SuccessCode     Code = 200
	FailValidCode   Code = 400
	FailServiceCode Code = 500
	TokenExpire     Code = 401
)

type Response struct {
	Code Code   `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

var empty = make(map[string]interface{})

func (r Response) JSON(c *gin.Context) {
	c.JSON(200, r)
}
func (c Code) String() string {
	switch c {
	case SuccessCode:
		return "成功"
	case FailValidCode:
		return "校验失败"
	case FailServiceCode:
		return "服务异常"
	}
	return ""
}

func Ok(data any, c *gin.Context) {
	Response{
		Code: SuccessCode,
		Data: data,
		Msg:  SuccessCode.String(),
	}.JSON(c)
}
func OkMsg(msg string, c *gin.Context) {
	Response{
		Code: SuccessCode,
		Data: empty,
		Msg:  msg,
	}.JSON(c)
}
func OkAny(data any, msg string, c *gin.Context) {
	Response{
		Code: SuccessCode,
		Data: data,
		Msg:  msg,
	}.JSON(c)
}
func Failed(data any, c *gin.Context) {
	Response{
		Code: FailServiceCode,
		Data: data,
		Msg:  FailServiceCode.String(),
	}.JSON(c)
}
func FailedAny(data any, msg string, c *gin.Context) {
	Response{
		Code: FailServiceCode,
		Data: data,
		Msg:  msg,
	}.JSON(c)
}
func FailedMsg(msg string, c *gin.Context) {
	Response{
		Code: FailValidCode,
		Data: empty,
		Msg:  msg,
	}.JSON(c)
}
func ExpireMsg(msg string, c *gin.Context) {
	Response{
		Code: TokenExpire,
		Data: empty,
		Msg:  msg,
	}.JSON(c)
}
