package page

const (
	QuerySize = "page_size"
	QueryPage = "page"
)

type page struct {
	page  int `json:"page"`
	skip  int `json:"skip"`
	limit int `json:"limit"`
}

func NewPage(skip, limit int) Pagination {
	return &page{skip: skip, limit: limit}
}

func (p page) GetSkip() uint {

	if p.skip < 0 {
		return 0
	}
	return uint(p.skip)
}

func (p page) GetPageNum() uint {

	if p.page < 0 {
		return 0
	}
	return uint(p.page)
}

func (p *page) SetPageNum(page int) {
	p.page = page
}

func (p page) GetLimit() uint {
	if p.limit <= 0 {
		return 10
	}
	return uint(p.limit)
}

type cursorPagination struct {
	Limit   int    `json:"limit"`
	Cursor  string `json:"cursor"`
	IsAfter bool   `json:"after""`
}

func NewCursorPagination(limit int, cursor string, isAfter bool) cursorPagination {
	return cursorPagination{Limit: limit, Cursor: cursor, IsAfter: isAfter}
}

func (p cursorPagination) GetLimit() uint {
	if p.Limit <= 0 {
		return 10
	}
	return uint(p.Limit)
}

func (c cursorPagination) GetCursor() string {
	return c.Cursor
}

func (c cursorPagination) After() bool {
	return c.IsAfter
}
