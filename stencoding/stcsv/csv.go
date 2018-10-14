package stcsv

import (
	"encoding/csv"
	"fmt"
	"os"
)

func CSV() {
	f, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 3
	reader.Comment = '#'
	for {
		record, e := reader.Read()
		if e != nil {
			fmt.Println(e)
			break
		}
		fmt.Println(record)
	}
}
