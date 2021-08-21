package driver

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

// hold connection
var DB *sql.DB

const (
	openConns = 10
	idleConns = 3
	lifeTime  = 5 * time.Minute
)

func Conn() (*sql.DB, error) {
	port, _ := strconv.Atoi(os.Getenv("DBPORT"))
	connDB := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("HOSTNAME"), port, os.Getenv("DBUSER"), os.Getenv("DBPASSWORD"),
		os.Getenv("DBNAME"), os.Getenv("DBSSLMODE"),
	)

	DB, err := sql.Open("pgx", connDB)
	if err != nil {
		return nil, err
	}

	DB.SetMaxOpenConns(openConns)
	DB.SetMaxIdleConns(idleConns)
	DB.SetConnMaxLifetime(lifeTime)

	return DB, nil
}
