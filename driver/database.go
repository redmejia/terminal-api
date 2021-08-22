package driver

import "database/sql"

// dbRepo for postgres database
type dbRepo struct {
	conn *sql.DB
}

func NewDBRepo(db *sql.DB) *dbRepo {
	return &dbRepo{
		conn: db,
	}
}
