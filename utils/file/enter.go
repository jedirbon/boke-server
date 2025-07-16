package file

import (
	"boke-server/utils"
	"errors"
	"strings"
)

var whiteList = []string{
	"jpg",
	"jpeg",
	"png",
	"webg",
	"gif",
}

func ImageSuffixJudge(filename string) (error, string) {
	_list := strings.Split(filename, ".")
	//xxx.jpg
	if len(_list) == 1 {
		return errors.New("错误的文件名"), ""
	}
	suffix := _list[len(_list)-1]
	if !utils.InList(suffix, whiteList) {
		return errors.New("文件类型错误!"), ""
	}
	return nil, _list[len(_list)-1]
}
