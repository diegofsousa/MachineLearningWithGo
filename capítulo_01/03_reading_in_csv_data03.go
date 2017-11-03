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
	// serão lidos apenas 5 campos por linha
	reader.FieldsPerRecord = 5

	var rawCSVData [][]string
	// Lendo os registros um por um
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// Se houver erros de análise, registrar o erro e seguir em frente.
		if err != nil {
			log.Println(err)
			continue
		}

		rawCSVData = append(rawCSVData, record)
	}
	fmt.Println(rawCSVData)
}
