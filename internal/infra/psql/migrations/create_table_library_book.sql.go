package migrations

const CreateTableLibraryBook = `
CREATE TABLE IF NOT EXISTS library_book (
    library_id INT NOT NULL REFERENCES library(id),
    book_id INT NOT NULL REFERENCES book(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (library_id, book_id)
);

CREATE OR REPLACE TRIGGER update_library_book_updated_at
BEFORE UPDATE ON library_book
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();
`
