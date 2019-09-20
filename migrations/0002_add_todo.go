package migrations

import (
	"database/sql"
)

var addTodoMigration_0002 = &Migration{
	Number: 2,
	Name:   "Add todo",
	Forwards: func(db *sql.Tx) error {
		const addSQL = `
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
		`

		_, err := db.Exec(addSQL)

		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	Migrations = append(Migrations, addTodoMigration_0002)
}
