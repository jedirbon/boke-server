package core

import (
	"boke-server/conf"
	"boke-server/flags"
	"boke-server/global"
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
)

var confPath = "settings.yaml"

func ReadConf() *conf.Config {
	byteData, err := os.ReadFile(flags.FlagOptions.File)
	if err != nil {
		panic(err)
	}
	var c = new(conf.Config)
	err = yaml.Unmarshal(byteData, c)
	if err != nil {
		panic(fmt.Sprintf("yaml配置文件格式错误%s", err))
	}

	fmt.Printf("读取配置文件 %s 成功\n", flags.FlagOptions.File)
	return c
}
func SetConf() {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		logrus.Error("conf读取失败\n", err)
	}
	err = os.WriteFile(flags.FlagOptions.File, byteData, 0666)
	if err != nil {
		logrus.Errorf("设置文件失败 %s", err)
		return
	}
}
