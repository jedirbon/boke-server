package conf

import "fmt"

type DB struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Debug    bool   `yaml:"debug"`  // 打印全部日志
	Source   string `yaml:"source"` //数据库的源 mysql
}
type SlaveDB struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Debug    bool   `yaml:"debug"`  // 打印全部日志
	Source   string `yaml:"source"` //数据库的源 mysql
}

func (d DB) DSN() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.Dbname)
	return dsn
}

func (d SlaveDB) SlaveDSN() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.Dbname)
	return dsn
}
