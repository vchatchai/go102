package stsql

import (
	"database/sql"
	"fmt"
)

func Querymap() {
	db := createConnnection()
	defer db.Close()

	rows, err := db.Query(selOne, 1)
	if err != nil {
		panic(err)
	}

	cols, _ := rows.Columns()
	for rows.Next() {
		m := parseWithRawBytes(rows, cols)
		fmt.Println(m)
		m = parseToMap(rows, cols)
		fmt.Println(m)
	}
}
func parseWithRawBytes(rows *sql.Rows, cols []string) map[string]interface{} {
	vals := make([]sql.RawBytes, len(cols))
	scanArgs := make([]interface{}, len(vals))
	for i := range vals {
		scanArgs[i] = &vals[i]
	}
	if err := rows.Scan(scanArgs...); err != nil {
		panic(err)
	}
	fmt.Printf("parseWithRawBytes %v  %v\n", scanArgs, vals)

	m := make(map[string]interface{})
	for i, col := range vals {
		if col == nil {
			m[cols[i]] = nil
		} else {
			m[cols[i]] = string(col)
		}
	}
	return m
}

func parseToMap(rows *sql.Rows, cols []string) map[string]interface{} {
	values := make([]interface{}, len(cols))
	pointers := make([]interface{}, len(cols))
	for i := range values {
		pointers[i] = &values[i]
	}
	if err := rows.Scan(pointers...); err != nil {
		panic(err)
	}

	fmt.Printf("parseToMap %v  %v\n", pointers, values)
	m := make(map[string]interface{})
	for i, colName := range cols {
		if values[i] == nil {
			m[colName] = nil
		} else {
			m[colName] = values[i]
		}
	}
	return m
}
