package stsql

import "fmt"

const procedure = `
CREATE OR REPLACE FUNCTION format_name
(firstname Text,lastname Text,age INT) RETURNS
VARCHAR AS $$
BEGIN
RETURN trim(firstname) ||' '||trim(lastname) ||' ('||age||')';
END;
$$ LANGUAGE plpgsql;
`
const call = "select * from format_name($1,$2,$3)"

type Result struct {
	Name     string
	Category int
}

func Procedure() {
	db := createConnnection()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	tx.Exec(procedure)
	r := Result{}

	if err := tx.QueryRow(call, "John", "Doe", 32).Scan(&r.Name); err != nil {
		panic(err)
	}

	fmt.Printf("Result is : %+v\n", r)

	tx.Commit()
}
