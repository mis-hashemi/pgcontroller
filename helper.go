package pgcontroller

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mis-hashemi/pgcontroller/page"
	"github.com/mis-hashemi/pgcontroller/sort"
	"github.com/mis-hashemi/request-parameter/query"
)

func AddSortsQuery(query string, sorts []sort.Sort) string {
	if len(sorts) == 0 {
		return query
	}

	sortQuery := " ORDER BY"
	for i, sort := range sorts {
		if i > 0 {
			sortQuery += ","
		}
		sortQuery += " " + sort.GetName()
		if !sort.IsAscending() {
			sortQuery += " DESC"
		}
	}

	return query + sortQuery
}

func AddPaginationQuery(query string, page page.Pagination) string {
	if page == nil {
		return query
	}

	limit := page.GetLimit()
	skip := page.GetSkip()

	if limit > 0 {
		query += " LIMIT " + strconv.FormatUint(uint64(limit), 10)
	}

	if skip > 0 {
		query += " OFFSET " + strconv.FormatUint(uint64(skip), 10)
	}

	return query
}

const (
	SqlIsNull    = "IS NULL"
	SqlIsNotNull = "IS NOT NULL"
)

type SqlOperator string

const (
	SqlOperatorEqual      SqlOperator = "="
	SqlOperatorNotEqual   SqlOperator = "!="
	SqlEqualOrMoreThan    SqlOperator = ">="
	SqlMoreThan           SqlOperator = ">"
	SqlLessThan           SqlOperator = "<"
	SqlLessOrEqualThan    SqlOperator = "<="
	SqlOperatorContain    SqlOperator = "LIKE"
	SqlOperatorNotContain SqlOperator = "NOT LIKE"
	SqlOperatorIn         SqlOperator = "IN"
	SqlOperatorNotIn      SqlOperator = "NOT IN"
	SqlOperatorEmpty      SqlOperator = "LIKE ''"
	SqlOperatorNotEmpty   SqlOperator = "NOT LIKE ''"
)
const (
	SqlParameter         = "(?)"
	SqlParameterMultiple = "(?)"
)

type ParameterExpectation int

const (
	ParameterExpectationZero ParameterExpectation = iota
	ParameterExpectationSingle
	ParameterExpectationMultiple
)

func getOperatorInfo(op query.QueryOperator) (SqlOperator, ParameterExpectation) {

	switch op {
	case query.QueryOperatorEqual:
		return SqlOperatorEqual, ParameterExpectationSingle
	case query.QueryOperatorNotEqual:
		return SqlOperatorNotEqual, ParameterExpectationSingle
	case query.QueryOperatorMoreThan:
		return SqlMoreThan, ParameterExpectationSingle
	case query.QueryOperatorEqualOrMoreThan:
		return SqlEqualOrMoreThan, ParameterExpectationSingle
	case query.QueryOperatorLessThan:
		return SqlLessThan, ParameterExpectationSingle
	case query.QueryOperatorEqualOrLessThan:
		return SqlLessOrEqualThan, ParameterExpectationSingle
	case query.QueryOperatorContain:
		return SqlOperatorContain, ParameterExpectationSingle
	case query.QueryOperatorNotContain:
		return SqlOperatorNotContain, ParameterExpectationSingle
	case query.QueryOperatorEmpty:
		return SqlOperatorEmpty, ParameterExpectationZero
	case query.QueryOperatorNotEmpty:
		return SqlOperatorNotEmpty, ParameterExpectationZero
	case query.QueryOperatorIn:
		return SqlOperatorIn, ParameterExpectationMultiple
	case query.QueryOperatorNotIn:
		return SqlOperatorNotIn, ParameterExpectationMultiple
	}

	panic("unknown Operator")
}
func ConvertSqlQuery(queryInfo query.QueryInfo, tblAlias string, placeholderCounter *int) (string, []interface{}) {
	whereClause, values := getWhereClause(queryInfo, tblAlias, placeholderCounter)

	// Append the WHERE clause to the original query
	// sqlQuery := fmt.Sprintf("%s WHERE %s", query, whereClause)

	return whereClause, values
}

func getWhereClause(queryInfo query.QueryInfo, tblAlias string, placeholderCounter *int) (string, []interface{}) {
	logicOperator := "AND"
	if !queryInfo.IsAnd() {
		logicOperator = "OR"
	}

	queries := queryInfo.GetQuery()
	clauses := make([]string, len(queries))
	var clauseValues []interface{}

	for i, q := range queries {
		clause, values := buildClause(q, tblAlias, placeholderCounter)
		clauses[i] = clause
		if values != nil {
			clauseValues = append(clauseValues, values...)
		}
	}

	whereClause := strings.Join(clauses, " "+logicOperator+" ")
	return whereClause, clauseValues
}

func buildClause(q query.Query, tblAlias string, placeholderCounter *int) (string, []interface{}) {
	fieldName := q.GetName()
	if tblAlias != "" {
		fieldName = tblAlias + "." + fieldName
	}
	operator, _ := getOperatorInfo(q.GetOperator())
	operand := q.GetOperand()

	switch operator {
	case SqlOperatorIn, SqlOperatorNotIn:
		values := strings.Split(operand.Value.(string), ",")
		placeholders := make([]string, len(values))
		clauseValues := make([]interface{}, len(values))
		for i := range values {
			*placeholderCounter++
			placeholders[i] = fmt.Sprintf("$%d", *placeholderCounter)
			clauseValues[i] = values[i]
		}
		clause := fmt.Sprintf("%s %s (%s)", fieldName, operator, strings.Join(placeholders, ", "))
		return clause, clauseValues
	case SqlOperatorEmpty, SqlOperatorNotEmpty:
		clause := fmt.Sprintf("%s %s ", fieldName, operator)
		return clause, nil
	case SqlOperatorContain, SqlOperatorNotContain:
		*placeholderCounter++
		clause := fmt.Sprintf("%s %s '%%' || $%d || '%%'", fieldName, operator, *placeholderCounter)
		return clause, []interface{}{operand.Value}
	default:
		*placeholderCounter++
		clause := fmt.Sprintf("%s %s $%d", fieldName, operator, *placeholderCounter)
		return clause, []interface{}{operand.Value}
	}
}
