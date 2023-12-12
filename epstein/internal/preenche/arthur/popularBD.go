package main

// Programa do Arthur
import (
	"database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func LerCSV(db *sql.DB) {
	f, err := os.Open("teste_flight_logs.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csv := csv.NewReader(f)

	for {
		data, err := csv.Read()

		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {

			log.Fatal("Reading CSV", err)

		}
		if len(data) > 4 && data[0] != "ID" { //Valores Válidos
			id_localDep := inserirLocalidade(db, data[9])
			id_localArr := inserirLocalidade(db, data[10])

			id_aeroportoDep := inserirAeroporto(db, id_localDep, data[7])
			id_aeroportoArr := inserirAeroporto(db, id_localArr, data[8])

			id_aeronave := inserirAeronave(db, data[3:7])

			id_voo := inserirVoo(db, id_aeroportoDep, id_aeroportoArr, id_aeronave, data[0:2])

			id_pessoa := inserirPessoa(db, data[17], data[19], data[20])

			inserirEmbarcam(db, id_pessoa, id_voo)

		}

	}

	fmt.Println("Fim Leitura")

}

func inserirLocalidade(db *sql.DB, data string) int {

	var cidadeEstado, pais string
	split := strings.Split(data, ",")

	if len(split) >= 3 {
		cidadeEstado = strings.TrimSpace(split[0]) + "," + strings.TrimSpace(split[1])
		pais = split[2]
	} else if len(split) == 2 {
		cidadeEstado = strings.TrimSpace(split[0])
		pais = strings.TrimSpace(split[1])
	} else {
		cidadeEstado = strings.TrimSpace(split[0])
		pais = cidadeEstado
	}

	var pk int

	err := db.QueryRow(`SELECT LocalID FROM Localidade WHERE CidadeEstado = ?`, cidadeEstado).Scan(&pk)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			q := `INSERT INTO Localidade(CidadeEstado, NomePais) values (?, ?)`
			db.Exec(q, cidadeEstado, pais)

			err = db.QueryRow(`SELECT LocalID FROM Localidade WHERE CidadeEstado = ?`, cidadeEstado).Scan(&pk)
			if err != nil {
				log.Fatal("Localidade ", err)
			}

		} else {
			log.Fatal("Localidade", err)
		}
	}

	return pk

}

func inserirAeroporto(db *sql.DB, idLocal int, data string) string {
	var pk string
	err := db.QueryRow(`SELECT Codigo FROM Aeroporto WHERE Codigo = ?`, data).Scan(&pk)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			q := `INSERT INTO Aeroporto(Codigo, Localizacao) values (?, ?)`
			db.Exec(q, data, idLocal)

			err = db.QueryRow(`SELECT Codigo FROM Aeroporto WHERE Codigo = ?`, data).Scan(&pk)
			if err != nil {
				fmt.Println(idLocal)
				log.Fatal("Aeroporto", err)
			}
		} else {
			log.Fatal("Aeroporto", err)
		}
	}

	return pk
}

func inserirAeronave(db *sql.DB, data []string) string {
	var pk string
	err := db.QueryRow(`SELECT NumCauda FROM Aeronave WHERE NumCauda = ?`, data[1]).Scan(&pk)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			q := `INSERT INTO Aeronave(Modelo, NumCauda, NumDeAssentos) values (?, ?, ?)`
			db.Exec(q, data[0], data[1], data[3])

			err = db.QueryRow(`SELECT NumCauda FROM Aeronave WHERE NumCauda = ?`, data[1]).Scan(&pk)
			if err != nil {
				log.Fatal("Aeronave1", err)
			}
		} else {
			log.Fatal("Aeronave", err)
		}
	}

	return pk
}

func inserirVoo(db *sql.DB, id_Dep string, id_Arr string, id_aeronave string, data []string) int {
	var pk int

	dateLayouts := []string{
		"1/2/2006",
		"01/2/2006",
		"1/02/2006",
		"01/02/2006",
	}

	id, err := strconv.Atoi(data[0])
	if err != nil {
		log.Fatal("String to int", err)
	}

	var date time.Time
	for _, layout := range dateLayouts {
		date, err = time.Parse(layout, data[1])
		if err == nil {
			//Parse correto
			break
		}
	}

	if err != nil {
		log.Fatal("Date Parsing ", data[1], err)
	}

	err = db.QueryRow(`SELECT VooID FROM Voo WHERE VooID = ?`, id).Scan(&pk)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			q := `INSERT INTO Voo(VooID, DataVoo, Origem, Destino, Nave) values (?, ?, ?,?, ?)`
			db.Exec(q, id, date.Format("2006-01-02"), id_Dep, id_Arr, id_aeronave)

			err = db.QueryRow(`SELECT VooID FROM Voo WHERE VooID = ?`, id).Scan(&pk)
			if err != nil {
				log.Fatal("Voo", err)
			}
		} else {
			log.Fatal("Voo", err)
		}
	}
	return pk
}

func inserirPessoa(db *sql.DB, nome string, iniciais string, conhecido string) int {
	var pk int

	err := db.QueryRow(`SELECT PessoaID FROM Pessoa WHERE Nome = ?`, nome).Scan(&pk)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c := (conhecido == "Yes")
			q := `INSERT INTO Pessoa(Nome, Iniciais, Conhecido) values (?, ?, ?)`
			db.Exec(q, nome, iniciais, c)

			err = db.QueryRow(`SELECT PessoaID FROM Pessoa WHERE Nome = ?`, nome).Scan(&pk)
			if err != nil {
				log.Fatal("Pessoa", err)
			}

		} else {
			log.Fatal("Pessoa", err)
		}
	}

	return pk
}

func inserirEmbarcam(db *sql.DB, id_Pessoa int, id_Voo int) (int, int) {
	var pk1, pk2 int

	err := db.QueryRow(`SELECT fk_Pessoa, fk_Voo FROM Embarcam WHERE fk_Pessoa = ? and fk_Voo = ?`, id_Pessoa, id_Voo).Scan(&pk1, &pk2)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			db.Exec(`INSERT INTO Embarcam(fk_Pessoa, fk_Voo) values (?, ?)`, id_Pessoa, id_Voo)

			err = db.QueryRow(`SELECT fk_Pessoa, fk_Voo FROM Embarcam WHERE fk_Pessoa = ? and fk_Voo = ?`, id_Pessoa, id_Voo).Scan(&pk1, &pk2)
			if err != nil {
				log.Fatal("Embarcam", err)
			}

		} else {
			log.Fatal("Embarcam", err)
		}
	}

	return pk1, pk2
}
