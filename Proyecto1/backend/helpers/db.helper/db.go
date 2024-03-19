package dbhelper

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Db struct {
	Db *sql.DB
}

func NewConnection() (*Db, error) {
	src := "root:@tcp(127.0.0.1:3306)/"
	dbName := "sopes"
	db, err := sql.Open("mysql", src)
	if err != nil {
		return nil, err
	}

	// Crear la base de datos si no existe
	if err := createDatabaseIfNotExists(db, dbName); err != nil {
		db.Close()
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Usamos la base de datos actual
	if _, err := db.Exec(fmt.Sprintf("USE %s", dbName)); err != nil {
		db.Close()
		return nil, err
	}
	return &Db{Db: db}, nil
}

func (db *Db) Close() {
	if db.Db != nil {
		db.Db.Close()
	}
}

func createDatabaseIfNotExists(db *sql.DB, databaseName string) error {
	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", databaseName)
	_, err := db.Exec(query)
	return err
}

func CreateTableIfNotExists(db *Db, tables ...string) error {
	for _, table := range tables {
		query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s", table)
		_, err := db.Db.Exec(query)
		if err != nil {
			return err
		}
	}
	return nil
}

func Execute(db *Db, query string) (sql.Result, error) {
	return db.Db.Exec(query)
}

func Query(db *Db, query string) (*sql.Rows, error) {
	fmt.Println("Entra al query")
	return db.Db.Query(query)
}
