package common

type PageInfo struct {
	PageIndex int `json:"pageIndex" form:"pageIndex"`
	PageSize  int `json:"pageSize" form:"pageSize"`
}

type Delete struct {
	Ids []int `json:"ids"`
}

type IDRequest struct {
	Id []int `json:"id" form:"id"`
}

type ResponseList[T any] struct {
	List       []T   `json:"list"`
	TotalCount int64 `json:"totalCount"`
}
