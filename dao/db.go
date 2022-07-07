package dao

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

func Connection() *sql.DB {
	config := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 fmt.Sprint(os.Getenv("DB_HOST"), ":", os.Getenv("DB_PORT")),
		DBName:               os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", config.FormatDSN())

	if err != nil {
		log.Fatalf("failed to connect db: %s", err.Error())

	} else if err = db.Ping(); err != nil {
		log.Fatalf("failed to ping db: %s" + err.Error())
	}
	return db
}
