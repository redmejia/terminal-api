package driver

import "database/sql"

// dbRepo for postgres database
type DbRepo struct {
	Conn *sql.DB
}
