package global

// 分页查询标准结构体
type PageResponse struct {
	PageIndex int `json:"pageIndex" form:"pageIndex"`
	PageSize  int `json:"pageSize" form:"pageSize"`
}

// 列表返回标准结构体
type ResponseData[T any] struct {
	Data       []T   `json:"data"`
	TotalCount int64 `json:"totalCount"`
}
