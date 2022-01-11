CREATE TABLE account (
	id SERIAL NOT NULL,
	client_id INTEGER NOT NULL,
	currency_id INTEGER NOT NULL,
	money NUMERIC(SIZE) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT now(),
	updated_at TIMESTAMP,
	CONSTRAINT account_id_pk PRIMARY KEY (id)
);

COMMENT ON TABLE account IS 'Write your comment here';

-- Register the permission module for the routes
INSERT INTO modules (name) VALUES ('ACCOUNT');
