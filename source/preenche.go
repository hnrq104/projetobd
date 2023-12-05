package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"sync"
	"time"
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

func initCache() *Cache {
	return &Cache{
		make(map[string]*Pessoa),
		make(map[string]*Local),
		make(map[string]*Voo),
		make(map[string]*Aeroporto),
		make(map[string]*Aviao),
	}
}

func (c *Cache) line(record []string) {
	flight, ok := c.voos[record[11]]
	if !ok {
		//Lê partes
		dep, ok := c.locais[record[9]]
		if !ok {
			dep = new(Local)
			dep.nome = record[9]
			c.locais[record[9]] = dep
			dep.id = len(c.locais)
		}

		arr, ok := c.locais[record[10]]
		if !ok {
			arr = new(Local)
			arr.nome = record[10]
			c.locais[record[10]] = arr
			arr.id = len(c.locais)
		}

		aerodep, ok := c.aeroportos[record[7]]
		if !ok {
			aerodep = new(Aeroporto)
			aerodep.codigo = record[7]
			aerodep.localid = dep.id
			c.aeroportos[record[7]] = aerodep
		}

		aeroarr, ok := c.aeroportos[record[8]]
		if !ok {
			aeroarr = new(Aeroporto)
			aeroarr.codigo = record[8]
			aeroarr.localid = arr.id
			c.aeroportos[record[8]] = aeroarr
		}

		aviao, ok := c.avioes[record[4]]
		if !ok {
			aviao = new(Aviao)
			aviao.cauda = record[4]
			aviao.modelo = record[3]
			aviao.assentos, _ = strconv.Atoi(record[6])
			c.avioes[record[4]] = aviao
			aviao.id = len(c.avioes)
		}

		flight = new(Voo)
		var err error
		flight.flightID, err = strconv.Atoi(record[11])
		if err != nil {
			//não adiciona flight
			// log.Print(err)
			return
		}

		flight.date, err = time.Parse("1/2/2006", record[1])
		if err != nil {
			log.Fatal(err)
		}

		flight.aeroDEP = aerodep
		flight.aeroARR = aeroarr
		flight.passageiros = make([]*Pessoa, 0)
		c.voos[record[11]] = flight
	}

	pes, ok := c.pessoas[record[17]]
	if !ok {
		pes = new(Pessoa)
		pes.nome = record[17]
		pes.iniciais = record[19]
		c.pessoas[record[17]] = pes
		pes.id = len(c.pessoas)
	}
	flight.passageiros = append(flight.passageiros, pes)
}

func main() {
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

	// fmt.Println("Aeroportos")
	// for k, v := range c.aeroportos {
	// 	fmt.Println(k, v)
	// }
	// fmt.Println("Avioes")
	// for k, v := range c.avioes {
	// 	fmt.Println(k, v)
	// }
	// fmt.Println("Locais")
	// for k, v := range c.locais {
	// 	fmt.Println(k, v)
	// }
	// fmt.Println("Pessoas")
	// for k, v := range c.pessoas {
	// 	fmt.Println(k, v)
	// }

	// fmt.Println("Voos")
	// for k, v := range c.voos {
	// 	fmt.Println(k, v)
	// }

	// fmt.Println(len(c.pessoas))
	// fmt.Println(len(c.aeroportos))
	// fmt.Println(len(c.avioes))
	// fmt.Println(len(c.locais))
	// fmt.Println(len(c.voos))

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

	for _, v := range c.locais {
		fmt.Println(v)
	}
	for _, v := range c.aeroportos {
		fmt.Println(v)
	}
	for _, v := range c.pessoas {
		fmt.Println(v)
	}
}

const wikimediaSearch = "https://en.wikipedia.org/w/rest.php/v1/search/page?limit=1&q="
const wikisummary = "https://en.wikipedia.org/api/rest_v1/page/summary/"

type WikiSumJSON struct {
	Title     string
	Extract   string
	Thumbnail struct {
		Source string
	}
}

type WikiSearchJSON struct {
	Key string
}

type WikiSearchResult struct {
	Pages []*WikiSearchJSON
}

// Cada função procura fará 2 requests caso o primeiro seja bem sucedido
func (p *Pessoa) ProcuraWikipedia() (*WikiSearchResult, error) {
	escapedName := url.QueryEscape(p.nome)
	resp, err := http.Get(wikimediaSearch + escapedName)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result WikiSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("Procurando %s: %v", p.nome, err)
	}

	if len(result.Pages) == 0 {
		return nil, fmt.Errorf("Not Found")
	}
	return &result, err
}

func (p *Pessoa) SetaSumario(titulo *WikiSearchResult) error {
	resp, err := http.Get(wikisummary + titulo.Pages[0].Key)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var summary WikiSumJSON
	if err := json.NewDecoder(resp.Body).Decode(&summary); err != nil {
		return fmt.Errorf("Sumarizando %s: %v", p.nome, err)
	}

	p.URLimage = summary.Thumbnail.Source
	p.descricao = summary.Extract
	p.titulo = summary.Title
	return nil
}

func (loc *Local) ProcuraWikipedia() (*WikiSearchResult, error) {
	escapedName := url.QueryEscape(loc.nome)
	resp, err := http.Get(wikimediaSearch + escapedName)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result WikiSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("Procurando %s: %v", loc.nome, err)
	}

	if len(result.Pages) == 0 {
		return nil, fmt.Errorf("Not Found")
	}
	return &result, err
}

func (loc *Local) SetaSumario(titulo *WikiSearchResult) error {
	resp, err := http.Get(wikisummary + titulo.Pages[0].Key)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var summary WikiSumJSON
	if err := json.NewDecoder(resp.Body).Decode(&summary); err != nil {
		return fmt.Errorf("Sumarizando %s: %v", loc.nome, err)
	}

	loc.titulo = summary.Title
	loc.URLimage = summary.Thumbnail.Source
	loc.descricao = summary.Extract
	return nil
}

func (aero *Aeroporto) ProcuraWikipedia() (*WikiSearchResult, error) {
	escapedName := url.QueryEscape(aero.codigo + "_airport")
	resp, err := http.Get(wikimediaSearch + escapedName)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result WikiSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("Procurando %s: %v", aero.codigo, err)
	}

	if len(result.Pages) == 0 {
		return nil, fmt.Errorf("Not Found")
	}
	return &result, err
}

func (aero *Aeroporto) SetaSumario(titulo *WikiSearchResult) error {
	resp, err := http.Get(wikisummary + titulo.Pages[0].Key)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var summary WikiSumJSON
	if err := json.NewDecoder(resp.Body).Decode(&summary); err != nil {
		return fmt.Errorf("Sumarizando %s: %v", aero.codigo, err)
	}

	aero.URLimage = summary.Thumbnail.Source
	aero.descricao = summary.Extract
	aero.titulo = summary.Title
	return nil
}
