package driver

import "database/sql"

// dbRepo for postgres database struct
type dbRepo struct {
	db *sql.DB
}

// NewDBRepo
func NewDBRepo(db *sql.DB) *dbRepo {
	return &dbRepo{
		db: db,
	}
}
