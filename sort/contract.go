package sort

type Sort interface {
	GetName() string
	IsAscending() bool
	GetOrder() int
}
