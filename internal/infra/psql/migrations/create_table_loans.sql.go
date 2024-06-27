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
    subscription_id INT NOT NULL REFERENCES subscriptions(id),
    staff_id INT NOT NULL REFERENCES staffs(id),
    loan_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    due_date TIMESTAMP NOT NULL,
    return_date TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Add a partial unique index to enforce the constraint
CREATE UNIQUE INDEX IF NOT EXISTS loan_unique_active
ON loans (library_book_id)
WHERE return_date IS NULL;
`

// not used
const CreateTriggerLoans = `
CREATE FUNCTION check_library_ids_match() RETURNS TRIGGER AS $$
BEGIN
    IF (
        (SELECT library_id FROM libraries_books WHERE id = NEW.library_book_id) <> 
        (SELECT library_id FROM users_memberships WHERE id = NEW.user_membership_id) OR
        (SELECT library_id FROM libraries_books WHERE id = NEW.library_book_id) <> 
        (SELECT library_id FROM staffs WHERE id = NEW.staff_id)
    ) THEN
        RAISE EXCEPTION 'Library IDs do not match';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER ensure_library_id_match
BEFORE INSERT ON loans
FOR EACH ROW
EXECUTE FUNCTION check_library_ids_match();
`
