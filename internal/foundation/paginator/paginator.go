package paginator

var pageSize = 8

type Paginator struct {
	Rows    int    `form:"rows"`
	Page    int    `form:"page"`
	Keyword string `form:"keyword"`
	Desc    bool   `form:"desc"`
	Counts  int
}

func NewPagintor(page, pagenums int) *Paginator {
	return &Paginator{Page: page, Rows: pageSize, Counts: pagenums}
}
