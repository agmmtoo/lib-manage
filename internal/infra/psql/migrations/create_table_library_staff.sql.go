package migrations

const CreateTableLibraryStaff = `
CREATE TABLE IF NOT EXISTS library_staff (
    library_id INT NOT NULL REFERENCES library(id),
    staff_id INT NOT NULL REFERENCES staff(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (library_id, staff_id)
);

CREATE OR REPLACE TRIGGER update_library_staff_updated_at
BEFORE UPDATE ON library_staff
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();
`
