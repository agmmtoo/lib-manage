package migrations

const CreateTableSetting = `
CREATE TABLE IF NOT EXISTS setting (
    id SERIAL PRIMARY KEY,
    library_id INT NOT NULL REFERENCES library(id),
    key VARCHAR(255) NOT NULL,
    value VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    UNIQUE (library_id, key)
);

CREATE OR REPLACE TRIGGER update_setting_updated_at
BEFORE UPDATE ON setting
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();
`
