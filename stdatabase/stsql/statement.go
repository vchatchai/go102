package stsql

import (
	"database/sql"
	"fmt"
)

const sel = "SELECT * FROM post;"
const trunc = "TRUNCATE TABLE post;"
const ins = "INSERT INTO post(ID,TITLE, CONTENT) VALUES (1,'Title 1' , 'Content 1'), (2,  'Title 2', 'Content 2') "

func Statment() {

	db := createConnnection()
	defer db.Close()

	_, err := db.Exec(trunc)
	if err != nil {
		panic(err)
	}

	fmt.Println("Table truncated")
	r, err := db.Exec(ins)
	if err != nil {
		panic(err)
	}
	affected, err := r.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Inserted rows count: %d\n", affected)
	rs, err := db.Query(sel)
	if err != nil {
		panic(err)
	}

	count := 0

	for rs.Next() {
		count++

	}
	fmt.Printf("Total of %d was selected.\n", count)
}

func createConnnection() *sql.DB {
	db, err := sql.Open("postgres", conStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db

}
