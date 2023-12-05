package main

import (
	"database/sql"
	"log"
)

func inserirPessoasTabela(db *sql.DB, pessoas map[string]*Pessoa) error {
	const pesstring = "INSERT INTO Pessoa(PessoaID, Nome, titulo, Iniciais, urlImagem, Descricao) VALUES(?,?,?,?,?,?)"
	for _, v := range pessoas {
		_, err := db.Exec(pesstring, v.id, v.nome, v.titulo, v.iniciais, v.URLimage, v.descricao)
		if err != nil {
			log.Print(err)
			return err
		}
	}
	return nil
}

func inserirLocaisTabela(db *sql.DB, locais map[string]*Local) error {
	const localstring = "INSERT INTO Localidade(LocalID, Nome, titulo, Descricao, urlImagem) VALUES(?,?,?,?,?)"
	for _, v := range locais {
		_, err := db.Exec(localstring, v.id, v.nome, v.titulo, v.descricao, v.URLimage)
		if err != nil {
			log.Print(err)
			return err
		}
	}
	return nil
}

func inserirAvioesTabela(db *sql.DB, avioes map[string]*Aviao) error {
	const aviaostring = "INSERT INTO Aeronave(AeronaveID, NumCauda, NumAssentos, Modelo) VALUES(?,?,?,?)"
	for _, v := range avioes {
		_, err := db.Exec(aviaostring, v.id, v.cauda, v.assentos, v.modelo)
		if err != nil {
			log.Print(err)
			return err
		}
	}
	return nil
}

func inserirAeroportoTabela(db *sql.DB, portos map[string]*Aeroporto) error {
	const aerostring = "INSERT INTO Aeroporto(Codigo, Localizacao, urlImagem, Descricao, titulo) VALUES(?,?,?,?,?)"
	for _, v := range portos {
		_, err := db.Exec(aerostring, v.codigo, v.localid, v.URLimage, v.descricao, v.titulo)
		if err != nil {
			log.Print(err)
			return err
		}
	}
	return nil
}

// vai inserir voo e embarcam
func inserirVooTabela(db *sql.DB, voos map[string]*Voo) error {
	const voostring = "INSERT INTO Voo(VooID,DataVoo,Origem,Destino,Nave) VALUES(?,?,?,?,?)"

	const embarcastring = "INSERT INTO Embarcam(fk_Voo,fk_Pessoa) VALUES(?,?)"

	for _, v := range voos {
		_, err := db.Exec(voostring, v.flightID, v.date.Format("2006-01-02"), v.aeroDEP.codigo, v.aeroARR.codigo, v.aviao.id)
		if err != nil {
			log.Print(err)
			return err
		}

		for _, p := range v.passageiros {
			_, err := db.Exec(embarcastring, v.flightID, p.id)
			if err != nil {
				log.Print(err)
				return err
			}

		}

	}
	return nil
}

func deletaPessoas(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM Pessoa")
	return err
}

func deletaVoos(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM Voo")
	return err
}

func deletaAeroporto(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM Aeroporto")
	return err
}

func deletaAeronave(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM Aeronave")
	return err
}

func deletaLocal(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM Localidade")
	return err
}

func deletaEmbarca(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM Embarcam")
	return err
}
