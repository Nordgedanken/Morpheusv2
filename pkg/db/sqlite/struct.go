package sqlite

import "database/sql"

type SQLite struct {
	db *sql.DB
}
