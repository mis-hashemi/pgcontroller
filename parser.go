package pgcontroller

import (
	"fmt"
	"strings"
	"time"

	"github.com/mis-hashemi/pgcontroller/filter"
	"github.com/mis-hashemi/request-parameter/query"
)

type statelessParser struct{}

func NewParser() statelessParser {
	return statelessParser{}
}

func (statelessParser) ParseFilter(f filter.Filter) (string, error) {
	if f == nil {
		return "", nil
	}
	applier := &applyer{}
	sqlQuery, err := f.Execute(applier)
	if err != nil {
		return "", err
	}
	return sqlQuery.(string), nil
}

func getJsonOperatorInfo(op query.QueryOperator) (string, error) {

	switch op {
	case query.QueryOperatorEqual:
		return string(SqlOperatorEqual), nil
	case query.QueryOperatorNotEqual:
		return "<>", nil
	case query.QueryOperatorContain:
		return string(SqlOperatorContain), nil
	case query.QueryOperatorNotContain:
		return string(SqlOperatorNotContain), nil
	case query.QueryOperatorEmpty:
		return SqlIsNull, nil
	case query.QueryOperatorNotEmpty:
		return SqlIsNotNull, nil
	case query.QueryOperatorMoreThan:
		return string(SqlMoreThan), nil
	case query.QueryOperatorEqualOrMoreThan:
		return string(SqlEqualOrMoreThan), nil
	case query.QueryOperatorLessThan:
		return string(SqlLessThan), nil
	case query.QueryOperatorEqualOrLessThan:
		return string(SqlLessOrEqualThan), nil
	case query.QueryOperatorIn:
		return string(SqlOperatorIn), nil
	case query.QueryOperatorNotIn:
		return string(SqlOperatorIn), nil
	}

	return "", fmt.Errorf("Operator %s not supported", op)
}

type applyer struct {
	Vars      []any
	LastQuery filter.FilterQuery
}

func (a *applyer) Execute(filterQuery filter.FilterQuery) (any, error) {
	operand := filterQuery.GetOperand()
	op, err := getJsonOperatorInfo(filterQuery.GetOperator())
	if err != nil {
		return nil, err
	}
	if subFilter, ok := operand.Value.(filter.Filter); ok {
		subQuery, err := subFilter.Execute(a)
		if err != nil {
			return nil, err
		}
		return a.buildCondition(filterQuery.GetName(), op, subQuery), nil
	}
	if operand != nil && operand.Value != nil {
		return a.buildCondition(filterQuery.GetName(), op, operand.Value), nil
	}

	switch op {
	case SqlIsNull:
		return fmt.Sprintf("( %s IS NULL)", filterQuery.GetName()), nil
	case SqlIsNotNull:
		return fmt.Sprintf("(%s IS NOT NULL)", filterQuery.GetName()), nil
	default:
		return nil, fmt.Errorf("operator %s not supported for nil operand", op)
	}
}

func (a *applyer) buildCondition(fieldName, operator string, value interface{}) string {
	if operator == string(SqlOperatorIn) || operator == string(SqlOperatorNotIn) {
		valuesAny := value.([]any)
		values := make([]string, len(valuesAny))
		for i, v := range valuesAny {
			values[i] = fmt.Sprintf("'%s'", v)
		}
		return fmt.Sprintf("%s %s (%s)", fieldName, operator, strings.Join(values, ","))
	}
	if operator == string(SqlOperatorContain) {
		return fmt.Sprintf("%s %s '%%%s%%'", fieldName, operator, value)
	}
	var castType string
	switch v := value.(type) {
	case string:
		castType = "text"
	case int, int32, int64, float64:
		castType = "int"
	case bool:
		castType = "boolean"
	case time.Time:
		castType = "timestamp"
	// case float64:
	// 	castType = "double precision"
	default:
		fmt.Printf("value type %T!\n", v)
	}
	if castType != "" {
		return fmt.Sprintf("(%s)::%s %s %v", fieldName, castType, operator, value)
	}

	return fmt.Sprintf("%s %s %v", fieldName, operator, value)
}

func (a *applyer) And(left, right filter.Filter) (any, error) {
	leftEx, err := left.Execute(a)
	if err != nil {
		return nil, err
	}
	rightEx, err := right.Execute(a)
	if err != nil {
		return nil, err
	}
	return " " + leftEx.(string) + " AND " + rightEx.(string), nil
}

func (a *applyer) Not(wrapped filter.Filter) (any, error) {
	wrappedEx, err := wrapped.Execute(a)
	if err != nil {
		return nil, err
	}
	return " NOT (" + wrappedEx.(string) + ")", nil
}

func (a *applyer) Or(left, right filter.Filter) (any, error) {
	leftEx, err := left.Execute(a)
	if err != nil {
		return nil, err
	}
	rightEx, err := right.Execute(a)
	if err != nil {
		return nil, err
	}
	return " " + leftEx.(string) + " OR " + rightEx.(string), nil
}

func (a *applyer) GroupAnd(left, right filter.Filter) (any, error) {
	leftEx, err := left.Execute(a)
	if err != nil {
		return nil, err
	}
	rightEx, err := right.Execute(a)
	if err != nil {
		return nil, err
	}
	return " " + leftEx.(string) + " AND (" + rightEx.(string) + ")", nil
}

func (a *applyer) GroupOr(left, right filter.Filter) (any, error) {
	leftEx, err := left.Execute(a)
	if err != nil {
		return nil, err
	}
	rightEx, err := right.Execute(a)
	if err != nil {
		return nil, err
	}
	return " " + leftEx.(string) + " OR (" + rightEx.(string) + ")", nil
}
