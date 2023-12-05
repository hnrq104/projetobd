create table Localidade(
LocalID int auto_increment primary key,
Descricao text,
Nome varchar(20) not null,
NomeCidade varchar(20),
NomeEstado varchar(20),
NomePais varchar(20)
-- url da imagem
);

create table Pessoa(
PessoaID int auto_increment primary key, 
Nome varchar(20) not null,
Conhecido bit,
Datanasc date,
Datamort date,
BreveDescricao text,
CidadeNascimento int,
CidadeMorte int,
foreign key (CidadeNascimento) references Localidade(LocalID) on delete set null,
foreign key (CidadeMorte) references Localidade(LocalID) on delete set null
-- url da imagem
);

create table Aeronave(
AeronaveID int auto_increment primary key,
NumDeAssentos int,
NumCauda int,
Modelo varchar(20) not null,
Fabricante varchar(20) not null
-- url da imagem
);

create table Aeroporto(
Codigo varchar(6) primary key,
Descricao text,
Localizacao int,
foreign key (Localizacao) references Localidade(LocalID)
);

create table Passagem(
PassagemID int auto_increment primary key,
NumVoo int,
NumPassageiros int,
DataVoo date,
Passageiro int,
Origem varchar(6),
Destino varchar(6),
Nave int,
foreign key (Passageiro) references Pessoa(PessoaID) on delete set null,
foreign key (Origem) references Aeroporto(Codigo) on delete set null,
foreign key (Destino) references Aeroporto(Codigo) on delete set null,
foreign key (Nave) references Aeronave(AeronaveID) on delete set null
);

