package site

type SiteInfo struct {
	Title string `yaml:"title" json:"title"`
	Logo  string `yaml:"logo" json:"logo"`
	Beian string `yaml:"beian" json:"beian"`
	Mode  int8   `yaml:"mode" json:"mode"` //1社区模式，博客模式
}

type Project struct {
	Title   string `yaml:"title" json:"title"`
	Icon    string `yaml:"icon" json:"icon"`
	WebPath string `yaml:"webPath" json:"webPath"`
}

type Seo struct {
	Keywords    string `yaml:"keywords" json:"keywords"`
	Description string `yaml:"description" json:"description"`
}

type About struct {
	SiteDate string `yaml:"siteDate" json:"siteDate"` //年月日
	QQ       string `yaml:"qq" json:"qq"`
	Version  string `yaml:"version"`
	Wechat   string `yaml:"wechat" json:"wechat"`
	Gitee    string `yaml:"gitee" json:"gitee"`
	Bilibili string `yaml:"bilibili" json:"bilibili"`
	Github   string `yaml:"github" json:"github"`
}

type Login struct {
	QQLogin          bool `yaml:"qqLogin" json:"qqLogin"`
	UsernamePwdLogin bool `yaml:"usernamePwdLogin" json:"usernamePwdLogin"`
	EmailLogin       bool `yaml:"emailLogin" json:"emailLogin"`
	Captcha          bool `yaml:"captcha" json:"captcha"` //验证码
}
type ComponentInfo struct {
	Title  string `yaml:"title" json:"title"`
	Enable bool   `yaml:"enable" json:"enable"`
}
type IndexRight struct {
	List []ComponentInfo `json:"list" yaml:"list"`
}
type Article struct {
	NoExamine bool `yaml:"noExamine" json:"noExamine"` //免审核
}
