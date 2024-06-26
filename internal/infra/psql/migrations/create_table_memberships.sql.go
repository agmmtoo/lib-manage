package migrations

const CreateTableMemberships = `
CREATE TABLE IF NOT EXISTS memberships (
	id SERIAL PRIMARY KEY,
	library_id INT NOT NULL REFERENCES libraries(id),
	name VARCHAR(255) NOT NULL,
	duration_days INT NOT NULL,
	active_loan_limit INT NOT NULL,
	fine_per_day INT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP
);
`
