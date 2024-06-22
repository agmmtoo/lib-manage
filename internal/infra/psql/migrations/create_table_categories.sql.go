package migrations

const CreateTableCategories = `
CREATE TABLE IF NOT EXISTS categories (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL UNIQUE,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP
);
`

const CreateTableSubCategories = `
CREATE TABLE IF NOT EXISTS sub_categories (
	id SERIAL PRIMARY KEY,
	category_id INT NOT NULL REFERENCES categories(id)
	name VARCHAR(255) NOT NULL UNIQUE
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP
);
`

const CreateTableCategoriesSubCategories = `
CREATE TABLE IF NOT EXISTS category_sub_categories (
	category_id INT NOT NULL REFERENCES categories(id),
	sub_category_id INT NOT NULL REFERENCES sub_categories(id),
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP,
	PRIMARY KEY (category_id, sub_category_id)
);
`
