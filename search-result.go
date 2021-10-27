package gssdk

type Pagination struct {
	Total uint64 `json:"total"`
	Pages uint64 `json:"pages"`
	PageSize int `json:"page-size"`
	CurrPageNo uint64 `json:"curr-page"`
	DocsInPage int `json:"page-count"`
}

type Doc map[string]interface{}
