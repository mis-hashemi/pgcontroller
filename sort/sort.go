package sort

const (
	QuerySort = "sort"
)

type sortAdapter struct {
	name  string
	asc   bool
	order int
}

func NewSort(name string, asc bool, order int) Sort {
	return sortAdapter{name: name, asc: asc, order: order}
}

func (s sortAdapter) GetName() string {
	return s.name
}

func (s sortAdapter) IsAscending() bool {
	return s.asc
}

func (s sortAdapter) GetOrder() int {
	return s.order
}
