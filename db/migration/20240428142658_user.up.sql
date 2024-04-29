CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	email VARCHAR(100) NOT NULL UNIQUE,
	name VARCHAR(50) NOT NULL,
	password VARCHAR(15) NOT NULL,
	CONSTRAINT name_length CHECK (LENGTH(name) >= 5 AND LENGTH(name) <= 50),
	-- CONSTRAINT password_length CHECK (LENGTH(password) >= 5 AND LENGTH(password) <= 15),
	CONSTRAINT valid_email CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$')
);
