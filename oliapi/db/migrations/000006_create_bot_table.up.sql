CREATE TABLE bots (
	id UUID,
	company_id UUID NOT NULL,
	name VARCHAR(255) NOT NULL,
	greeting_message TEXT NOT NULL,
	custom_propmt TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	deleted_at TIMESTAMP DEFAULT NULL,
	archived_at TIMESTAMP DEFAULT NULL,
	FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE,
	PRIMARY KEY (id)
);
