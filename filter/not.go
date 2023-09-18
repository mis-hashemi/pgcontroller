package filter

type NotFilter struct {
	wrapped Filter
}

func NewNotFilter(condition Filter) Filter {
	f := NotFilter{wrapped: condition}
	return f
}

func (c NotFilter) And(other Filter) Filter {
	return NewAndFilter(c, other)
}

func (c NotFilter) Not() Filter {
	return NewNotFilter(c.wrapped)
}

func (c NotFilter) Or(left Filter) Filter {
	return NewOrFilter(c, left)
}

func (c NotFilter) GroupAnd(other Filter) Filter {
	return NewGroupAndFilter(c, other)
}

func (a NotFilter) Execute(app FilterExecutable) (any, error) {
	return app.Not(a.wrapped)
}

func (c NotFilter) GroupOr(other Filter) Filter {
	return NewGroupOrFilter(c, other)
}
