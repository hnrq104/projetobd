package main

import (
	"database/sql"
	"epstein"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
)

type Cache struct {
	db      *sql.DB
	Pessoas map[int64]*epstein.Pessoa
	Voos    map[int64]*epstein.Voo
	Portos  map[string]*epstein.PaginaPorto
	Locais  map[int64]*epstein.Local
	Naves   map[int64]*epstein.Aeronave
}

func (c Cache) listaPessoas(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Problem parsing form")
		log.Print("Problem parsing form")
		return
	}

	if r.Form.Has("pessoaid") {
		id, err := strconv.ParseInt(r.Form.Get("pessoaid"), 10, 64)
		pes, ok := c.Pessoas[id]
		if err != nil || !ok {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Problem parsing form")
			log.Print("Problem parsing form")
			return
		}
		var pp epstein.PessoaPagina

		pp.Pessoa = pes
		pp.VoosFeitos, err = epstein.VoosPorPessoa(id, c.db)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Problem retrieving flights %v", err)
			log.Print("Problem retrieving flights", err)
			return
		}

		log.Print(epstein.PessoaTemp.Execute(w, pp))
	} else {
		log.Print(epstein.MapPessoasTemp.Execute(w, c.Pessoas))
	}
}

func (c Cache) listaVoos(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Problem parsing form")
		log.Print("Problem parsing form")
		return
	}

	if r.Form.Has("vooid") {
		id, err := strconv.ParseInt(r.Form.Get("vooid"), 10, 64)
		voo, ok := c.Voos[id]
		if err != nil || !ok {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Problem parsing form")
			log.Print("Problem parsing form")
			return
		}
		var vp epstein.VooPagina

		vp.Voo = voo
		vp.Passageiros, err = epstein.PassageirosPorVoo(id, c.db)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Problem retrieving flights %v", err)
			log.Print("Problem retrieving flights", err)
			return
		}

		log.Print(epstein.VooTemp.Execute(w, vp))
	} else {
		log.Print(epstein.MapVoosTemp.Execute(w, c.Voos))
	}
}

func (c Cache) listaPortos(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Problem parsing form")
		log.Print("Problem parsing form")
		return
	}

	if r.Form.Has("codigo") {
		id := r.Form.Get("codigo")
		pp, ok := c.Portos[id]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Problem parsing form")
			log.Print("Problem parsing form")
			return
		}

		log.Print(epstein.PaginaPortoTemp.Execute(w, pp))
	} else {
		log.Print(epstein.MapAeroportosTemp.Execute(w, c.Portos))
	}
}

func (c Cache) listaLocais(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Problem parsing form")
		log.Print("Problem parsing form")
		return
	}

	if r.Form.Has("localid") {
		id, err := strconv.ParseInt(r.Form.Get("localid"), 10, 64)
		loc, ok := c.Locais[id]
		if err != nil || !ok {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Problem parsing form")
			log.Print("Problem parsing form")
			return
		}
		var pagLoc epstein.PaginaLocal
		pagLoc.Local = loc

		pagLoc.PessoasNascidas, err = epstein.PessoasPorLocal(loc.LocalID, c.db)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Problema Procurando Pessoas")
			log.Print("Problema procurando pessoas")
		}

		pagLoc.Aeroportos, err = epstein.AeroportosPorLocal(loc.LocalID, c.db)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Problema Procurando Aeroportos")
			log.Print("Problema procurando Aeroportos")
		}

		log.Print(epstein.LocalTemp.Execute(w, pagLoc))
	} else {
		log.Print(epstein.MapLocaisTemp.Execute(w, c.Locais))
	}
}

func (c Cache) listaNaves(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Problem parsing form")
		log.Print("Problem parsing form")
		return
	}

	if r.Form.Has("naveid") {
		id, err := strconv.ParseInt(r.Form.Get("naveid"), 10, 64)
		nave, ok := c.Naves[id]
		if err != nil || !ok {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Problem parsing form")
			log.Print("Problem parsing form")
			return
		}

		var np epstein.AeronavePagina
		np.Aeronave = nave

		np.VoosFeitos, err = epstein.VoosPorNave(nave.AeronaveID, c.db)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "problem retrieving Flights")
			log.Print("Problem retrieving flights")
			return
		}

		log.Print(epstein.AeronaveTemp.Execute(w, np))
	} else {
		log.Print(epstein.MapAeronaveTemp.Execute(w, c.Naves))
	}
}

func main() {
	t := time.Now()
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "epstein2",
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

	var c Cache
	c.db = db
	pessoas, err := epstein.TodasPessoas(c.db)

	if err != nil {
		log.Fatal("TodasPessoas: ", err)
	}

	voos, err := epstein.TodosVoos(c.db)
	if err != nil {
		log.Fatal("TodosVoos: ", err)
	}

	portos, err := epstein.TodosPaginaPortos(c.db)
	if err != nil {
		log.Fatal("TodosPaginaPortos: ", err)
	}

	locais, err := epstein.TodosLocais(c.db)
	if err != nil {
		log.Fatal("TodosLocais: ", err)
	}

	naves, err := epstein.TodasAeronaves(c.db)
	if err != nil {
		log.Fatal("TodasAeronaves: ", err)
	}

	c.Portos = portos
	c.Pessoas = pessoas
	c.Voos = voos
	c.Locais = locais
	c.Naves = naves

	fmt.Println(time.Since(t).Milliseconds(), "ms")

	http.HandleFunc("/pessoas", c.listaPessoas)
	http.HandleFunc("/voos", c.listaVoos)
	http.HandleFunc("/aeroportos", c.listaPortos)
	http.HandleFunc("/locais", c.listaLocais)
	http.HandleFunc("/naves", c.listaNaves)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	epstein.HomeTemplate.Execute(w, nil)
}
