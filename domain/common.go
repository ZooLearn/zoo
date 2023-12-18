package domain

type Pagination struct {
	Page     uint64 `json:"page"`
	PageSize uint64 `json:"page_size"`
	Total    uint64 `json:"total"`
}

type Response struct {
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
}
