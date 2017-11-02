package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	/* abrindo o arquivo CSV */
	f, err := os.Open("iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	reader := csv.NewReader(f)

	reader.FieldsPerRecord = -1

	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	/* formato [][]string */
	fmt.Println(rawCSVData)
}
