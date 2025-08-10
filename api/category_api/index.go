package category_api

import (
	"boke-server/common/res"
	"boke-server/global"
	"boke-server/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

type CateGory struct {
}

type PageResponse struct {
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}

// 读取类别列表
func (CateGory) GetGoryList(c *gin.Context) {
	var categoryList []models.CategoryModel
	var pg PageResponse
	pg.PageIndex = 1
	pg.PageSize = 10
	err := c.ShouldBind(&pg)
	if err != nil {
		res.FailedMsg(err.Error(), c)
		return
	}
	result := global.DB.Model(&models.CategoryModel{}).Limit(pg.PageSize).Offset((pg.PageIndex - 1) * pg.PageSize).Find(&categoryList)
	if result.Error != nil {
		res.FailedMsg(result.Error.Error(), c)
		return
	} else {
		res.Ok(categoryList, c)
		return
	}
}

// 添加类别
func (CateGory) AddCateGory(c *gin.Context) {
	var data models.CategoryModel
	err := c.ShouldBind(&data)
	if err != nil {
		res.FailedMsg(err.Error(), c)
		return
	}
	result := global.DB.Create(&data)
	if result.Error != nil {
		res.FailedMsg(result.Error.Error(), c)
		return
	} else {
		res.Ok(data, c)
		return
	}
}

// 删除类别
func (CateGory) DelCateGory(c *gin.Context) {
	var dels []int
	err := c.ShouldBindJSON(&dels)
	fmt.Println(dels)
	if err != nil {
		res.FailedMsg(err.Error(), c)
		return
	}
	global.DB.Model(&models.CategoryModel{})
	result := global.DB.Where("id in ?", dels).Delete(&models.CategoryModel{})
	if result.Error != nil {
		res.FailedMsg(result.Error.Error(), c)
		return
	}
	if result.RowsAffected > 0 {
		res.OkMsg(fmt.Sprintf("删除了%d条数据", result.RowsAffected), c)
	} else {
		res.FailedMsg("未找到该数据", c)
	}
}

// 上传文章封面通用接口
func (CateGory) UploadCover(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		res.FailedMsg(err.Error(), c)
		return
	}
	dst := fmt.Sprintf("uploads/cover/%s", file.Filename)
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		res.FailedMsg(err.Error(), c)
		return
	}
	res.Ok(map[string]string{"url": dst}, c)
}
