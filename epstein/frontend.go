package epstein

import (
	"text/template"
)

type PessoaPagina struct {
	*Pessoa
	VoosFeitos []*Voo
}

// Recebe uma PessoaPagina
// ADICIONAR LINK PARA CADA VOO
// ADICIONAR LINK PARA TODOS OS VOOS
// ADICIONAR LINK PARA CADA AEROPORTO
// ADICIONAR LINK PARA TODOS AEROPORTOS
// TROCAR IMAGEM DEPOIS
var PessoaTemp = template.Must(template.New("PessoaTemplate").Parse(`
<a href=/>home</a>

<h1><a href=pessoas>Pessoa</a> {{.Titulo}}</h1>
<img src="{{.ImagemURL.Value}}" alt="certamente nao eh um gato">
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
</tbody>
</table>
<h3>Descricao</h3>
<p>{{.Descricao}}</p>
<h3>Voos</h3>
<table>
	<tr>
		<th><a href=voos>NumVoo</a></th>
		<th>Origem</th>
		<th>Destino</th>
	</tr>
{{range .VoosFeitos}}
	<tr>
		<td><a href=voos?vooid={{.VooID}}>{{.VooID}}</a></td>
		<td><a href=aeroportos?codigo={{.OrigemID}}>{{.OrigemID}}</a></td>
		<td><a href=aeroportos?codigo={{.DestinoID}}>{{.DestinoID}}</a></td>
	<tr>
{{end}}
</table>
`))

// Recebe um Mapa de Pessoas
var MapPessoasTemp = template.Must(template.New("PessoasTemplate").Parse(`
<a href=/>
home</a>

<h1>Pessoas</h1>
<table>
	<tr>
		<th>Nome</th>
	</tr>
{{range $k, $v := .}}
	<tr>
		<td><a href=/pessoas?pessoaid={{$k}}>{{ $v.Nome }}</a></td>
	</tr>
{{end}}
</table>
`))

type VooPagina struct {
	*Voo
	Passageiros []*Pessoa
}

// RECEBE UM vooPagina
var VooTemp = template.Must(template.New("PessoaTemplate").Parse(`
<a href=/>home</a>

<h1><a href=voos>Voo</a> #{{.VooID}}</h1>
<table>
<tbody>
	<tr>
		<th>Numero</th>
		<td>{{.VooID}}
	</tr>
	<tr>
		<th>Data</th>
		<td>{{.Data}}</td>
	</tr>
	<tr>
		<th><a href=naves>Aeronave</a></th>
		<td><a href=naves?naveid={{.AeronaveID}}>{{.AeronaveID}}</a></td>
	</tr>
	<tr>
		<th>Origem</th>
		<td><a href=aeroportos?codigo={{.OrigemID}}>{{.OrigemID}}</a></td>
	</tr>
	<tr>
		<th>Destino</th>
		<td><a href=aeroportos?codigo={{.DestinoID}}>{{.DestinoID}}</a></td>
	</tr>
</tbody>
</table>
<h3>Passageiros</h3>
<table>
	<tr>
		<th>Nome</th>
	</tr>
{{range .Passageiros}}
	<tr>
		<td><a href=pessoas?pessoaid={{.PessoaID}}>{{ .Nome }}</a></td>
	<tr>
{{end}}
</table>
`))

// Recebe um Mapa de Voos
var MapVoosTemp = template.Must(template.New("VoosTemplate").Parse(`
<a href=/>
home</a>

<h1>Voos</h1>
<table>
	<tr>
		<th>NumVoo</th>
		<th>Data</th>
		<th><a href=naves>NumAeronave</a></th>
		<th>Origem</th>
		<th>Destino</th>
	</tr>
{{range $k, $v := .}}
	<tr>
		<td><a href=voos?vooid={{$k}}>{{$k}}</a></td>
		<td>{{$v.Data}}</td>
		<td><a href=naves?naveid={{.AeronaveID}}>{{.AeronaveID}}</a></td>
		<td><a href=aeroportos?codigo={{$v.OrigemID}}>{{$v.OrigemID}}</a></td>
		<td><a href=aeroportos?codigo={{$v.DestinoID}}>{{$v.DestinoID}}</a></td
	</tr>
{{end}}
</table>
`))

// Deve guardar localização dos aeroportos também
type PaginaPorto struct {
	Aeroporto
	Nome        string //Cidade Estado
	VoosOrigem  []*Voo
	VoosDestino []*Voo
	// Pais string
}

// Recebe um *PaginaPorto
var PaginaPortoTemp = template.Must(template.New("PaginaPortoTemplate").Parse(`
<a href=/>home</a>

<h1><a href=aeroportos>Aeroporto</a> #{{.Titulo}}</h1>
<img src="{{.ImagemURL.Value}}" alt="certamente nao eh um gato">
<table>
<tbody>
	<tr>
		<th>Codigo</th>
		<td>{{.CodigoAeroporto}}</td>
	</tr>
	<tr>
		<th><a href=locais>Cidade</a></th>
		<td><a href=locais?localid={{.LocalID}}>{{.Nome}}</a></td>
	</tr>
</tbody>
</table>
<h3>Descricao</h3>
<p>{{.Descricao}}</p>

<h3>Voos Origem</h3>
<table>
	<tr>
		<th><a href=voos>NumVoo</a></th>
		<th>Origem</th>
		<th>Destino</th>
	</tr>
{{range .VoosOrigem}}
	<tr>
		<td><a href=voos?vooid={{.VooID}}>{{.VooID}}</a></td>
		<td><a href=aeroportos?codigo={{.OrigemID}}>{{.OrigemID}}</a></td>
		<td><a href=aeroportos?codigo={{.DestinoID}}>{{.DestinoID}}</a></td>
	<tr>
{{end}}
</table>

<h3>Voos Destino</h3>
<table>
	<tr>
		<th><a href=voos>NumVoo</a></th>
		<th>Origem</th>
		<th>Destino</th>
	</tr>
{{range .VoosDestino}}
	<tr>
		<td><a href=voos?vooid={{.VooID}}>{{.VooID}}</a></td>
		<td><a href=aeroportos?codigo={{.OrigemID}}>{{.OrigemID}}</a></td>
		<td><a href=aeroportos?codigo={{.DestinoID}}>{{.DestinoID}}</a></td>
	<tr>
{{end}}
</table>



`))

// Recebe um map[string]*PaginaPorto
var MapAeroportosTemp = template.Must(template.New("MapPaginaPortoSTemplate").Parse(`
<a href=/>home</a>

<h1>Aeroportos</h1>
<table>
	<tr>
		<th>Codigo</th>
		<th>Cidade</th>
	</tr>
{{range $k, $v := .}}
	<tr>
		<td><a href=aeroportos?codigo={{$k}}>{{$k}}</a></td>
		<td><a href=locais?localid={{$v.LocalID}}>{{$v.Nome}}</a></td>
	</tr>
{{end}}
</table>
`))

type PaginaLocal struct {
	*Local
	// PessoasNascidas []*Pessoa
	Aeroportos []*Aeroporto
}

// recebe um PaginaLocal
var LocalTemp = template.Must(template.New("LocalTemplate").Parse(`
<a href=/>home</a>

<h1><a href=locais>Local</a> {{.Titulo}}</h1>
<img src="{{.ImagemURL.Value}}" alt="certamente nao eh um gato">
<table>
<tbody>
	<tr>
		<th>LocalID</th>
		<td>{{.LocalID}}</td>
	</tr>
	<tr>
		<th>Nome</th>
		<td>{{.CidadeEstado}}</td>
	</tr>
</tbody>
</table>
<h3>Descricao</h3>
<p>{{.Descricao}}</p>

<h3><a href=aeroportos>Aeroportos</a> da Cidade</h3>
<table>
	{{range .Aeroportos}}
	<tr>
		<th><a href=aeroportos?codigo={{.CodigoAeroporto}}>{{.CodigoAeroporto}}</th>
	</tr>
	{{end}}
</table>
`))

// Recebe um map[int]*Local
var MapLocaisTemp = template.Must(template.New("MapLocaisTemplate").Parse(`
<a href=/>home</a>

<h1>Locais</h1>
<table>
	<tr>
		<th>LocalID</th>
		<th>Nome</th>
	</tr>
{{range $k, $v := .}}
	<tr>
		<td><a href=locais?localid={{$k}}>{{$k}}</a></td>
		<td>{{$v.CidadeEstado}}</td>
	</tr>
{{end}}
</table>
`))

type AeronavePagina struct {
	*Aeronave
	VoosFeitos []*Voo
}

// Recebe um AeronavePagina
var AeronaveTemp = template.Must(template.New("AeronaveTemplate").Parse(`
<a href=/>home</a>


<h1><a href=naves>Aeronave</a> #{{.NumCauda}}</h1>
<img src="{{.ImagemURL.Value}}" alt="certamente nao eh um gato">
<table>
<tbody>
	<tr>
		<th>AeronaveID</th>
		<td>{{.AeronaveID}}</td>
	</tr>
	<tr>
		<th>NumCauda</th>
		<td>{{.NumCauda}}</td>
	</tr>
	<tr>
		<th>Modelo</th>
		<td>{{.Modelo}}</td>
	</tr>
	<tr>
		<th>NumAssentos</th>
		<td>{{.NumAssentos}}</td>
	</tr>
	<tr>
		<th>Fabricante</th>
		<td>{{.Fabricante}}</td>
	<tr>
</tbody>
</table>
<h3>Voos</h3>
<table>
	<tr>
		<th><a href=voos>NumVoo</a></th>
		<th>Origem</th>
		<th>Destino</th>
	</tr>
{{range .VoosFeitos}}
	<tr>
		<td><a href=voos?vooid={{.VooID}}>{{.VooID}}</a></td>
		<td><a href=aeroportos?codigo={{.OrigemID}}>{{.OrigemID}}</a></td>
		<td><a href=aeroportos?codigo={{.DestinoID}}>{{.DestinoID}}</a></td>
	<tr>
{{end}}
</table>
`))

// Recebe um mapa de Aeronaves
var MapAeronaveTemp = template.Must(template.New("MapaAeronaveTemplate").Parse(`
<a href=/>home</a>

<h1>Aeronave</h1>
<table>
	<tr>
		<th>ID</th>
		<th>Cauda</th>
	</tr>
{{range $k, $v := .}}
	<tr>
		<td><a href=naves?naveid={{$k}}>{{$k}}</a></td>
		<td>{{$v.NumCauda}}</td>
	</tr>
{{end}}
</table>
`))

// Adicionar link para Jeffrey epstein
var HomeTemplate = template.Must(template.New("HomeTemplate").Parse(`
<h1> Raestrando Epstein </h1>
<p>Esse projeto tem como intuito rastrear o pedófilo e traficante de humanos Jeffrey Epstein,
tendo como base, os voos que ele fez em um período de 20 anos</p>
<table>
	<tr>
	<th><a href=pessoas> Pessoas </a></th>
	</tr>

	<tr>
	<th><a href=voos> Voos </a></th>
	</tr>

	<tr>
	<th><a href=locais> Locais </a></th>
	</tr>

	<tr>
	<th><a href=naves> Aeronaves </a></th>
	</tr>

	<tr>
	<th><a href=aeroportos> Aeroportos </a></th>
	</tr>
</table>
`))
