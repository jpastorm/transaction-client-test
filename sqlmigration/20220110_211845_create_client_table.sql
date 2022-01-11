CREATE TABLE client (
	id SERIAL NOT NULL,
	name VARCHAR(SIZE) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT now(),
	updated_at TIMESTAMP,
	CONSTRAINT client_id_pk PRIMARY KEY (id)
);

COMMENT ON TABLE client IS 'Write your comment here';

-- Register the permission module for the routes
INSERT INTO modules (name) VALUES ('CLIENT');
