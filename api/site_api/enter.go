package site_api

import (
	"boke-server/common/res"
	"boke-server/conf"
	"boke-server/core"
	"boke-server/global"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
)

type SiteApi struct {
}
type SiteInfoRequest struct {
	Name string `uri:"name"`
}

func (SiteApi) SiteInfoView(c *gin.Context) {
	var cr SiteInfoRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.Failed(err, c)
		return
	}
	if cr.Name == "site" {
		global.Config.Site.About.Version = global.Version
		res.Ok(global.Config.Site, c)
		return
	}
	//
	var data any
	switch cr.Name {
	case "site":
		data = global.Config.Site
	case "email":
		rep := global.Config.Email
		rep.AuthCode = "*******"
		data = rep
	case "qq":
		rep := global.Config.QQ
		rep.AppKey = "*******"
		data = rep
	case "qiNiu":
		rep := global.Config.QiNiu
		rep.SecretKey = "*******"
		data = rep
	case "ai":
		rep := global.Config.Ai
		rep.SecretKey = "*******"
		data = rep
	default:
		res.FailedMsg("不存在的配置", c)
		return
	}
	res.Ok(data, c)
	return
}

func (SiteApi) SiteUpdateView(c *gin.Context) {
	var cr SiteInfoRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailedMsg("参数错误", c)
		return
	}
	var rep any
	switch cr.Name {
	case "site":
		var data conf.Site
		err = c.ShouldBindJSON(&data)
		rep = data
	case "email":
		var data conf.Email
		err = c.ShouldBindJSON(&data)
		rep = data
	case "qq":
		var data conf.QQ
		err = c.ShouldBindJSON(&data)
		rep = data
	case "qiNiu":
		var data conf.QiNiu
		err = c.ShouldBindJSON(&data)
		rep = data
	case "ai":
		var data conf.Ai
		err = c.ShouldBindJSON(&data)
		rep = data
	default:
		res.FailedMsg("不存在的配置", c)
		return
	}
	if err != nil {
		res.Failed("参数错误", c)
		return
	}
	switch s := rep.(type) {
	case conf.Site:
		//判断站点信息更新前端文件部分
		err := UpdateSite(s)
		if err != nil {
			res.Failed(err, c)
			return
		}
		global.Config.Site = s
	case conf.Email:
		if s.AuthCode == "******" {
			s.AuthCode = global.Config.Email.AuthCode
		}
		global.Config.Email = s
	case conf.QQ:
		if s.AppKey == "******" {
			s.AppKey = global.Config.QQ.AppKey
		}
		global.Config.QQ = s
	case conf.QiNiu:
		if s.SecretKey == "******" {
			s.SecretKey = global.Config.QiNiu.SecretKey
		}
		global.Config.QiNiu = s
	case conf.Ai:
		if s.SecretKey == "******" {
			s.SecretKey = global.Config.Ai.SecretKey
		}
		global.Config.Ai = s

	}
	//改配置文件
	core.SetConf()

	res.OkMsg("更新站点配置成功！", c)
	return
}

func (SiteApi) SiteInfoQQView(c *gin.Context) {
	res.Ok(global.Config.QQ.Url(), c)
}

func UpdateSite(site conf.Site) error {
	if site.Project.Icon == "" && site.Project.Title == "" ||
		site.Seo.Keywords == "" && site.Seo.Description == "" &&
			site.Project.WebPath == "" {
		return nil
	}
	if site.Project.WebPath == "" {
		return errors.New("请配置前端地址")
	}

	file, err := os.Open(site.Project.WebPath)
	if err != nil {
		return errors.New(fmt.Sprintf("%s文件不存在", site.Project.WebPath))
	}
	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		logrus.Errorf("goquery 解析失败 %s", err)
		return errors.New("文件解析失败")
	}
	if site.Project.Title != "" {
		doc.Find("title").SetText(site.Project.Title)
	}
	if site.Project.Icon != "" {
		if doc.Is("link[ref='icon']") {
			//有就修改
			doc.Find("link[ref='icon']").SetAttr("href", site.Project.Icon)
		} else {
			//没有就创建
			doc.Find("head").AppendHtml(fmt.Sprintf("<link rel=\"icon\" href=\"/logo.png\">"))
		}
	}
	html, err := doc.Html()
	if err != nil {
		logrus.Errorf("生成html失败")
		return errors.New("生成html失败")
	}
	err = os.WriteFile(site.Project.WebPath, []byte(html), 0666)
	if err != nil {
		logrus.Errorf("文件写入失败 %s", err)
		return errors.New("文件写入失败")
	}
	return nil
}
