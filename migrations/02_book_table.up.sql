CREATE TABLE IF NOT EXISTS book(
    id uuid PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    category_id uuid REFERENCES book_category(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT(Now()),
	updated_at TIMESTAMP DEFAULT(Now())
);