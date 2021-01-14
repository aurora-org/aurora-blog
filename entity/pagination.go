package entity

type PaginationData struct {
	Pagination Pagination  `json:"pagination"`
	Objects    interface{} `json:"objects"`
}

type Pagination struct {
	Total int `json:"total"`
	Page  int `json:"page"`
	Size  int `json:"size"`
}

func NewEmptyPaginationData() PaginationData {
	return PaginationData{
		Pagination: Pagination{
			Total: 0,
			Page:  0,
			Size:  0,
		},
		Objects: []struct{}{},
	}
}
