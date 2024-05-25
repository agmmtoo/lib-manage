
CREATE TABLE IF NOT EXISTS loan (
    id SERIAL PRIMARY KEY,
    book_id INT NOT NULL,
    user_id INT NOT NULL REFERENCES "user"(id),
    library_id INT NOT NULL,
    staff_id INT NOT NULL,
    loan_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    due_date TIMESTAMP NOT NULL,
    return_date TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (library_id, staff_id) REFERENCES library_staff(library_id, staff_id),
    FOREIGN KEY (library_id, book_id) REFERENCES library_book(library_id, book_id)
);

-- Add a partial unique index to enforce the constraint
CREATE UNIQUE INDEX IF NOT EXISTS loan_unique_active
ON loan (library_id, book_id)
WHERE return_date IS NULL;

CREATE OR REPLACE TRIGGER update_loan_updated_at
BEFORE UPDATE ON loan
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();

-- CREATE OR REPLACE FUNCTION can_borrow_book(
--     p_user_id INT, 
--     p_library_id INT
-- ) RETURNS BOOLEAN AS $$
-- DECLARE
--     active_loans_count INT;
--     max_loans_per_user INT;
-- BEGIN
--     -- Step 1: Check active loans count
--     SELECT COUNT(*) INTO active_loans_count
--     FROM loan 
--     WHERE user_id = p_user_id 
--     AND library_id = p_library_id 
--     AND return_date IS NULL;

--     -- Step 2: Retrieve max loans per user setting
--     SELECT value::int INTO max_loans_per_user
--     FROM settings 
--     WHERE key = 'max_loans_per_user';

--     -- Step 3: Compare and return
--     IF active_loans_count < max_loans_per_user THEN
--         RETURN TRUE;
--     ELSE
--         RETURN FALSE;
--     END IF;
-- END;
-- $$ LANGUAGE plpgsql;

-- SELECT can_borrow_book(123, 1); -- Returns true or false

