package stsql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const driver = "postgres"
const conStr = "postgres://postgres@localhost:5432/FINWIZ?sslmode=disable"

func Connect() {
	db, err := sql.Open(driver, conStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()

	if err != nil {
		panic(err)
	}
	fmt.Println("Ping OK")
}
