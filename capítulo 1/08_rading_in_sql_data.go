package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	//"os"
)

func main() {
	db, err := sql.Open("mysql", "root:dieguim961100@/serralheria_sousa")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	imprime(db)

	stmt, err := db.Prepare("INSERT serralheria_sousa.cliente SET nome=?,telefone=?,cpf=?,sexo=?,atividade=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec("LEO SILVA", "89994242321", 85856394729, "M", 1)
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	fmt.Println("Novo elemento adicionado no ID ", id)
	imprime(db)
}

func imprime(db *sql.DB) {
	rows, err := db.Query(`SELECT idcliente, nome, telefone FROM serralheria_sousa.cliente;`)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			idcliente int
			nome      string
			telefone  string
		)

		if err := rows.Scan(&idcliente, &nome, &telefone); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d, %s, %s\n", idcliente, nome, telefone)
	}
}
