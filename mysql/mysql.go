package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Dsn(dbName string) string {
	username := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	hostname := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, hostname, port, dbName)
}

func CreatePool() (*sql.DB, error) {
	dbname := os.Getenv("MYSQL_DEFAULT_DB")
	db, err := sql.Open("mysql", Dsn(dbname))
	if err != nil {
		log.Printf("Error %s when opening DB", err)
		return nil, err
	}
	maxOpenConnection, _ := strconv.Atoi(os.Getenv("MYSQL_CONNECTION_LIMIT"))
	db.SetMaxOpenConns(maxOpenConnection)
	db.SetMaxIdleConns(maxOpenConnection)
	db.SetConnMaxLifetime(time.Minute * 5)
	err = db.Ping()
	if err != nil {
		log.Printf("Errors %s pinging DB", err)
		return nil, err
	}
	log.Printf("Connected to DB %s successfully\n", dbname)
	return db, nil
}
