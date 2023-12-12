package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/go-sql-driver/mysql"
)

type Pessoa struct {
	id        int
	nome      string
	titulo    string
	iniciais  string
	URLimage  string
	descricao string
}

type Local struct {
	id        int
	titulo    string
	nome      string
	URLimage  string
	descricao string
	// pais      string
}

type Voo struct {
	flightID    int
	passageiros []*Pessoa
	aeroDEP     *Aeroporto
	aeroARR     *Aeroporto
	aviao       *Aviao
	date        time.Time
}

type Aeroporto struct {
	codigo    string
	localid   int
	URLimage  string
	descricao string
	titulo    string
}

type Aviao struct {
	id       int
	cauda    string
	modelo   string
	assentos int
}

// ID,Date,Year,Aircraft Model,Aircraft Tail #,Aircraft Type,# of Seats,DEP: Code,ARR: Code,DEP,ARR,Flight_No.,Pass #,Unique ID,First Name,Last Name,"Last, First",First Last,Comment,Initials,Known,Data Source

type Cache struct {
	pessoas    map[string]*Pessoa
	locais     map[string]*Local
	voos       map[string]*Voo
	aeroportos map[string]*Aeroporto
	avioes     map[string]*Aviao
}

func main() {
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "epfinal",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	defer db.Close()
	if err != nil {
		log.Fatal("open: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("ping: ", err)
	}
	fmt.Println("Connected")

	f, err := os.Open("flight_logs.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Read()

	c := initCache()

	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		c.line(rec)

		// fmt.Println(rec)
	}

	var wg sync.WaitGroup

	counter := 0
	for _, v := range c.pessoas {
		wg.Add(1)
		go func(p *Pessoa) {
			defer wg.Done()
			res, err := p.ProcuraWikipedia()
			if err != nil {
				return
			}
			err = p.SetaSumario(res)
			if err != nil {
				//do something
				return
			}

		}(v)
		counter++
		if counter%10 == 0 {
			time.Sleep(500 * time.Millisecond)
		}
	}

	counter = 0
	for _, v := range c.locais {
		wg.Add(1)
		go func(loc *Local) {
			defer wg.Done()
			res, err := loc.ProcuraWikipedia()
			if err != nil {
				return
			}
			err = loc.SetaSumario(res)
			if err != nil {
				//do something
				return
			}

		}(v)
		counter++
		if counter%10 == 0 {
			time.Sleep(500 * time.Millisecond)
		}
	}

	counter = 0
	for _, v := range c.aeroportos {
		wg.Add(1)
		go func(aero *Aeroporto) {
			defer wg.Done()
			res, err := aero.ProcuraWikipedia()
			if err != nil {
				return
			}
			err = aero.SetaSumario(res)
			if err != nil {
				//do something
				return
			}

		}(v)
		counter++
		if counter%20 == 0 {
			time.Sleep(500 * time.Millisecond)
		}
	}

	wg.Wait()

	fmt.Println("deletando valores da tabela")
	deletaAeronave(db)
	deletaAeroporto(db)
	deletaEmbarca(db)
	deletaLocal(db)
	deletaPessoas(db)
	deletaVoos(db)

	fmt.Println(inserirPessoasTabela(db, c.pessoas))
	fmt.Println(inserirLocaisTabela(db, c.locais))
	fmt.Println(inserirAvioesTabela(db, c.avioes))
	fmt.Println(inserirAeroportoTabela(db, c.aeroportos))
	fmt.Println(inserirVooTabela(db, c.voos))
}
