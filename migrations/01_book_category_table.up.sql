CREATE TABLE IF NOT EXISTS book_category(
    id uuid PRIMARY KEY , 
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT(Now()),
	updated_at TIMESTAMP DEFAULT(Now())
);
