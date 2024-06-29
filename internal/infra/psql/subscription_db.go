package psql

import (
	"context"
	"fmt"
	"strings"

	cm "github.com/agmmtoo/lib-manage/internal/core/models"
	"github.com/agmmtoo/lib-manage/internal/core/subscription"
	"github.com/agmmtoo/lib-manage/internal/infra/psql/models"
)

func (l *LibraryAppDB) ListSubscriptions(ctx context.Context, input subscription.ListRequest) (*subscription.ListResponse, error) {
	qb := QueryBuilder{
		Table:        "subscriptions s",
		ParamCounter: 1,
		Cols: []string{
			"s.id, s.user_id", "s.membership_id", "s.expiry_date", "s.created_at", "s.updated_at", "s.deleted_at",
			"u.id", "u.username",
			"m.id, m.library_id", "m.name", "m.duration_days", "m.active_loan_limit", "m.fine_per_day",
			"l.id", "l.name",
		},
	}

	// Join with "users"
	qb.JoinTables = append(qb.JoinTables, "JOIN users u ON s.user_id = u.id")

	// Join with "memberships"
	qb.JoinTables = append(qb.JoinTables, "JOIN memberships m ON s.membership_id = m.id")

	// Join with "libraries"
	qb.JoinTables = append(qb.JoinTables, "JOIN libraries l ON m.library_id = l.id")

	if len(input.IDs) > 0 {
		qb.AddClause("m.id = ANY($%d)", input.IDs)
	}
	if len(input.UserIDs) > 0 {
		qb.AddClause("s.user_id = ANY($%d)", input.UserIDs)
	}
	if len(input.MembershipIDs) > 0 {
		qb.AddClause("s.membership_id = ANY($%d)", input.MembershipIDs)
	}
	if input.ExpiryDate != nil {
		qb.AddClause("DATE(s.expiry_date) = DATE($%d)", input.ExpiryDate)
	}
	if input.ExpiredBefore != nil {
		qb.AddClause("s.expiry_date < $%d", input.ExpiredBefore)
	}
	if input.ExpiredAfter != nil {
		qb.AddClause("s.expiry_date > $%d", input.ExpiredAfter)
	}
	if len(input.LibraryIDs) > 0 {
		qb.AddClause("m.library_id = ANY($%d)", input.LibraryIDs)
	}
	// if input.Name != "" {
	// 	qb.AddClause("m.name ILIKE $%d", fmt.Sprintf("%%%s%%", input.Name))
	// }

	obParts := make([]string, len(input.OrderBy))
	for i, ob := range input.OrderBy {
		// Default to ASC if Direction is not specified
		dir := "ASC"
		if ob.Dir != "" {
			dir = ob.Dir
		}
		obParts[i] = fmt.Sprintf("%s %s", ob.Col, dir)
	}
	ob := strings.Join(obParts, ", ")
	qb.SetOrderBy(ob)

	qb.SetLimit(input.Limit)
	qb.SetOffset(input.Offset)
	q, args := qb.Build()

	rows, err := l.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var subscriptions []*cm.Subscription
	for rows.Next() {
		var s models.Subscription
		err := rows.Scan(
			&s.ID, &s.UserID, &s.MembershipID, &s.ExpiryDate, &s.CreatedAt, &s.UpdatedAt, &s.DeletedAt,
			&s.UserID, &s.UserUsername,
			&s.MembershipID, &s.MembershipLibraryID, &s.MembershipName, &s.MembershipDurationDays, &s.MembershipActiveLoanLimit, &s.MembershipFinePerDay,
			&s.MembershipLibraryID, &s.MembershipLibraryName,
		)
		if err != nil {
			return nil, err
		}
		subscriptions = append(subscriptions, s.ToCoreModel())
	}

	return &subscription.ListResponse{
		Subscriptions: subscriptions,
	}, nil
}

func (l *LibraryAppDB) CreateSubscription(ctx context.Context, input subscription.CreateRequest) (*cm.Subscription, error) {
	q := `
		INSERT INTO subscriptions (user_id, membership_id, expiry_date)
		VALUES ($1, $2, $3)
		RETURNING id, user_id, membership_id, expiry_date, created_at, updated_at, deleted_at
	`
	var s models.Subscription
	err := l.db.QueryRowContext(ctx, q, input.UserID, input.MembershipID, input.ExpiryDate).Scan(
		&s.ID, &s.UserID, &s.MembershipID, &s.ExpiryDate, &s.CreatedAt, &s.UpdatedAt, &s.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	return s.ToCoreModel(), nil
}
