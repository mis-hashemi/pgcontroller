package userservice

import (
	"context"

	"github.com/mis-hashemi/pgcontroller/example/example01/entity"
	"github.com/mis-hashemi/pgcontroller/page"
	"github.com/mis-hashemi/pgcontroller/sort"
	"github.com/mis-hashemi/request-parameter/query"
)

func (s Service) GetAll(ctx context.Context, filter query.QueryInfo, sorts []sort.Sort, page page.Pagination) ([]entity.User, error) {
	return s.repo.GetAll(ctx, filter, sorts, page)
}
