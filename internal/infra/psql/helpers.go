package psql

import (
	"fmt"
	"strings"
)

type QueryBuilder struct {
	Table        string
	Clauses      []string
	Params       []interface{}
	Limit        int
	Offset       int
	ParamCounter int
	Cols         []string
}

func (qb *QueryBuilder) AddClause(clause string, params ...interface{}) {
	if len(params) > 0 {
		clause = fmt.Sprintf(clause, qb.ParamCounter)
		qb.Params = append(qb.Params, params...)
		qb.ParamCounter++
	}
	qb.Clauses = append(qb.Clauses, clause)
}

func (qb *QueryBuilder) SetLimit(limit int) {
	qb.Limit = limit
}

func (qb *QueryBuilder) SetOffset(offset int) {
	qb.Offset = offset
}

func (qb *QueryBuilder) Build() (string, []interface{}) {
	fields := "*"
	if len(qb.Cols) > 0 {
		fields = strings.Join(qb.Cols, ", ")
	}
	query := fmt.Sprintf("SELECT %s FROM %s", fields, qb.Table)

	if len(qb.Clauses) > 0 {
		query += " WHERE " + strings.Join(qb.Clauses, " AND ")
	}
	if qb.Limit > 0 {
		query += fmt.Sprintf(" LIMIT $%d", len(qb.Params)+1)
		qb.Params = append(qb.Params, qb.Limit)
	}
	if qb.Offset > 0 {
		query += fmt.Sprintf(" OFFSET $%d", len(qb.Params)+1)
		qb.Params = append(qb.Params, qb.Offset)
	}

	fmt.Println("QueryBuilder: ", query, qb.Params)

	return query, qb.Params
}
