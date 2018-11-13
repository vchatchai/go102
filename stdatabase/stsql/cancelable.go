package stsql

import (
	"context"
	"fmt"
	"time"
)

const sel_cancel = "SELECT * FROM  post p CROSS JOIN (SELECT 1 FROM generate_series(1,1000000)) tbl"

func Cancelable() {
	db := createConnnection()

	defer db.Close()

	ctx, canc := context.WithTimeout(context.Background(), 20*time.Microsecond)
	rows, err := db.QueryContext(ctx, sel_cancel)
	canc()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer rows.Close()

	count := 0
	for rows.Next() {
		if rows.Err() != nil {
			fmt.Println(rows.Err())
			continue
		}
		count++
	}

	fmt.Printf("%d rows returned\n", count)
}
