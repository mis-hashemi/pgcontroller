package page

type Pagination interface {
	GetPageNum() uint
	SetPageNum(page int)
	GetSkip() uint
	GetLimit() uint
}

type CursorPagination interface {
	GetLimit() uint
	GetCursor() string
	After() bool
}
