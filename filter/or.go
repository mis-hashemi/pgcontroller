package filter

type OrFilter struct {
	leftCondition, rightCondition Filter
}

func NewOrFilter(leftCondition, rightCondition Filter) Filter {
	f := OrFilter{leftCondition: leftCondition, rightCondition: rightCondition}
	return f
}

func (c OrFilter) And(other Filter) Filter {
	return NewAndFilter(c, other)
}

func (c OrFilter) Or(left Filter) Filter {
	return NewOrFilter(c, left)
}

func (c OrFilter) GroupAnd(other Filter) Filter {
	return NewGroupAndFilter(c, other)
}

func (c OrFilter) Not() Filter {
	return NewNotFilter(c)
}

func (a OrFilter) Execute(app FilterExecutable) (any, error) {
	return app.Or(a.leftCondition, a.rightCondition)
}

func (c OrFilter) GroupOr(other Filter) Filter {
	return NewGroupOrFilter(c, other)
}
