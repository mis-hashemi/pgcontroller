package userservice

import (
	"context"

	"github.com/mis-hashemi/pgcontroller/example/example01/entity"
	"github.com/mis-hashemi/pgcontroller/page"
	"github.com/mis-hashemi/pgcontroller/sort"
	"github.com/mis-hashemi/request-parameter/query"
)

type userRepository interface {
	GetAll(ctx context.Context, filter query.QueryInfo, sorts []sort.Sort, page page.Pagination) ([]entity.User, error)
}

type Service struct {
	repo userRepository
}

func New(repo userRepository) Service {
	return Service{
		repo: repo,
	}
}
