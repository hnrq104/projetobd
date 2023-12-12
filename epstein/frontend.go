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
<!DOCTYPE html>
<html class="side-page">
<head>
	<meta charset="UTF-8">
	<title>Pessoa</title>
	<link rel="stylesheet" href="/css/style.css">
</head>
<body>
    <header class="side-header">
    	<a class="return" href=/>Home</a>
	</header>
	
    <h1 class="titulo"><a class="links" href=pessoas>Pessoa</a> {{.Nome}}</h1>
    <img class="main-image" src="{{.ImagemURL.Value}}" alt="certamente nao eh um gato">
    <table class="infoTable">
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
        <h3 class="descricao">Descricao</h3>
        <p class="text">{{.Descricao}}</p>

    <table class="infoTable">
        <caption>Voos</caption>
        <tr>
            <th><a class="links" href=voos>NumVoo</a></th>
            <th>Origem</th>
            <th>Destino</th>
        </tr>
    {{range .VoosFeitos}}
        <tr>
            <td><a class="links" href=voos?vooid={{.VooID}}>{{.VooID}}</a></td>
            <td><a class="links" href=aeroportos?codigo={{.OrigemID}}>{{.OrigemID}}</a></td>
            <td><a class="links" href=aeroportos?codigo={{.DestinoID}}>{{.DestinoID}}</a></td>
            <tr>
    {{end}}
    </table>
</body>
</html>
`))

// Recebe um Mapa de Pessoas
var MapPessoasTemp = template.Must(template.New("PessoasTemplate").Parse(`
<!DOCTYPE html>
<html class="side-page">
<head>
	<meta charset="UTF-8">
	<title>Pessoas</title>
	<link rel="stylesheet" href="/css/style.css">
</head>
<body>
	<header class="side-header">
    	<a class="return" href=/>Home</a>
	</header>
    <h1 class="titulo">Pessoas</h1>
    <table class="infoTable">
    	<tr>
    		<th>Nome</th>
    	</tr>
		{{range $k, $v := .}}
    	<tr>
    		<td><a class="links" href=/pessoas?pessoaid={{$k}}>{{ $v.Nome }}</a></td>
		</tr>
		{{end}}
    </table>
</body>
</html>
`))

type VooPagina struct {
	*Voo
	Passageiros []*Pessoa
}

// RECEBE UM vooPagina
var VooTemp = template.Must(template.New("VooTemplate").Parse(`
<!DOCTYPE html>
<html class="side-page">
<head>
	<meta charset="UTF-8">
	<title>Voo</title>
	<link rel="stylesheet" href="/css/style.css">
</head>
<body>
	<header class="side-header">
    	<a class="return" href=/>Home</a>
	</header>
	<h1 class="titulo"><a class="links" href=voos>Voo</a> #{{.VooID}}</h1>
	<table class="infoTable"> 
    <tbody>
    	<tr>
    		<th>Numero</th>
			<th>Data</th>
			<th><a class="links" href=naves>Aeronave</a></th>
			<th>Origem</th>
			<th>Destino</th>
    	</tr>
    	<tr>
			<td>{{.VooID}}</td>
    		<td>{{.Data}}</td>
			<td><a class="links" href=naves?naveid={{.AeronaveID}}>{{.AeronaveID}}</a></td>
			<td><a class="links" href=aeroportos?codigo={{.OrigemID}}>{{.OrigemID}}</a></td>
			<td><a class="links" href=aeroportos?codigo={{.DestinoID}}>{{.DestinoID}}</a></td>

		</tr>
    </tbody>
    </table>

    <table class="infoTable">
		<caption>Passageiros</caption>
	    <tr>
	    	<th>Nome</th>
	    </tr>
		{{range .Passageiros}}
    	<tr>
    		<td><a class="links" href=pessoas?pessoaid={{.PessoaID}}>{{ .Nome }}</a></td>
		<tr>
		{{end}}
    </table>
</body>
</html>
`))

// Recebe um Mapa de Voos
var MapVoosTemp = template.Must(template.New("VoosTemplate").Parse(`
<!DOCTYPE html>
<html class="side-page">
<head>
	<meta charset="UTF-8">
	<title>Voos</title>
	<link rel="stylesheet" href="/css/style.css">
</head>

<body>
	<header class="side-header">
		<a class="return" href=/>Home</a>
	</header>
	<h1 class="titulo">Voos</h1>
	<table class="infoTable">
		<tr>
			<th>NumVoo</th>
			<th>Data</th>
			<th><a class="links" href=naves>NumAeronave</a></th>
			<th>Origem</th>
			<th>Destino</th>
		</tr>
		{{range $k, $v := .}}
		<tr>
			<td><a class="links" href=voos?vooid={{$k}}>{{$k}}</a></td>
			<td>{{$v.Data}}</td>
			<td><a class="links" href=naves?naveid={{.AeronaveID}}>{{.AeronaveID}}</a></td>
			<td><a class="links" href=aeroportos?codigo={{$v.OrigemID}}>{{$v.OrigemID}}</a></td>
			<td><a class="links" href=aeroportos?codigo={{$v.DestinoID}}>{{$v.DestinoID}}</a></td>
		</tr>
		{{end}}
	</table>
</body>
</html>
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
<!DOCTYPE html>
<html class="side-page">
<head>
	<meta charset="UTF-8">
	<title>Aeroporto</title>
	<link rel="stylesheet" href="/css/style.css">
</head>
<body>

	<header class="side-header">
    	<a class="return" href=/>Home</a>
	</header>

<h1 class="titulo"><a class="links" href=aeroportos>Aeroporto</a> #{{.Titulo}}</h1>
<img src="{{.ImagemURL.Value}}" alt="certamente nao eh um gato">
<table class="infoTable">
<tbody>
	<tr>
		<th>Codigo</th>
		<td>{{.CodigoAeroporto}}</td>
	</tr>
	<tr>
		<th><a class="links" href=locais>Cidade</a></th>
		<td><a class="links" href=locais?localid={{.LocalID}}>{{.Nome}}</a></td>
	</tr>
</tbody>
</table>
<h3 class="descricao">Descricao</h3>
<p class="text">{{.Descricao}}</p>

<table class="infoTable">
	<caption>Voos Origem</caption>
	<tr>
		<th><a class="links" href=voos>NumVoo</a></th>
		<th>Origem</th>
		<th>Destino</th>
	</tr>
{{range .VoosOrigem}}
	<tr>
		<td><a class="links" href=voos?vooid={{.VooID}}>{{.VooID}}</a></td>
		<td><a class="links" href=aeroportos?codigo={{.OrigemID}}>{{.OrigemID}}</a></td>
		<td><a class="links" href=aeroportos?codigo={{.DestinoID}}>{{.DestinoID}}</a></td>
	<tr>
{{end}}
</table>

<table class="infoTable">
	<caption>Voos Destino</caption>
	<tr>
		<th><a class="links" href=voos>NumVoo</a></th>
		<th>Origem</th>
		<th>Destino</th>
	</tr>
{{range .VoosDestino}}
	<tr>
		<td><a class="links" href=voos?vooid={{.VooID}}>{{.VooID}}</a></td>
		<td><a class="links" href=aeroportos?codigo={{.OrigemID}}>{{.OrigemID}}</a></td>
		<td><a class="links" href=aeroportos?codigo={{.DestinoID}}>{{.DestinoID}}</a></td>
	<tr>
{{end}}
</table>
</body>
</html>
`))

// Recebe um map[string]*PaginaPorto
var MapAeroportosTemp = template.Must(template.New("MapPaginaPortoSTemplate").Parse(`
<!DOCTYPE html>
<html class="side-page">
<head>
	<meta charset="UTF-8">
	<title>Aeroportos</title>
	<link rel="stylesheet" href="/css/style.css">
</head>

<body>
	<header class="side-header">
    	<a class="return" href=/>Home</a>
	</header>

<h1 class="titulo">Aeroportos</h1>
<table class="infoTable">
	<tr>
		<th>Codigo</th>
		<th>Cidade</th>
	</tr>
{{range $k, $v := .}}
	<tr>
		<td><a class="links" href=aeroportos?codigo={{$k}}>{{$k}}</a></td>
		<td><a class="links" href=locais?localid={{$v.LocalID}}>{{$v.Nome}}</a></td>
	</tr>
{{end}}
</table>
</body>
</html>
`))

type PaginaLocal struct {
	*Local
	// PessoasNascidas []*Pessoa
	Aeroportos []*Aeroporto
}

// recebe um PaginaLocal
var LocalTemp = template.Must(template.New("LocalTemplate").Parse(`
<!DOCTYPE html>
<html class="side-page">
<head>
	<meta charset="UTF-8">
	<title>Aeronave</title>
	<link rel="stylesheet" href="/css/style.css">
</head>

<body>
    <header class="side-header">
		<a class="return" href=/>Home</a>
	</header>

    <h1 class="titulo"><a class="links" href=locais>Local</a> {{.Titulo}}</h1>
    <img src="{{.ImagemURL.Value}}" alt="certamente nao eh um gato">
    <table class="infoTable">
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
    <h3 class="descricao">Descricao</h3>
    <p class="text">{{.Descricao}}</p>

    <h3 class="titulo"><a class="links" href=aeroportos>Aeroportos</a> da Cidade</h3>
    <table class="infoTable">
    	{{range .Aeroportos}}
    	<tr>
    		<th><a class="links" href=aeroportos?codigo={{.CodigoAeroporto}}>{{.CodigoAeroporto}}</th>
    	</tr>
    	{{end}}
    </table>
    </body>
</html>
`))

// Recebe um map[int]*Local
var MapLocaisTemp = template.Must(template.New("MapLocaisTemplate").Parse(`
<!DOCTYPE html>
<html class="side-page">
<head>
	<meta charset="UTF-8">
	<title>Locais</title>
	<link rel="stylesheet" href="/css/style.css">
</head>
<body>
	    <header class="side-header">
	    	<a class="return" href=/>Home</a>
		</header>


	<h1 class="titulo">Locais</h1>
	<table class="infoTable">
		<tr>
			<th>LocalID</th>
			<th>Nome</th>
		</tr>
	{{range $k, $v := .}}
		<tr>
			<td><a class="links" href=locais?localid={{$k}}>{{$k}}</a></td>
			<td>{{$v.CidadeEstado}}</td>
		</tr>
	{{end}}
	</table>
</body>
</html>
`))

type AeronavePagina struct {
	*Aeronave
	VoosFeitos []*Voo
}

// Recebe um AeronavePagina
var AeronaveTemp = template.Must(template.New("AeronaveTemplate").Parse(`
<!DOCTYPE html>
<html class="side-page">
<head>
	<meta charset="UTF-8">
	<title>Aeronave</title>
	<link rel="stylesheet" href="/css/style.css">
</head>
<body>
	<header class="side-header">
    	<a class="return" href=/>Home</a>
	</header>
    <h1 class="titulo"><a class="links" href=naves>Aeronave</a> #{{.NumCauda}}</h1>
<img src="{{.ImagemURL.Value}}" alt="certamente nao eh um gato">
<table class="infoTable">
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
<table>
	<caption>Voos</caption>
	<tr>
		<th><a class="links" href=voos>NumVoo</a></th>
		<th>Origem</th>
		<th>Destino</th>
	</tr>
{{range .VoosFeitos}}
	<tr>
		<td><a class="links" href=voos?vooid={{.VooID}}>{{.VooID}}</a></td>
		<td><a class="links" href=aeroportos?codigo={{.OrigemID}}>{{.OrigemID}}</a></td>
		<td><a class="links" href=aeroportos?codigo={{.DestinoID}}>{{.DestinoID}}</a></td>
	<tr>
{{end}}
</table>
</body>
</html>
`))

// Recebe um mapa de Aeronaves
var MapAeronaveTemp = template.Must(template.New("MapaAeronaveTemplate").Parse(`
<!DOCTYPE html>
<html class="side-page">
<head>
	<meta charset="UTF-8">
	<title>Aeronaves</title>
	<link rel="stylesheet" href="/css/style.css">
</head>
<body>
	<header class="side-header">
    	<a class="return" href=/>Home</a>
	</header>
	<h1 class="titulo">Aeronaves</h1>
	<table class="infoTable">
		<tr>
			<th>ID</th>
			<th>Cauda</th>
		</tr>
	{{range $k, $v := .}}
		<tr>
			<td><a class="links" href=naves?naveid={{$k}}>{{$k}}</a></td>
			<td>{{$v.NumCauda}}</td>
		</tr>
	{{end}}
	</table>
</body>
</html>
`))

// Adicionar link para Jeffrey epstein
var HomeTemplate = template.Must(template.New("HomeTemplate").Parse(`
<!DOCTYPE html>
<html class="main-page">
<head>
	<meta charset="UTF-8">
	<title>Rastreando Epstein</title>
	<link rel="stylesheet" href="/css/style.css" media="screen">
</head>
<body>
	<header class="main-header">
		<h1 class="Titulo"> Rastreando Epstein </h1>
	</header>
	<p class="descricao">Esse projeto tem como intuito rastrear o pedófilo e traficante de humanos Jeffrey Epstein,
		tendo como base, os voos que ele fez em um período de 20 anos.</p>
	<table class="home-table">
		<tr>
		<th><a class="links" href=pessoas> Pessoas </a></th>
		</tr>

		<tr>
		<th><a class="links" href=voos> Voos </a></th>
		</tr>

		<tr>
		<th><a class="links" href=locais> Locais </a></th>
		</tr>

		<tr>
		<th><a class="links" href=naves> Aeronaves </a></th>
		</tr>

		<tr>
		<th><a class="links" href=aeroportos> Aeroportos </a></th>
		</tr>
	</table>
</body>
</html>
`))
