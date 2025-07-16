package core

import (
	"boke-server/utils"
	"errors"
	"fmt"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"github.com/sirupsen/logrus"
	"strings"
)

var searcher *xdb.Searcher

func InitIPDB() {
	var dbPath = "init/ip2region.xdb"
	_searcher, err := xdb.NewWithFileOnly(dbPath)
	if err != nil {
		logrus.Fatalf("ip地址数据库加载失败")
		return
	}
	searcher = _searcher
}

func GetIpAddr(ip string) (arr string) {
	if utils.IsPrivateIP(ip) {
		fmt.Println("内网")
		return "内网"
	}
	region, err := searcher.SearchByStr(ip)
	if err != nil {
		logrus.Errorf("ip解析失败")
		return
	}
	_addrList := strings.Split(region, "|")
	if len(_addrList) != 5 {
		err = errors.New("")
		return
	}
	return strings.Join(_addrList, "-")
}
