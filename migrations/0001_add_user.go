package migrations

import (
	"database/sql"
)

var addUserMigration_0001 = &Migration{
	Number: 1,
	Name:   "Add user",
	Forwards: func(db *sql.Tx) error {
		const addSQL = `
			CREATE TABLE users(
 				id serial PRIMARY KEY,
 				email text UNIQUE NOT NULL,
				hashed_password bytea NOT NULL,
 				created_at TIMESTAMP NOT NULL,
 				updated_at TIMESTAMP NOT NULL,
 				deleted_at TIMESTAMP
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
	Migrations = append(Migrations, addUserMigration_0001)
}
