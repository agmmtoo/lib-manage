package migrations

const CreateTableStaffs = `
CREATE TABLE IF NOT EXISTS staffs (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
	library_id INT NOT NULL REFERENCES libraries(id),
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP,
	UNIQUE (user_id, library_id)
);
`
