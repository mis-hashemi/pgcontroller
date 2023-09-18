package filter

import "github.com/mis-hashemi/request-parameter/query"

type FilterQuery interface {
	GetName() string
	SetName(string)
	GetOperator() query.QueryOperator
	GetOperand() *query.Operand
}

type basicFilterQuery struct {
	name    string
	op      query.QueryOperator
	operand *query.Operand
}

func NewFilterQuery(name string, op query.QueryOperator, operand *query.Operand) FilterQuery {

	return &basicFilterQuery{name: name, op: op, operand: operand}
}

func (b basicFilterQuery) GetName() string {
	return b.name
}

func (b *basicFilterQuery) SetName(name string) {
	b.name = name
}

func (b basicFilterQuery) GetOperator() query.QueryOperator {
	return b.op
}

func (b basicFilterQuery) GetOperand() *query.Operand {
	return b.operand
}
