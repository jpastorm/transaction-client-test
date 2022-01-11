CREATE TABLE transaction (
	id SERIAL NOT NULL,
	money NUMERIC(SIZE) NOT NULL,
	type VARCHAR(SIZE) NOT NULL,
	account_holder INTEGER NOT NULL,
	subject INTEGER NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT now(),
	updated_at TIMESTAMP,
	CONSTRAINT transaction_id_pk PRIMARY KEY (id)
);

COMMENT ON TABLE transaction IS 'Write your comment here';

-- Register the permission module for the routes
INSERT INTO modules (name) VALUES ('TRANSACTION');
