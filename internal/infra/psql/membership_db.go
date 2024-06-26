package psql

import (
	"context"
	"fmt"

	"github.com/agmmtoo/lib-manage/internal/core/membership"
	cm "github.com/agmmtoo/lib-manage/internal/core/models"
	"github.com/agmmtoo/lib-manage/internal/infra/psql/models"
)

func (l *LibraryAppDB) ListMemberships(ctx context.Context, input membership.ListRequest) (*membership.ListResponse, error) {
	qb := QueryBuilder{
		Table:        "memberships m",
		ParamCounter: 1,
		Cols: []string{
			"m.id, m.library_id", "m.name", "m.duration_days", "m.active_loan_limit", "m.fine_per_day", "m.created_at", "m.updated_at", "m.deleted_at",
			"l.id", "l.name",
		},
	}

	// Join with "libraries"
	qb.JoinTables = append(qb.JoinTables, "JOIN libraries l ON m.library_id = l.id")

	if len(input.IDs) > 0 {
		qb.AddClause("m.id = ANY($%d)", input.IDs)
	}
	if len(input.LibraryIDs) > 0 {
		qb.AddClause("m.library_id = ANY($%d)", input.LibraryIDs)
	}
	if input.Name != "" {
		qb.AddClause("m.name ILIKE $%d", fmt.Sprintf("%%%s%%", input.Name))
	}
	if input.DurationDays != nil {
		qb.AddClause("m.duration_days = $%d", input.DurationDays)
	}
	if input.ActiveLoanLimit != nil {
		qb.AddClause("m.active_loan_limit = $%d", input.ActiveLoanLimit)
	}
	if input.FinePerDay != nil {
		qb.AddClause("m.fine_per_day = $%d", input.FinePerDay)
	}
	qb.SetLimit(input.Limit)
	qb.SetOffset(input.Offset)
	q, args := qb.Build()

	rows, err := l.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var memberships []*cm.Membership
	for rows.Next() {
		var m models.Membership
		err := rows.Scan(
			&m.ID, &m.LibraryID, &m.Name, &m.DurationDays, &m.ActiveLoanLimit, &m.FinePerDay, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt,
			&m.LibraryID, &m.LibraryName,
		)
		if err != nil {
			return nil, err
		}
		memberships = append(memberships, m.ToCoreModel())
	}

	return &membership.ListResponse{
		Memberships: memberships,
	}, nil
}
