package conf

type Jwt struct {
	Key    string `yaml:"key"`
	Time   int    `yaml:"time"`
	Person string `yaml:"person"`
}
