CREATE TYPE type_transaction AS ENUM ('DEPOSIT', 'WITHDRAW', 'TRANSFER');

CREATE TABLE client (
	id SERIAL primary key ,
	name varchar(50),
	created_at TIMESTAMP NOT NULL DEFAULT now(),
	updated_at TIMESTAMP
);

CREATE TABLE currency (
	id SERIAL primary key ,
	name varchar(50) NOT NULL,
	code varchar(50) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT now(),
	updated_at TIMESTAMP
);

INSERT INTO currency(name, code) VALUES
('American dollar', 'USD'),('COP','Colombian Peso'),('MXN', 'Mexico peso');


CREATE TABLE account (
	id SERIAL primary key ,
	client_id int,
	currency_id int,
	money decimal(10,2),
	FOREIGN KEY (client_id) REFERENCES client(id),
	FOREIGN KEY (currency_id) REFERENCES currency(id),
	created_at TIMESTAMP NOT NULL DEFAULT now(),
	updated_at TIMESTAMP
);

CREATE TABLE transaction (
	id SERIAL primary key ,
	money decimal(10,2),
	type type_transaction,
	account_holder int,
	subject int null,
	FOREIGN KEY (account_holder) REFERENCES account(id),
	created_at TIMESTAMP NOT NULL DEFAULT now(),
	updated_at TIMESTAMP
);
