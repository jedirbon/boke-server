package models

type BannerModel struct {
	Show  bool   `json:"show"`  //是否可展示
	Cover string `json:"cover"` //图片链接
	Href  string `json:"href"`  //跳转链接
	Model
}
