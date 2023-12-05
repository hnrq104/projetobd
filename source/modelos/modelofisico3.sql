create table Localidade(
LocalID int primary key,
Descricao text,
nome varchar(60),
titulo text,
urlImagem text
);

create table Pessoa(
PessoaID int primary key, 
Nome varchar(60) not null,
titulo text,
Iniciais varchar(10),
-- Conhecido bit,
-- DataNasc date,
-- DataMorte date,
urlImagem text,
Descricao text
-- CidadeNascimento int,
-- CidadeMorte int,
-- foreign key (CidadeNascimento) references Localidade(LocalID) on delete set null,
-- foreign key (CidadeMorte) references Localidade(LocalID) on delete set null,
);

create table Aeronave(
AeronaveID int primary key,
NumAssentos int,
NumCauda varchar(15),
Modelo varchar(30),
Fabricante varchar(40),
urlImagem text
);

create table Aeroporto(
Codigo varchar(10) primary key,
Localizacao int,
urlImagem text,
Descricao text,
titulo text,
foreign key (Localizacao) references Localidade(LocalID)
);

create table Voo(
VooID int primary key,
-- NumPassageiros int,
DataVoo date,
Origem varchar(10),
Destino varchar(10),
Nave int,
foreign key (Origem) references Aeroporto(Codigo) on delete set null,
foreign key (Destino) references Aeroporto(Codigo) on delete set null,
foreign key (Nave) references Aeronave(AeronaveID) on delete set null
);

create table Embarcam(
fk_Voo int not null,
fk_Pessoa int not null,
foreign key (fk_Voo) references Voo(VooID) on delete cascade,
foreign key (fk_Pessoa) references Pessoa(PessoaID) on delete cascade
);



