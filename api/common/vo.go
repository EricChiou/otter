package common

// EmptyVo empty struct
type EmptyVo struct{}

// PageRespVo page response vo
type PageRespVo struct {
	Records []interface{} `json:"records"`
	Page    int         `json:"page"`
	Limit   int         `json:"limit"`
	Total   int         `json:"total"`
}

// OrderByID order by id struct
type OrderByID struct {
	ID int `json:"id" column:"id"`
}
