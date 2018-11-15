package stsql

import "fmt"

const insert4 = "INSERT INTO post(ID,TITLE,CONTENT) VALUES(4,'Transaction Title', 'Transaction Content'); "

type POST struct {
	ID      int
	Title   string
	Content string
}

func Transaction() {
	db := createConnnection()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	_, err = tx.Exec(insert4)
	if err != nil {
		panic(err)
	}

	p := POST{}
	rs := tx.QueryRow(selOne, 4)

	err = rs.Scan(&p.Title, &p.Content)

	if err != nil {
		fmt.Println("Got error for db.Query:" + err.Error())
	}

	fmt.Println(p)
	tx.Rollback()

}
