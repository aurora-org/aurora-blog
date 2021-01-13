package entity

type PaginationData struct {
	Pagination Pagination  `json:"pagination"`
	Objects    interface{} `json:"objects"`
}

type Pagination struct {
	Total int `json:"total"`
	From  int `json:"from"`
	To    int `json:"to"`
}

func NewEmptyPaginationData() PaginationData {
	return PaginationData{
		Pagination: Pagination{
			Total: 0,
			From:  0,
			To:    0,
		},
		Objects: []struct{}{},
	}
}
