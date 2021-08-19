package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func Conn() (*sql.DB, error) {

	port, _ := strconv.Atoi(os.Getenv("DBPORT"))
	connDB := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("HOSTNAME"), port, os.Getenv("DBUSER"), os.Getenv("DBPASSWORD"),
		os.Getenv("DBNAME"), os.Getenv("DBSSLMODE"),
	)

	db, err := sql.Open("pgx", connDB)
	if err != nil {
		return nil, err
	}

	return db, nil
}
