CREATE TABLE currency (
	id SERIAL NOT NULL,
	name VARCHAR(SIZE) NOT NULL,
	code VARCHAR(SIZE) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT now(),
	updated_at TIMESTAMP,
	CONSTRAINT currency_id_pk PRIMARY KEY (id)
);

COMMENT ON TABLE currency IS 'Write your comment here';

-- Register the permission module for the routes
INSERT INTO modules (name) VALUES ('CURRENCY');
