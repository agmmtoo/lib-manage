
CREATE TABLE IF NOT EXISTS categories (
	id UUID PRIMARY KEY,
	name VARCHAR(255) NOT NULL UNIQUE,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS sub_categories (
	id UUID PRIMARY KEY,
	category_id UUID NOT NULL REFERENCES categories(id),
	name VARCHAR(255) NOT NULL UNIQUE,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS libraries (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS books (
    id UUID PRIMARY KEY,
    title VARCHAR(225) NOT NULL,
    author VARCHAR(225) NOT NULL,
	sub_category_id UUID REFERENCES sub_categories(id),
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS libraries_books (
    id UUID PRIMARY KEY,
    library_id UUID NOT NULL REFERENCES libraries(id),
    book_id UUID NOT NULL REFERENCES books(id),
    code VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    UNIQUE (library_id, code)
);

CREATE TABLE IF NOT EXISTS users (
	id UUID PRIMARY KEY,
	username VARCHAR(255) NOT NULL UNIQUE,
	password VARCHAR(255) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS staffs (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id),
	library_id UUID NOT NULL REFERENCES libraries(id),
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP,
	UNIQUE (user_id, library_id)
);

CREATE TABLE IF NOT EXISTS memberships (
	id UUID PRIMARY KEY,
	library_id UUID NOT NULL REFERENCES libraries(id),
	name VARCHAR(255) NOT NULL,
	duration_days INT NOT NULL,
	active_loan_limit INT NOT NULL,
	fine_per_day INT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS subscriptions (
	id UUID PRIMARY KEY,
	user_id UUID NOT NULL REFERENCES users(id),
	membership_id UUID NOT NULL REFERENCES memberships(id),
	-- since membership can be edited, we need to store the expiry date
	expiry_date TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS loans (
    id UUID PRIMARY KEY,
    library_book_id UUID NOT NULL REFERENCES libraries_books(id),
    subscription_id UUID NOT NULL REFERENCES subscriptions(id),
    staff_id UUID NOT NULL REFERENCES staffs(id),
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

CREATE TABLE IF NOT EXISTS settings (
    id UUID PRIMARY KEY,
    library_id UUID NOT NULL REFERENCES libraries(id),
    key VARCHAR(255) NOT NULL,
    value VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    UNIQUE (library_id, key)
);
