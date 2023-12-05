package epstein

import (
	"database/sql"
	"time"
)

type Pessoa struct {
	PessoaID           int64
	Nome               string
	Iniciais           sql.NullString
	Conhecido          sql.NullString
	Descricao          sql.NullString
	ImagemURL          sql.NullString
	LocalNascimentoID  sql.NullInt64
	DataNascimento     sql.NullTime
	LocalFalecimentoID sql.NullInt64
	DataFalecimento    sql.NullTime
}

type Voo struct {
	VooID          int64
	Data           time.Time
	NumPassageiros sql.NullInt64
	AeronaveID     int64
	DestinoID      string
	OrigemID       string
}

type Aeronave struct {
	AeronaveID  int64
	NumAssentos sql.NullInt64
	// Tipo        sql.NullString
	NumCauda   sql.NullString
	Modelo     sql.NullString
	Fabricante sql.NullString
	ImagemURL  sql.NullString
}

type Aeroporto struct {
	CodigoAeroporto string
	Titulo          sql.NullString
	Descricao       sql.NullString
	ImagemURL       sql.NullString
	LocalID         int64
}

type Local struct {
	LocalID      int64
	CidadeEstado string
	Descricao    sql.NullString
	Pais         string
	ImagemURL    sql.NullString
}

type Embarca struct {
	VooID    int64
	PessoaID int64
}
