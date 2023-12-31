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
		err := rows.Scan(&pes.PessoaID, &pes.Nome, &pes.Iniciais,
			&pes.ImagemURL, &pes.Descricao, &pes.Titulo)

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
		err := rows.Scan(&loc.LocalID, &loc.Descricao, &loc.CidadeEstado, &loc.ImagemURL, &loc.Titulo)

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
		err := rows.Scan(&voo.VooID, &voo.Data, &voo.OrigemID,
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
		err := rows.Scan(&porto.CodigoAeroporto, &porto.LocalID,
			&porto.ImagemURL, &porto.Descricao, &porto.Titulo)

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
	rows, err := conn.Query(`SELECT Codigo, A.urlImagem, A.Descricao, A.titulo,nome,LocalID nome from 
		Aeroporto as A join Localidade on Localizacao = LocalID`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mapaPortos = make(map[string]*PaginaPorto, 20)
	for rows.Next() {
		var pp PaginaPorto
		err := rows.Scan(&pp.CodigoAeroporto, &pp.ImagemURL, &pp.Descricao,
			&pp.Titulo, &pp.Nome, &pp.LocalID)

		// n usarei reflect, mt complicado
		if err != nil {
			return nil, fmt.Errorf("TodosPaginaPortos: %v", err)
		}

		// pp.VoosDestino, err = VoosPorAeroporto(pp.CodigoAeroporto, true, conn)
		// pp.VoosOrigem, err = VoosPorAeroporto(pp.CodigoAeroporto, false, conn)

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
		err = rows.Scan(&aero.AeronaveID, &aero.NumCauda, &aero.Modelo,
			&aero.Fabricante, &aero.ImagemURL, &aero.NumAssentos)

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

// func PessoasPorLocal(LocalID int64, conn *sql.DB) ([]*Pessoa, error) {
// 	rows, err := conn.Query("SELECT PessoaID,Nome,DataNasc FROM Pessoa WHERE CidadeNascimento = ?", LocalID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	pessoas := make([]*Pessoa, 0, 10)
// 	for rows.Next() {
// 		var pes Pessoa
// 		err := rows.Scan(&pes.PessoaID, &pes.Nome, &pes.DataNascimento)
// 		if err != nil {
// 			return nil, fmt.Errorf("PessoasPorLocal %d: %v", LocalID, err)
// 		}
// 		pessoas = append(pessoas, &pes)
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, fmt.Errorf("PessoasPorLocal %d: %v", LocalID, err)
// 	}
// 	return pessoas, nil
// }

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

// ArrDep determina se você buscará por destino ou por origem: true, false
func VoosPorAeroporto(Codigo string, ArrDep bool, conn *sql.DB) ([]*Voo, error) {
	const dest = `SELECT VooID, Origem, Destino, DataVoo FROM 
	Voo WHERE Destino = ? ORDER BY DataVoo`
	const orig = `SELECT VooID, Origem, Destino, DataVoo FROM 
	Voo WHERE Origem = ? ORDER BY DataVoo`

	query := orig
	if ArrDep {
		query = dest
	}

	rows, err := conn.Query(query, Codigo)
	if err != nil {
		return nil, err
	}

	resp := make([]*Voo, 0, 20)
	for rows.Next() {
		var v Voo
		err := rows.Scan(&v.VooID, &v.OrigemID, &v.DestinoID, &v.Data)
		if err != nil {
			return nil, fmt.Errorf("VoosPorAeroporto %s: %v", Codigo, err)
		}
		resp = append(resp, &v)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("VoosPorAeroporto %s: %v", Codigo, err)
	}
	return resp, nil
}

func NumVoosPorPessoa(PessoaID int64, conn *sql.DB) (int, error) {
	rows, err := conn.Query(`SELECT count(fk_Voo) FROM 
	Embarcam WHERE fk_Pessoa = ?`, PessoaID)
	if err != nil {
		return 0, err
	}

	var numVoos int
	err = rows.Scan(&numVoos)

	if err := rows.Err(); err != nil {
		return 0, fmt.Errorf("NumVoosPorPessoa %d: %v", PessoaID, err)
	}
	return numVoos, nil
}

type numVoosNave struct {
	numVoos int
	idNave  int
}

func NumVoosPorPessoaPorNave(PessoaID int64, conn *sql.DB) ([]*numVoosNave, error) {
	rows, err := conn.Query(`SELECT AeronaveID, count(VooID) FROM 
	Embarcam JOIN Voo on fk_Voo=VooID JOIN Aeronave on Nave=AeronaveID WHERE fk_Pessoa = ? group by AeronaveID`, PessoaID)
	if err != nil {
		return nil, err
	}

	TotalVoosNave := make([]*numVoosNave, 0, 6)
	for rows.Next() {
		var nvn *numVoosNave
		err := rows.Scan(&nvn.idNave, &nvn.numVoos)
		if err != nil {
			return nil, fmt.Errorf("NumVoosPorPessoaPorNave %d: %v", PessoaID, err)
		}

		TotalVoosNave = append(TotalVoosNave, nvn)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("NumVoosPorPessoaPorNave %d: %v", PessoaID, err)
	}

	return TotalVoosNave, nil
}

func PessoasQueVisitaramAeroporto(codigo string, conn *sql.DB) ([]*Pessoa, error) {
	rows, err := conn.Query(`select PessoaID, Nome from Pessoa as P where exists 
	(select * from Embarcam join Voo on fk_Voo = VooID and fk_Pessoa = P.PessoaID 
		where Origem = '?' or Destino = '?')`, codigo, codigo)
	if err != nil {
		return nil, err
	}
	var pessoas = make([]*Pessoa, 0, 6)
	for rows.Next() {
		var pes Pessoa
		err := rows.Scan(&pes.PessoaID, &pes.Nome)
		if err != nil {
			return nil, fmt.Errorf("PessoasQueVisitaramAeroporto: %v", err)
		}
		pessoas = append(pessoas, &pes)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("NuncaVisitaramIlha: %v", err)
	}
	return pessoas, nil
}

func PessoasQueViajaramCom(pessoaID int64, conn *sql.DB) ([]*Pessoa, error) {
	rows, err := conn.Query(`select E.fk_Pessoa from 
		Embarcam as E join Embarcam as B on 
		E.fk_Voo = B.fk_Voo and B.fk_Pessoa= ? and E.fk_Pessoa != ?`, pessoaID, pessoaID)
	if err != nil {
		return nil, err
	}
	var pessoas = make([]*Pessoa, 0, 6)
	for rows.Next() {
		var pes Pessoa
		err := rows.Scan(&pes.PessoaID)
		if err != nil {
			return nil, fmt.Errorf("PessoasQueViajaramCom: %v", err)
		}
		pessoas = append(pessoas, &pes)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("PessoasQueViajaramCom: %v", err)
	}
	return pessoas, nil
}

func VoosPorComNPessoas(n int64, conn *sql.DB) ([]*Voo, error) {
	rows, err := conn.Query(`select * from Voo join Embarcam on 
	VooID = fk_Voo group by VooID having count(fk_Pessoa) > ?`, n)
	if err != nil {
		return nil, err
	}

	resp := make([]*Voo, 0, 20)
	for rows.Next() {
		var v Voo
		err := rows.Scan(&v.VooID, &v.OrigemID, &v.DestinoID, &v.Data)
		if err != nil {
			return nil, fmt.Errorf("VoosPorNPessoas %05d: %v", n, err)
		}
		resp = append(resp, &v)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("VoosPorNPessoas %05d: %v", n, err)
	}
	return resp, nil
}
