package migrations

const CreateTableBook = `
CREATE TABLE IF NOT EXISTS book (
    id SERIAL PRIMARY KEY,
    title VARCHAR(225) NOT NULL,
    author VARCHAR(225) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP
);

CREATE OR REPLACE TRIGGER update_book_updated_at
BEFORE UPDATE ON book
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();
`
