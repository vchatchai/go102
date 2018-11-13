package stsql

import (
	"fmt"
)

func Metadata() {
	db := createConnnection()
	defer db.Close()

	rs, err := db.Query(sel)
	if err != nil {
		panic(err)
	}

	defer rs.Close()
	columns, err := rs.Columns()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Selected columns: %v\n", columns)

	colTypes, err := rs.ColumnTypes()

	if err != nil {
		panic(err)
	}

	for _, col := range colTypes {
		fmt.Println()
		fmt.Printf("%+v\n", col)
	}

}
