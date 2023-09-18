package page

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func ParseEchoQueryParamPagination(c echo.Context) (Pagination, bool) {
	size := c.QueryParam(QuerySize)
	if size == "" {
		return nil, false
	}
	sizeNum, err := strconv.Atoi(size)
	if err != nil {
		return nil, false
	}
	limit := sizeNum
	pageQuery := c.QueryParam(QueryPage)
	if pageQuery == "" {
		return NewPage(0, 10), false
	}
	pageNum, err := strconv.Atoi(pageQuery)
	if err != nil {
		return nil, false
	}
	skip := (pageNum - 1) * limit
	p := NewPage(skip, limit)
	p.SetPageNum(pageNum)
	return p, true

}
