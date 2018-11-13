package stsql

import (
	"fmt"
)

var testTable = []struct {
	ID      int
	Title   string
	Content string
}{
	{1, "Title One", "Content of title  one"},
	{2, "Title Two", "Conetnt of title two"},
	{3, "Title Three", "Content of title three"},
}

const insert = "INSERT INTO post(ID,TITLE, CONTENT) VALUES ($1,$2 , $3) "

func Prepare() {
	db := createConnnection()
	defer db.Close()

	_, err := db.Exec(trunc)
	if err != nil {
		panic(err)
	}

	stm, err := db.Prepare(insert)
	if err != nil {
		panic(err)
	}

	inserted := int64(0)

	for _, val := range testTable {
		fmt.Printf("Inserting record ID: %d\n", val.ID)
		//Executeig the prepared statement
		r, err := stm.Exec(val.ID, val.Title, val.Content)
		if err != nil {
			fmt.Printf("Connot insert record ID: %d\n", val.ID)
		}

		if affected, err := r.RowsAffected(); err == nil {
			inserted = inserted + affected
		}
	}
	fmt.Printf("Result: Insert %d rows.\n", inserted)
}
