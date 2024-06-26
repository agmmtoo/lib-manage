package migrations

const CreateTableSubscriptions = `
CREATE TABLE IF NOT EXISTS subscriptions (
	id SERIAL PRIMARY KEY,
	user_id INT NOT NULL REFERENCES users(id),
	membership_id INT NOT NULL REFERENCES memberships(id),
	-- since membership can be edited, we need to store the expiry date
	expiry_date TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
`
