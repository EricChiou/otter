package codemap

// AddReqVo add codemap vo
type AddReqVo struct {
	Type   string `json:"type"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	SortNo int    `json:"sortNo"`
	Enable bool   `json:"enable"`
}

// UpdateReqVo update codemap vo
type UpdateReqVo struct {
	ID     int    `json:"id"`
	Type   string `json:"type"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	SortNo int    `json:"sortNo"`
	Enable bool   `json:"enable"`
}

// DeleteReqVo delete codemap vo
type DeleteReqVo struct {
	ID int `json:"id"`
}

// ListResVo get codemap list response vo
type ListResVo struct {
	Records []ListDataVo `json:"records"`
	Total   int          `json:"total"`
}

// ListDataVo user list data vo
type ListDataVo struct {
	ID     int    `json:"id"`
	Type   string `json:"type"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	SortNo int    `json:"sortNo"`
	Enable bool   `json:"enable"`
}
