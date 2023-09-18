package sort

import (
	"strings"

	"github.com/labstack/echo/v4"
)

func ParseEchoQueryParamSort(c echo.Context) []Sort {
	out := []Sort{}
	sortQuery := c.QueryParam(QuerySort)

	if sortQuery == "" {
		return out
	}

	parts := strings.Split(sortQuery, ",")

	for i := 0; i < len(parts); i++ {
		part := parts[i]

		if part == "" {
			continue
		}
		part = strings.TrimSpace(part)
		rawArr := strings.Split(part, "-")

		if len(rawArr) == 2 {
			asc := false
			name := rawArr[1]
			order := i + 1
			if strings.ContainsAny(name, "-+") {
				continue
			}
			s := NewSort(name, asc, order)
			out = append(out, s)
			continue
		}

		rawArr = strings.Split(part, "+")

		if len(rawArr) == 2 || len(rawArr) == 1 {
			asc := true
			name := ""
			if len(rawArr) == 1 {
				name = rawArr[0]
			} else {
				name = rawArr[1]
			}
			order := i + 1
			if strings.ContainsAny(name, "-+") {
				continue
			}
			s := NewSort(name, asc, order)
			out = append(out, s)
		}

	}
	return out

}
