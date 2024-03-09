CREATE TABLE documents (
	id UUID,
	category_id UUID NOT NULL,
	content TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	embedding_date TIMESTAMP DEFAULT NULL,
	deleted_at TIMESTAMP DEFAULT NULL,
	archived_at TIMESTAMP DEFAULT NULL,
	FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE,
	PRIMARY KEY (id)
);
