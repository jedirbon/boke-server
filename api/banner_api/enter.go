package banner_api

import (
	"boke-server/common"
	"boke-server/common/res"
	"boke-server/global"
	"boke-server/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

type BannerApi struct {
}

type BannerCreateRequest struct {
	Cover string `json:"cover" binding:"required"`
	Href  string `json:"href"`
	Show  bool   `json:"show"`
}

func (BannerApi) BannerCreateView(c *gin.Context) {
	var cr BannerCreateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailedMsg(err.Error(), c)
		return
	}
	result := global.DB.Create(&models.BannerModel{
		Cover: cr.Cover,
		Href:  cr.Href,
		Show:  cr.Show,
	})
	if result.Error != nil {
		res.FailedMsg("创建失败", c)
		return
	}
	res.Ok(cr, c)
}

func (BannerApi) BannerList(c *gin.Context) {
	var bannerList []models.BannerModel
	var totalCount int64
	var pageInfo common.PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		res.FailedMsg(err.Error(), c)
		return
	}
	global.DB.Find(&models.BannerModel{}).Count(&totalCount)
	result := global.DB.
		Find(&bannerList).
		Limit(pageInfo.PageIndex).
		Offset((pageInfo.PageIndex - 1) * pageInfo.PageSize)
	if result.Error != nil {
		res.FailedMsg("创建失败", c)
		return
	}
	res.Ok(common.ResponseList[models.BannerModel]{
		List:       bannerList,
		TotalCount: totalCount,
	}, c)
}

func (BannerApi) BannerRemoveView(c *gin.Context) {
	var del common.Delete
	err := c.ShouldBind(&del)
	if err != nil {
		res.FailedMsg("参数错误", c)
		return
	}
	fmt.Println(del.Ids)
	var list []models.BannerModel
	global.DB.Find(&list, "id in ?", del.Ids)
	if len(list) > 0 {
		global.DB.Delete(&list)
	}
	res.OkMsg(fmt.Sprintf("删除了%d条数据,成功删除了%d条", len(del.Ids), len(list)), c)
}

func (BannerApi) BannerUpdateView(c *gin.Context) {
	var banner models.BannerModel
	err := c.ShouldBind(&banner)
	if err != nil {
		res.FailedMsg("参数错误", c)
		return
	}
	result := global.DB.Find(&models.BannerModel{}, "id = ?", banner.ID)
	if result.Error == nil && result.RowsAffected > 0 {
		err := global.DB.Save(&banner).Error
		if err != nil {
			res.FailedMsg(err.Error(), c)
			return
		}
		res.Ok(banner, c)
	}
}
