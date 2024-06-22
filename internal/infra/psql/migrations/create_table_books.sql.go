package migrations

const CreateTableBooks = `
CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(225) NOT NULL,
    author VARCHAR(225) NOT NULL,
	category_id INT REFERENCES categories(id),
	sub_category_id INT REFERENCES sub_categories(id),
	FOREIGN KEY (category_id, sub_category_id) REFERENCES categories_sub_categories(category_id, sub_category_id),
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP
);
`
