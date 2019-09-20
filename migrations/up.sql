CREATE TABLE users (
	id serial PRIMARY KEY,
	email text UNIQUE NOT NULL,
	hashed_password bytea NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	deleted_at TIMESTAMP
);

CREATE TABLE todos (
	id serial PRIMARY KEY,
	name text NOT NULL,
	done boolean NOT NULL,
	user_id int NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	deleted_at TIMESTAMP,
	CONSTRAINT todos_user_id_fkey FOREIGN KEY (user_id)
    REFERENCES users (id) MATCH SIMPLE
    ON UPDATE NO ACTION ON DELETE CASCADE
);
