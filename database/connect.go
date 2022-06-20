package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

const (
	host     = "172.30.1.25"
	port     = "5432"
	user     = "postgres"
	password = "Dlwpak1024!!"
	dbname   = "shoppingmall"
)

func Connect() error {
	var err error
	DB, err = sql.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname))
	if err != nil {
		return err
	}
	if err = DB.Ping(); err != nil {
		return err
	}
	return nil
}
