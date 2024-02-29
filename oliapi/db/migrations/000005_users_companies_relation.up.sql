CREATE TABLE users_companies (
	user_id UUID,
	company_id UUID,
	PRIMARY KEY (user_id, company_id),
	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
	FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE
);
