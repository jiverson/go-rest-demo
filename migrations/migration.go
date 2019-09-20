package migrations

import "database/sql"

type Migration struct {
	Number uint
	Name   string

	Forwards func(db *sql.Tx) error
}

var Migrations []*Migration
