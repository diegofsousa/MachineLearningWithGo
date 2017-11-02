package main

import (
	"encoding/csv"
	"fmt"
	"io"
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

	var rawCSVData [][]string
	// Lendo os registros um por um
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		rawCSVData = append(rawCSVData, record)
	}
	fmt.Println(rawCSVData)
}
