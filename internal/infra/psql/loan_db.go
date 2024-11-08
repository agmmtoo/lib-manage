package psql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/agmmtoo/lib-manage/internal/core"
	"github.com/agmmtoo/lib-manage/internal/core/loan"
	cm "github.com/agmmtoo/lib-manage/internal/core/models"
	"github.com/agmmtoo/lib-manage/internal/infra/psql/models"
	"github.com/jackc/pgx/v5/pgconn"
)

func (l *LibraryAppDB) ListLoans(ctx context.Context, input loan.ListRequest) (*loan.ListResponse, error) {

	qb := &QueryBuilder{
		Table:        "loans l",
		ParamCounter: 1,
		Cols: []string{
			"l.id", "l.library_book_id", "l.subscription_id", "l.staff_id", "l.loan_date", "l.due_date", "l.return_date", "l.created_at", "l.updated_at", "l.deleted_at",
		},
	}

	if input.IncludeLibraryBook {
		qb.JoinTables = append(qb.JoinTables, "JOIN libraries_books lb ON l.library_book_id = lb.id")
		qb.Cols = append(qb.Cols, "lb.id", "lb.library_id", "lb.book_id", "lb.code")
		qb.JoinTables = append(qb.JoinTables, "JOIN libraries lbl ON lb.library_id = lbl.id")
		qb.Cols = append(qb.Cols, "lbl.id", "lbl.name")
		qb.JoinTables = append(qb.JoinTables, "JOIN books b ON lb.book_id = b.id")
		qb.Cols = append(qb.Cols, "b.id", "b.title", "b.author", "b.sub_category_id")
		qb.JoinTables = append(qb.JoinTables, "JOIN sub_categories sc ON b.sub_category_id = sc.id")
		qb.Cols = append(qb.Cols, "sc.id", "sc.category_id", "sc.name")
		qb.JoinTables = append(qb.JoinTables, "JOIN categories c ON sc.category_id = c.id")
		qb.Cols = append(qb.Cols, "c.id", "c.name")
	}

	if input.IncludeSubscription {
		qb.JoinTables = append(qb.JoinTables, "JOIN subscriptions s ON l.subscription_id = s.id")
		qb.Cols = append(qb.Cols, "s.id", "s.user_id", "s.membership_id", "s.expiry_date")
		qb.JoinTables = append(qb.JoinTables, "JOIN users u ON s.user_id = u.id")
		qb.Cols = append(qb.Cols, "u.id", "u.username")
		qb.JoinTables = append(qb.JoinTables, "JOIN memberships m ON s.membership_id = m.id")
		qb.Cols = append(qb.Cols, "m.id", "m.library_id", "m.name", "m.duration_days", "m.active_loan_limit", "m.fine_per_day")
		qb.JoinTables = append(qb.JoinTables, "JOIN libraries lib ON m.library_id = lib.id")
		qb.Cols = append(qb.Cols, "lib.id", "lib.name")
	}

	if input.IncludeStaff {
		qb.JoinTables = append(qb.JoinTables, "JOIN staffs st ON l.staff_id = st.id")
		qb.Cols = append(qb.Cols, "st.id", "st.user_id", "st.library_id")
		qb.JoinTables = append(qb.JoinTables, "JOIN users stu ON st.user_id = stu.id")
		qb.Cols = append(qb.Cols, "stu.id", "stu.username")
		qb.JoinTables = append(qb.JoinTables, "JOIN libraries stl ON st.library_id = stl.id")
		qb.Cols = append(qb.Cols, "stl.id", "stl.name")
	}

	if len(input.IDs) > 0 {
		qb.AddClause("l.id = ANY($%d)", input.IDs)
	}
	if input.Active {
		qb.AddClause("l.return_date IS NULL")
	}
	if input.ExpiryDate != nil {
		qb.AddClause("DATE(l.due_date) = DATE($%d)", input.ExpiryDate)
	}

	if input.IncludeSubscription {
		if len(input.UserIDs) > 0 {
			qb.AddClause("s.user_id = ANY($%d)", input.UserIDs)
		}
		if len(input.LibraryIDs) > 0 {
			qb.AddClause("m.library_id = ANY($%d)", input.LibraryIDs)
		}
	}

	if len(input.LibraryBookIDs) > 0 {
		qb.AddClause("l.library_book_id = ANY($%d)", input.LibraryBookIDs)
	}

	qb.SetLimit(input.Limit)
	qb.SetOffset(input.Offset)

	query, params := qb.Build()
	rows, err := l.db.QueryContext(ctx, query, params...)
	if err != nil {
		return nil, core.NewCoreError(core.ErrCodeDBQuery, "error listing loans", err)
	}

	defer rows.Close()

	var loans []*cm.Loan
	for rows.Next() {
		var l models.Loan
		dests := []interface{}{&l.ID, &l.LibraryBookID, &l.SubscriptionID, &l.StaffID, &l.LoanDate, &l.DueDate, &l.ReturnDate, &l.CreatedAt, &l.UpdatedAt, &l.DeletedAt}

		if input.IncludeLibraryBook {
			// "lb.id", "lb.library_id", "lb.book_id", "lb.code"
			dests = append(dests, &l.LibraryBookID, &l.LibraryBookLibraryID, &l.LibraryBookBookID, &l.LibraryBookCode)
			// "lbl.id", "lbl.name"
			dests = append(dests, &l.LibraryBookLibraryID, &l.LibraryBookLibraryName)
			// "b.id", "b.title", "b.author"
			dests = append(dests, &l.LibraryBookBookID, &l.LibraryBookBookTitle, &l.LibraryBookBookAuthor, &l.LibraryBookBookSubCategoryID)
			// "sc.id", "sc.category_id", "sc.name"
			dests = append(dests, &l.LibraryBookBookSubCategoryID, &l.LibraryBookBookSubCategoryCategoryID, &l.LibraryBookBookSubCategoryName)
			// "c.id", "c.name"
			dests = append(dests, &l.LibraryBookBookSubCategoryCategoryID, &l.LibraryBookBookSubCategoryCategoryName)
		}

		if input.IncludeSubscription {
			// "s.id", "s.user_id", "s.membership_id", "s.expiry_date"
			dests = append(dests, &l.SubscriptionID, &l.SubscriptionUserID, &l.SubscriptionMembershipID, &l.SubscriptionExpiryDate)
			// "u.id", "u.username"
			dests = append(dests, &l.SubscriptionUserID, &l.SubscriptionUserUsername)
			// "m.id", "m.library_id", "m.name", "m.duration_days", "m.active_loan_limit", "m.fine_per_day"
			dests = append(dests, &l.SubscriptionMembershipID, &l.SubscriptionMembershipLibraryID, &l.SubscriptionMembershipName, &l.SubscriptionMembershipDurationDays, &l.SubscriptionMembershipActiveLoanLimit, &l.SubscriptionMembershipFinePerDay)
			// "lib.id", "lib.name"
			dests = append(dests, &l.SubscriptionMembershipLibraryID, &l.SubscriptionMembershipLibraryName)
		}

		if input.IncludeStaff {
			// "st.id", "st.user_id", "st.library_id"
			dests = append(dests, &l.StaffID, &l.StaffUserID, &l.StaffLibraryID)
			// "u.id", "u.username"
			dests = append(dests, &l.StaffUserID, &l.StaffUserUsername)
			// "lib.id", "lib.name"
			dests = append(dests, &l.StaffLibraryID, &l.StaffLibraryName)
		}

		err := rows.Scan(dests...)
		if err != nil {
			return nil, core.NewCoreError(core.ErrCodeDBScan, "error scanning loan", err)
		}
		loans = append(loans, l.ToCoreModel())
	}

	if err := rows.Err(); err != nil {
		return nil, core.NewCoreError(core.ErrCodeDBQuery, "error listing loans", err)
	}

	return &loan.ListResponse{
		Loans: loans,
	}, nil
}

func (l *LibraryAppDB) GetLoanByID(ctx context.Context, id int) (*cm.Loan, error) {
	qb := &QueryBuilder{
		Table:        "loan l",
		ParamCounter: 1,
		Cols: []string{
			"l.id", "l.library_book_id", "l.subscription_id", "l.staff_id", "l.loan_date", "l.due_date", "l.return_date", "l.created_at", "l.updated_at", "l.deleted_at",
		},
	}

	qb.JoinTables = append(qb.JoinTables, "JOIN libraries_books lb ON l.library_book_id = lb.id JOIN libraries lbl ON lb.library_id = lbl.id JOIN books lbb ON lb.book_id = lbb.id JOIN sub_categories lbbsc ON lbb.sub_category_id = lbbsc.id JOIN categories lbbscc ON lbbsc.category_id = lbbscc.id")
	qb.Cols = append(qb.Cols, "lb.id", "lb.library_id", "lb.book_id", "lb.code", "lbl.id", "lbl.name", "lbb.id", "lbb.title", "lbb.author", "lbb.sub_category_id", "lbbsc.id", "lbbsc.category_id", "lbbsc.name", "lbbscc.id", "lbbscc.name")

	qb.JoinTables = append(qb.JoinTables, "JOIN subscriptions s ON l.subscription_id = s.id JOIN users su ON s.user_id = su.id JOIN memberships sm ON s.membership_id = sm.id JOIN libraries sml ON sm.library_id = sml.id")
	qb.Cols = append(qb.Cols, "s.id", "s.user_id", "s.membership_id", "s.expiry_date", "su.id", "su.username", "sm.id", "sm.library_id", "sm.name", "sm.duration_days", "sm.active_loan_limit", "sm.fine_per_day", "sml.id", "sml.name")

	qb.JoinTables = append(qb.JoinTables, "JOIN staffs st ON l.staff_id = st.id JOIN users stu ON st.user_id = stu.id JOIN libraries stl ON st.library_id = stl.id")
	qb.Cols = append(qb.Cols, "st.id", "st.user_id", "st.library_id", "stu.id", "stu.username", "stl.id", "stl.name")

	qb.AddClause("l.id = $%d", id)

	q, params := qb.Build()

	row := l.db.QueryRowContext(ctx, q, params...)

	var lo models.Loan
	err := row.Scan(
		&lo.ID, &lo.LibraryBookID, &lo.SubscriptionID, &lo.StaffID, &lo.LoanDate, &lo.DueDate, &lo.ReturnDate, &lo.CreatedAt, &lo.UpdatedAt, &lo.DeletedAt,
		&lo.LibraryBookID, &lo.LibraryBookLibraryID, &lo.LibraryBookBookID, &lo.LibraryBookCode, &lo.LibraryBookLibraryID, &lo.LibraryBookLibraryName, &lo.LibraryBookBookID, &lo.LibraryBookBookTitle, &lo.LibraryBookBookAuthor, &lo.LibraryBookBookSubCategoryID, &lo.LibraryBookBookSubCategoryID, &lo.LibraryBookBookSubCategoryCategoryID, &lo.LibraryBookBookSubCategoryName, &lo.LibraryBookBookSubCategoryCategoryID, &lo.LibraryBookBookSubCategoryCategoryName,
		&lo.SubscriptionID, &lo.SubscriptionUserID, &lo.SubscriptionMembershipID, &lo.SubscriptionExpiryDate, &lo.SubscriptionUserID, &lo.SubscriptionUserUsername, &lo.SubscriptionMembershipID, &lo.SubscriptionMembershipLibraryID, &lo.SubscriptionMembershipName, &lo.SubscriptionMembershipDurationDays, &lo.SubscriptionMembershipActiveLoanLimit, &lo.SubscriptionMembershipFinePerDay, &lo.SubscriptionMembershipLibraryID, &lo.SubscriptionMembershipLibraryName,
		&lo.StaffID, &lo.StaffUserID, &lo.StaffLibraryID, &lo.StaffUserID, &lo.StaffUserUsername, &lo.StaffLibraryID, &lo.StaffLibraryName,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, core.NewCoreError(core.ErrCodeDBNotFound, "loan not found", err)
		}
		return nil, core.NewCoreError(core.ErrCodeDBQuery, "error getting loan", err)
	}

	return lo.ToCoreModel(), nil
}

func (l *LibraryAppDB) CreateLoan(ctx context.Context, input loan.CreateRequest) (*cm.Loan, error) {
	q := "INSERT INTO loans (library_book_id, subscription_id, staff_id, loan_date, due_date, return_date) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, library_book_id, subscription_id, staff_id, loan_date, due_date, return_date, created_at, updated_at, deleted_at;"
	args := []any{input.LibraryBookID, input.SubscriptionID, input.StaffID, input.LoanDate, input.DueDate, input.ReturnDate}

	row := l.db.QueryRowContext(ctx, q, args...)

	var lo models.Loan
	err := row.Scan(&lo.ID, &lo.LibraryBookID, &lo.SubscriptionID, &lo.StaffID, &lo.LoanDate, &lo.DueDate, &lo.ReturnDate, &lo.CreatedAt, &lo.UpdatedAt, &lo.DeletedAt)

	// TODO: handle pg error (old code)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			fmt.Println(pgErr.Detail)
			if pgErr.ConstraintName == "loan_user_id_fkey" {
				return nil, core.NewCoreError(core.ErrCodeDBNotFound, "user not found", err)
			}
			if pgErr.ConstraintName == "loan_library_id_staff_id_fkey" {
				return nil, core.NewCoreError(core.ErrCodeDBNotFound, "library or staff not found", err)
			}
			if pgErr.ConstraintName == "loan_library_id_book_id_fkey" {
				return nil, core.NewCoreError(core.ErrCodeDBNotFound, "library or book not found", err)
			}
			if pgErr.ConstraintName == "loan_unique_active" {
				return nil, core.NewCoreError(core.ErrCodeDBDuplicate, "book already loaned", err)
			}
		}
		return nil, core.NewCoreError(core.ErrCodeDBScan, "error creating loan", err)
	}

	return lo.ToCoreModel(), nil
}
