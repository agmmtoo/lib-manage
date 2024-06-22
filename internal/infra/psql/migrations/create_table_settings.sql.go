package migrations

const CreateTableSetting = `
CREATE TABLE IF NOT EXISTS settings (
    id SERIAL PRIMARY KEY,
    library_id INT NOT NULL REFERENCES library(id),
    key VARCHAR(255) NOT NULL,
    value VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    UNIQUE (library_id, key)
);
`
