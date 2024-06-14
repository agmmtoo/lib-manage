package migrations

const CreateTableCategory = `
CREATE TABLE IF NOT EXISTS "category" (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL UNIQUE,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP
);

CREATE OR REPLACE TRIGGER update_category_updated_at
BEFORE UPDATE ON "category"
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();
`

const CreateTableSubCategory = `
CREATE TABLE IF NOT EXISTS "sub_category" (
	id SERIAL PRIMARY KEY,
	category_id INT NOT NULL REFERENCES "category"(id)
	name VARCHAR(255) NOT NULL UNIQUE
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP
);

CREATE OR REPLACE TRIGGER update_sub_category_updated_at
BEFORE UPDATE ON "sub_category"
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();
`

const CreateTableCategorySubCategory = `
CREATE TABLE IF NOT EXISTS "category_sub_category" (
	category_id INT NOT NULL REFERENCES "category"(id),
	sub_category_id INT NOT NULL REFERENCES "sub_category"(id),
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP,
	PRIMARY KEY (category_id, sub_category_id)
);

CREATE OR REPLACE TRIGGER update_category_sub_category_updated_at
BEFORE UPDATE ON "category_sub_category"
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();
`
