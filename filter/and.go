package filter

// AndFilter
type AndFilter struct {
	rightCondition, leftCondition Filter
}

func NewAndFilter(leftCondition, rightCondition Filter) Filter {
	f := AndFilter{leftCondition: leftCondition, rightCondition: rightCondition}

	return f
}

func (c AndFilter) GroupAnd(other Filter) Filter {
	return NewGroupAndFilter(c, other)
}

func (c AndFilter) And(other Filter) Filter {
	return NewAndFilter(c, other)
}

func (c AndFilter) Not() Filter {
	return NewNotFilter(c)
}

func (c AndFilter) Or(left Filter) Filter {
	return NewOrFilter(c, left)
}

func (a AndFilter) Execute(app FilterExecutable) (any, error) {
	return app.And(a.leftCondition, a.rightCondition)
}

func (c AndFilter) GroupOr(other Filter) Filter {
	return NewGroupOrFilter(c, other)
}
