package migrations

// const CreateTableLoans = `
// CREATE TABLE IF NOT EXISTS loan (
//     id SERIAL PRIMARY KEY,
//     book_id INT NOT NULL,
//     user_id INT NOT NULL,
//     library_id INT NOT NULL,
//     staff_id INT NOT NULL,
//     loan_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
//     due_date TIMESTAMP NOT NULL,
//     return_date TIMESTAMP,
//     created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
//     updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
//     deleted_at TIMESTAMP,
//     FOREIGN KEY (library_id, staff_id) REFERENCES library_staff(library_id, staff_id),
//     FOREIGN KEY (library_id, book_id) REFERENCES library_book(library_id, book_id),
//     FOREIGN KEY (library_id, user_id) REFERENCES library_user(library_id, user_id)
// );

// -- Add a partial unique index to enforce the constraint
// CREATE UNIQUE INDEX IF NOT EXISTS loan_unique_active
// ON loan (library_id, book_id)
// WHERE return_date IS NULL;
// `

const CreateTableLoans = `
CREATE TABLE IF NOT EXISTS loans (
    id SERIAL PRIMARY KEY,
    library_book_id INT NOT NULL REFERENCES libraries_books(id),
    user_membership_id INT NOT NULL REFERENCES users_memberships(id),
    staff_id INT NOT NULL REFERENCES staffs(id),
    loan_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    due_date TIMESTAMP NOT NULL,
    return_date TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
);

-- Add a partial unique index to enforce the constraint
CREATE UNIQUE INDEX IF NOT EXISTS loan_unique_active
ON loan library_book_id
WHERE return_date IS NULL;
`
