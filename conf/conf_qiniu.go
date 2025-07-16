package conf

type QiNiu struct {
	Enable    bool   `yaml:"enable" json:"enable"`
	AccessKey string `yaml:"accessKey" json:"accessKey"`
	SecretKey string `yaml:"secretKey" json:"secretKey"`
	Bucket    string `yaml:"Bucket" json:"Bucket"`
	Uri       string `yaml:"uri" json:"uri"`
	Region    string `yaml:"region" json:"region"`
	Prefix    string `yaml:"prefix" json:"prefix"`
	Size      int64  `yaml:"size" json:"size"` //大小限制
	Expiry    int    `yaml:"expiry"`           //国企时间
}
