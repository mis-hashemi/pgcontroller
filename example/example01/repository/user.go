package repository

import (
	"context"
	"database/sql"

	"github.com/mis-hashemi/pgcontroller"
	"github.com/mis-hashemi/pgcontroller/example/example01/entity"
	"github.com/mis-hashemi/pgcontroller/page"
	"github.com/mis-hashemi/pgcontroller/sort"
	"github.com/mis-hashemi/request-parameter/query"
)

type userRepository struct {
	conn *sql.DB
}

func NewUserRepository(getter pgcontroller.DataContextGetter) *userRepository {
	return &userRepository{
		conn: getter.GetDataContext(),
	}
}
func (repo *userRepository) GetAll(ctx context.Context, filter query.QueryInfo, sorts []sort.Sort, page page.Pagination) ([]entity.User, error) {
	query := `
	SELECT  id, first_name, last_name, phone_number	FROM users  `
	var values []interface{}
	placeholderCounter := 0
	if filter != nil && len(filter.GetQuery()) != 0 {
		whereClause, valuesFilter := pgcontroller.ConvertSqlQuery(filter, "", &placeholderCounter)
		query = query + " where " + whereClause
		values = append(values, valuesFilter...)
	}
	query = pgcontroller.AddSortsQuery(query, sorts)
	query = pgcontroller.AddPaginationQuery(query, page)
	rows, err := repo.conn.QueryContext(ctx, query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.PhoneNumber)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}
