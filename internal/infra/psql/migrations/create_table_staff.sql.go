package migrations

const CreateTableStaff = `
CREATE TABLE IF NOT EXISTS staff (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL UNIQUE REFERENCES "user"(id),
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP
);

CREATE OR REPLACE TRIGGER update_staff_updated_at
BEFORE UPDATE ON staff
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();

-- ALTER TABLE staff
-- ADD CONSTRAINT fk_staff_users
-- FOREIGN KEY (user_id) REFERENCES users(id);
`
