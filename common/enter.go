package common

type PageInfo struct {
	PageIndex int `json:"pageIndex" from:"pageIndex"`
	PageSize  int `json:"pageSize" from:"pageSize"`
}

type Delete struct {
	Ids []int `json:"ids"`
}

type IDRequest struct {
	Id []int `json:"id" from:"id"`
}

type ResponseList[T any] struct {
	List       []T   `json:"list"`
	TotalCount int64 `json:"totalCount"`
}
