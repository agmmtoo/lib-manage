package migrations

const CreateTableLibrariesBooks = `
CREATE TABLE IF NOT EXISTS libraries_books (
    id SERIAL PRIMARY KEY,
    library_id INT NOT NULL REFERENCES libraries(id),
    book_id INT NOT NULL REFERENCES books(id),
    code VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    UNIQUE (library_id, book_id, code)
);
`
