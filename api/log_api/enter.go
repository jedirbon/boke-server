package log_api

import (
	"boke-server/common/res"
	"boke-server/global"
	"boke-server/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

type LogApi struct {
}

type QueryInfo struct {
	PageIndex int    `form:"pageIndex" json:"pageIndex"`
	PageSize  int    `form:"pageSize" json:"pageSize"`
	Key       string `form:"key" json:"key"`
	models.LogModel
}
type Response struct {
	models.LogModel
	UserName   string `json:"userName"`
	UserAvatar string `json:"userAvatar"`
}

func (LogApi) LogList(c *gin.Context) {
	//分页 查询
	var query_info QueryInfo
	var totalCount int64
	var logList []models.LogModel
	err := c.ShouldBindQuery(&query_info)
	if err != nil {
		res.Failed(err.Error(), c)
		return
	}
	var model = models.LogModel{
		LogType: query_info.LogType,
		Level:   query_info.Level,
		Ip:      query_info.Ip,
		Addr:    query_info.Addr,
	}
	var like = global.DB.Where("title like ?", fmt.Sprintf("%%%s%%", query_info.Key))
	result := global.DB.
		Preload("UserModel").
		Where(model).
		Limit(query_info.PageSize).
		Where(like).
		Offset((query_info.PageIndex - 1) * query_info.PageSize).
		Find(&logList)
	if result.Error != nil {
		res.Failed(result.Error, c)
	} else {
		var resList = make([]Response, 0)
		for _, logModel := range logList {
			resList = append(resList, Response{
				LogModel:   logModel,
				UserName:   logModel.UserModel.Username,
				UserAvatar: logModel.UserModel.Avatar,
			})
		}

		global.DB.Model(models.LogModel{}).Where(like).Where(model).Count(&totalCount)
		res.Ok(struct {
			TotalCount int64 `json:"totalCount"`
			Data       any   `json:"data"`
		}{
			TotalCount: totalCount,
			Data:       resList,
		}, c)
	}
}

type DelIds struct {
	IdList []uint `json:"idList"`
}

// 删除
func (LogApi) DeleteLog(c *gin.Context) {
	var ids DelIds
	var logList []models.LogModel
	err := c.ShouldBind(&ids)
	if err != nil {
		res.Failed(err, c)
		return
	}
	global.DB.Find(&logList).Where("id in ?", ids)
	if len(logList) > 0 {
		reslut := global.DB.Where("id in ?", ids).Delete(&models.LogModel{})
		if reslut.RowsAffected > 0 {
			res.OkMsg(fmt.Sprintf("删除了%d个数据", reslut.RowsAffected), c)
			return
		}
	}
	res.OkMsg("删除失败", c)
}
