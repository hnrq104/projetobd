package epstein

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func TodasPessoas(conn *sql.DB) (map[int64]*Pessoa, error) {
	rows, err := conn.Query("SELECT * FROM Pessoa")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pessoas = make(map[int64]*Pessoa, 200)
	for rows.Next() {
		var pes Pessoa
		err := rows.Scan(&pes.PessoaID, &pes.Nome, &pes.Iniciais, &pes.Conhecido,
			&pes.DataNascimento, &pes.DataFalecimento, &pes.Descricao,
			&pes.LocalNascimentoID, &pes.LocalFalecimentoID, &pes.ImagemURL)

		// n usarei reflect, mt complicado
		if err != nil {
			return nil, fmt.Errorf("TodasPessoas: %v", err)
		}
		pessoas[pes.PessoaID] = &pes
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("TodasPessoas: %v", err)
	}
	return pessoas, nil
}

func TodosLocais(conn *sql.DB) (map[int64]*Local, error) {
	rows, err := conn.Query("SELECT * FROM Localidade")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var locais = make(map[int64]*Local, 200)
	for rows.Next() {
		var loc Local
		err := rows.Scan(&loc.LocalID, &loc.Descricao, &loc.CidadeEstado, &loc.Pais, &loc.ImagemURL)

		// n usarei reflect, mt complicado
		if err != nil {
			return nil, fmt.Errorf("TodosLocais: %v", err)
		}
		locais[loc.LocalID] = &loc
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("TodosLocais: %v", err)
	}
	return locais, nil
}

func TodosVoos(conn *sql.DB) (map[int64]*Voo, error) {
	rows, err := conn.Query("SELECT * FROM Voo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var MapaVoos = make(map[int64]*Voo, 200)
	for rows.Next() {
		var voo Voo
		err := rows.Scan(&voo.VooID, &voo.NumPassageiros, &voo.Data, &voo.OrigemID,
			&voo.DestinoID, &voo.AeronaveID)

		// n usarei reflect, mt complicado
		if err != nil {
			return nil, fmt.Errorf("TodosVoos: %v", err)
		}
		MapaVoos[voo.VooID] = &voo
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("TodosVoos: %v", err)
	}
	return MapaVoos, nil
}

func TodosAeroportos(conn *sql.DB) (map[string]*Aeroporto, error) {
	rows, err := conn.Query("SELECT * FROM Aeroporto")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var aeroportos = make(map[string]*Aeroporto, 40)
	for rows.Next() {
		var porto Aeroporto
		err := rows.Scan(&porto.CodigoAeroporto, &porto.Descricao, &porto.LocalID, &porto.ImagemURL)

		// n usarei reflect, mt complicado
		if err != nil {
			return nil, fmt.Errorf("TodosAeroportos: %v", err)
		}
		aeroportos[porto.CodigoAeroporto] = &porto
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("TodosAeroportos: %v", err)
	}
	return aeroportos, nil
}

func VoosPorPessoa(IDPessoa int64, conn *sql.DB) ([]*Voo, error) {
	rows, err := conn.Query(`SELECT VooID, Origem, Destino, DataVoo FROM 
	Voo JOIN Embarcam ON VooID = fk_Voo WHERE fk_Pessoa = ? ORDER BY DataVoo`, IDPessoa)
	if err != nil {
		return nil, err
	}

	resp := make([]*Voo, 0, 20)
	for rows.Next() {
		var v Voo
		err := rows.Scan(&v.VooID, &v.OrigemID, &v.DestinoID, &v.Data)
		if err != nil {
			return nil, fmt.Errorf("VoosPorPessoa %05d: %v", IDPessoa, err)
		}
		resp = append(resp, &v)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("VoosPorPessoa %05d: %v", IDPessoa, err)
	}
	return resp, nil
}

func PassageirosPorVoo(IDVoo int64, conn *sql.DB) ([]*Pessoa, error) {
	rows, err := conn.Query(`SELECT PessoaID, Nome FROM 
	Pessoa JOIN Embarcam ON PessoaID = fk_Pessoa WHERE fk_Voo = ?`, IDVoo)
	if err != nil {
		return nil, err
	}

	resp := make([]*Pessoa, 0, 6)
	for rows.Next() {
		var p Pessoa
		err := rows.Scan(&p.PessoaID, &p.Nome)
		if err != nil {
			return nil, fmt.Errorf("PassageirosPorVoo %05d: %v", IDVoo, err)
		}
		resp = append(resp, &p)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("PassageirosPorVoo %05d: %v", IDVoo, err)
	}
	return resp, nil
}

func TodosPaginaPortos(conn *sql.DB) (map[string]*PaginaPorto, error) {
	rows, err := conn.Query(`SELECT Codigo, LocalID, CidadeEstado, NomePais, A.Descricao, A.urlImagem FROM 
		Aeroporto as A JOIN Localidade ON Localizacao = LocalID`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mapaPortos = make(map[string]*PaginaPorto, 20)
	for rows.Next() {
		var pp PaginaPorto
		err := rows.Scan(&pp.CodigoAeroporto, &pp.LocalID, &pp.CEP, &pp.Pais, &pp.Descricao, &pp.ImagemURL)

		// n usarei reflect, mt complicado
		if err != nil {
			return nil, fmt.Errorf("TodosPaginaPortos: %v", err)
		}
		mapaPortos[pp.CodigoAeroporto] = &pp
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("TodosPaginaPortos: %v", err)
	}
	return mapaPortos, nil
}

func TodasAeronaves(conn *sql.DB) (map[int64]*Aeronave, error) {
	rows, err := conn.Query("SELECT * FROM Aeronave")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var naves = make(map[int64]*Aeronave, 10)
	for rows.Next() {
		var aero Aeronave
		err = rows.Scan(&aero.AeronaveID, &aero.NumAssentos, &aero.NumCauda,
			&aero.Modelo, &aero.Fabricante, &aero.ImagemURL)

		if err != nil {
			return nil, fmt.Errorf("TodasAeronaves: %v", err)
		}
		naves[aero.AeronaveID] = &aero
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("TodasAeronaves: %v", err)
	}

	return naves, nil
}

func VoosPorNave(AeronaveID int64, conn *sql.DB) ([]*Voo, error) {
	rows, err := conn.Query(`SELECT VooID, Origem, Destino, DataVoo FROM 
	Voo WHERE Nave = ? ORDER BY DataVoo`, AeronaveID)
	if err != nil {
		return nil, err
	}

	resp := make([]*Voo, 0, 20)
	for rows.Next() {
		var v Voo
		err := rows.Scan(&v.VooID, &v.OrigemID, &v.DestinoID, &v.Data)
		if err != nil {
			return nil, fmt.Errorf("VoosPorNave %05d: %v", AeronaveID, err)
		}
		resp = append(resp, &v)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("VoosPorPessoa %05d: %v", AeronaveID, err)
	}
	return resp, nil
}

func PessoasPorLocal(LocalID int64, conn *sql.DB) ([]*Pessoa, error) {
	rows, err := conn.Query("SELECT PessoaID,Nome,DataNasc FROM Pessoa WHERE CidadeNascimento = ?", LocalID)
	if err != nil {
		return nil, err
	}

	pessoas := make([]*Pessoa, 0, 10)
	for rows.Next() {
		var pes Pessoa
		err := rows.Scan(&pes.PessoaID, &pes.Nome, &pes.DataNascimento)
		if err != nil {
			return nil, fmt.Errorf("PessoasPorLocal %d: %v", LocalID, err)
		}
		pessoas = append(pessoas, &pes)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("PessoasPorLocal %d: %v", LocalID, err)
	}
	return pessoas, nil
}

func AeroportosPorLocal(LocalID int64, conn *sql.DB) ([]*Aeroporto, error) {
	rows, err := conn.Query("SELECT Codigo FROM Aeroporto WHERE Localizacao = ?", LocalID)
	if err != nil {
		return nil, err
	}

	aeroportos := make([]*Aeroporto, 0, 10)
	for rows.Next() {
		var porto Aeroporto
		err := rows.Scan(&porto.CodigoAeroporto)
		if err != nil {
			return nil, fmt.Errorf("AeroportoPorLocal %d: %v", LocalID, err)
		}
		aeroportos = append(aeroportos, &porto)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("AeroportoPorLocal %d: %v", LocalID, err)
	}
	return aeroportos, nil
}
