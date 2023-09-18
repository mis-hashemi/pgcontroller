package filter

// AndFilter
type GrouporFilter struct {
	rightCondition, leftCondition Filter
}

func NewGroupOrFilter(leftCondition, rightCondition Filter) Filter {
	f := GrouporFilter{leftCondition: leftCondition, rightCondition: rightCondition}
	return f
}

func (c GrouporFilter) And(other Filter) Filter {
	return NewAndFilter(c, other)
}

func (c GrouporFilter) GroupAnd(other Filter) Filter {
	return NewGroupAndFilter(c, other)
}

func (c GrouporFilter) GroupOr(other Filter) Filter {
	return NewGroupOrFilter(c, other)
}

func (c GrouporFilter) Not() Filter {
	return NewNotFilter(c)
}

func (c GrouporFilter) Or(left Filter) Filter {
	return NewOrFilter(c, left)
}

func (a GrouporFilter) Execute(app FilterExecutable) (any, error) {
	return app.GroupOr(a.leftCondition, a.rightCondition)
}
