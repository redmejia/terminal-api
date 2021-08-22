package driver

import "database/sql"

// Database for postgres database
type Database struct {
	Conn *sql.DB
}
