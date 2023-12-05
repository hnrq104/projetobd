package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"

	"github.com/go-sql-driver/mysql"
	// _ "github.com/go-sql-driver/mysql"
)

type Pessoa struct {
	PessoaID         int
	Nome             string
	Iniciais         sql.NullString
	Conhecido        sql.NullString
	DataNascimento   time.Time
	DataMorte        sql.NullTime
	BreveDescricao   sql.NullString
	CidadeNascimento sql.NullInt64
	CidadeMorte      sql.NullInt64
	URLImagem        string
}

var pessoaTemp = template.Must(template.New("PessoaTemplate").Parse(`
<h1>Pessoa</h1>
<img src="{{.URLImagem}}" alt="certamente nao eh um gato">
<table>
<tbody>
<tr>
	<th>Nome</th>
	<td>{{.Nome}}</td>
</tr>
<tr>
	<th>Iniciais</th>
	<td>{{.Iniciais}}</td>
</tr>
<tr>
	<th>Nascimento</th>
	<td>{{.CidadeNascimento}}</td>
	<td>{{.DataNascimento}}</td>
</tr>
<tr>
	<th>Morte</th>
	<td>{{.CidadeMorte}}</td>
	<td>{{.DataMorte}}</td>
</tr>
</tbody>
</table>
<p> {{.BreveDescricao}}</p>
`))

var listaPessoasTemp = template.Must(template.New("PessoasTemplate").Parse(`
<h1>Pessoas</h1>
<table>
	<tr>
		<th>Nome</th>
		<th>Nascimento</th>
	</tr>
{{range $k, $v := .Pessoas}}
	<tr>
		<td> <a href=/list?pessoaid={{$k}}>{{ $v.Nome }}</a> </td>
		<td>{{ $v.DataNascimento}}</td>
	</tr>
{{end}}
</table>
`))

type PessoasCache struct {
	Pessoas map[int]*Pessoa
	Tam     int
}

func (c PessoasCache) lista(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Problem parsing form")
		log.Print("Problem parsing form")
		return
	}

	if r.Form.Has("pessoaid") {
		id, err := strconv.Atoi(r.Form.Get("pessoaid"))
		pes, ok := c.Pessoas[id]

		if err != nil || !ok {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Problem parsing form")
			log.Print("Problem parsing form")
			return
		}
		log.Print(pessoaTemp.Execute(w, pes))
	} else {
		log.Print(listaPessoasTemp.Execute(w, c))
	}
}

var db *sql.DB

func main() {
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "epstein2",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	var err error
	// log.Println(cfg.FormatDSN())
	db, err = sql.Open("mysql", cfg.FormatDSN())
	// db, err := sql.Open("mysql", "goapp:1234@tcp(localhost:3306)/epsteindb")
	defer db.Close()
	if err != nil {
		log.Fatal("open: ", err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal("ping: ", pingErr)
	}
	fmt.Println("Connected")

	pessoas, err := ListaPessoas(db)
	if err != nil {
		log.Fatal(err)
	}

	// for _, p := range pessoas {
	// 	fmt.Println(p.Nome, p.DataNascimento)
	// }

	var cache PessoasCache = PessoasCache{Pessoas: pessoas, Tam: len(pessoas)}
	http.HandleFunc("/list", cache.lista)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func ListaPessoas(conn *sql.DB) (map[int]*Pessoa, error) {
	var pessoas = make(map[int]*Pessoa, 200)
	rows, err := conn.Query("SELECT * FROM Pessoa")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var pes Pessoa
		err := rows.Scan(&pes.PessoaID, &pes.Nome, &pes.Iniciais, &pes.Conhecido,
			&pes.DataNascimento, &pes.DataMorte, &pes.BreveDescricao,
			&pes.CidadeNascimento, &pes.CidadeMorte, &pes.URLImagem)

		if err != nil {
			return nil, fmt.Errorf("ListaPessoas: %v", err)
		}
		pessoas[pes.PessoaID] = &pes

	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ListaPessoas: %v", err)
	}

	return pessoas, nil
}
