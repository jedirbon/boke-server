package requestModel

import "boke-server/global"

type RequestParams struct {
	Title  string `json:"title" form:"title"`
	Status string `json:"status" form:"status"`
	global.PageResponse
}
