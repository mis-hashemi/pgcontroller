package filter

type FilterExecutable interface {
	Execute(FilterQuery) (any, error)
	And(left, right Filter) (any, error)
	Or(left, right Filter) (any, error)
	Not(f Filter) (any, error)
	GroupAnd(left, right Filter) (any, error)
	GroupOr(left, right Filter) (any, error)
}
