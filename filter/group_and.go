package filter

// AndFilter
type GroupAndFilter struct {
	rightCondition, leftCondition Filter
}

func NewGroupAndFilter(leftCondition, rightCondition Filter) Filter {
	f := GroupAndFilter{leftCondition: leftCondition, rightCondition: rightCondition}
	return f
}

func (c GroupAndFilter) And(other Filter) Filter {
	return NewAndFilter(c, other)
}

func (c GroupAndFilter) GroupAnd(other Filter) Filter {
	return NewGroupAndFilter(c, other)
}

func (c GroupAndFilter) Not() Filter {
	return NewNotFilter(c)
}

func (c GroupAndFilter) Or(left Filter) Filter {
	return NewOrFilter(c, left)
}

func (a GroupAndFilter) Execute(app FilterExecutable) (any, error) {
	return app.GroupAnd(a.leftCondition, a.rightCondition)
}

func (c GroupAndFilter) GroupOr(other Filter) Filter {
	return NewGroupOrFilter(c, other)
}
