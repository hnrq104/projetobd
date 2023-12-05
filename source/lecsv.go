package main

import (
	"log"
	"strconv"
	"time"
)

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
		flight.aviao = aviao
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
