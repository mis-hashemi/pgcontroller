package filter

import "github.com/mis-hashemi/request-parameter/query"

type FilterExecuter interface {
	Execute(app FilterExecutable) (any, error)
}

type Filter interface {
	FilterExecuter
	And(Filter) Filter
	GroupAnd(Filter) Filter
	GroupOr(Filter) Filter
	Or(Filter) Filter
	Not() Filter
}

type CompositeFilter struct {
	query FilterQuery
}

func NewFilter(query FilterQuery) Filter {
	return CompositeFilter{query: query}
}

func NewQueryFilter(name string, op query.QueryOperator, operand *query.Operand) Filter {
	return CompositeFilter{query: NewFilterQuery(name, op, operand)}
}

func NewEmptyFilter(name string, op query.QueryOperator) Filter {
	return CompositeFilter{query: NewFilterQuery(name, op, query.NewOperand(nil))}
}

func (c CompositeFilter) Execute(applier FilterExecutable) (any, error) {
	return applier.Execute(c.query)
}

func (c CompositeFilter) And(other Filter) Filter {
	return NewAndFilter(c, other)
}

func (c CompositeFilter) Or(left Filter) Filter {
	return NewOrFilter(c, left)
}

func (c CompositeFilter) GroupAnd(other Filter) Filter {
	return NewGroupAndFilter(c, other)
}

func (c CompositeFilter) Not() Filter {
	return NewNotFilter(c)
}

func (c CompositeFilter) GroupOr(other Filter) Filter {
	return NewGroupOrFilter(c, other)
}
