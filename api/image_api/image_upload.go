package image_api

import (
	"boke-server/common/res"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"strings"
)

func (ImageApi) ImageUploadView(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		res.Failed(err, c)
		return
	}

	//文件大小判断
	if (fileHeader.Size / 1024 / 1024) > 5 {
		res.FailedMsg("图片不得超过5MB", c)
		return
	}
	//后缀判断
	filename := fileHeader.Filename
	err, suffixName := imageSuffixJudge(filename)
	if err != nil {
		res.FailedMsg("文件类型错误！", c)
		return
	}
	//文件hash
	file, err := fileHeader.Open()
	if err != nil {
		res.FailedMsg(err.Error(), c)
		return
	}
	byteData, _ := io.ReadAll(file)
	hash := Hash(byteData)
	filePath := fmt.Sprintf("uploads/images/%s", hash+"."+suffixName)
	err = c.SaveUploadedFile(fileHeader, filePath)
	if err == nil {
		res.OkAny("/"+filePath, "上传成功", c)
	}
}

var whiteList = []string{
	"jpg",
	"jpeg",
	"png",
	"webg",
	"gif",
}

func InList(key string, list []string) bool {
	for _, s := range list {
		if key == s {
			return true
		}
	}
	return false
}

func imageSuffixJudge(filename string) (error, string) {
	_list := strings.Split(filename, ".")
	//xxx.jpg
	if len(_list) == 1 {
		return errors.New("错误的文件名"), ""
	}
	suffix := _list[len(_list)-1]
	if !InList(suffix, whiteList) {
		return errors.New("文件类型错误!"), ""
	}
	return nil, _list[len(_list)-1]
}
func Hash(data []byte) string {
	md5New := md5.New()
	md5New.Write(data)
	//hex转字符串
	return hex.EncodeToString(md5New.Sum(nil))
}

type RequestData struct {
	Name string `json:"name" form:"name"`
}

func (ImageApi) OnUploadImage(c *gin.Context) {
	var rq RequestData
	err := c.ShouldBind(&rq)
	fmt.Println(rq.Name)
	if err != nil {
		res.FailedMsg("参数错误", c)
		return
	}
	filePath := fmt.Sprintf("uploads/images/%s", rq.Name)
	fmt.Println(filePath)

	err = os.Remove(filePath)
	if err != nil {
		if !fileExists(filePath) {
			res.FailedAny(err, "删除地址不存在", c)
		} else {
			res.FailedAny(err, "删除失败", c)
		}
		return
	}
	res.OkMsg("删除成功", c)
	return
}
func fileExists(filepath string) bool {
	_, err := os.Open(filepath)
	return err == nil
}
