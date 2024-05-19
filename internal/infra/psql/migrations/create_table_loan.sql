
CREATE TABLE IF NOT EXISTS loan (
    id SERIAL PRIMARY KEY,
    book_id INT NOT NULL REFERENCES book(id),
    user_id INT NOT NULL REFERENCES "user"(id),
    library_id INT NOT NULL,
    staff_id INT NOT NULL,
    FOREIGN KEY (library_id, staff_id) REFERENCES library_staff(library_id, staff_id),
    loan_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    due_date TIMESTAMP NOT NULL,
    return_date TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE OR REPLACE TRIGGER update_loan_updated_at
BEFORE UPDATE ON loan
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();
