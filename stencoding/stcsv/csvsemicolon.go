package stcsv

import (
	"encoding/csv"
	"fmt"
	"os"
)

func CSVSemiColon() {
	f, err := os.Open("data_uncommon.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Comma = ';'
	for {
		record, err := reader.Read()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(record)
	}
}
