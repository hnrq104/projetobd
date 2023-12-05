create table Localidade(
LocalID int auto_increment primary key,
Descricao text,
CidadeEstado varchar(40),
NomePais varchar(20),
urlImagem text
);

create table Pessoa(
PessoaID int auto_increment primary key, 
Nome varchar(30) not null,
Iniciais varchar(6),
Conhecido bit,
DataNasc date,
DataMorte date,
Descricao text,
CidadeNascimento int,
CidadeMorte int,
urlImagem text
foreign key (CidadeNascimento) references Localidade(LocalID) on delete set null,
foreign key (CidadeMorte) references Localidade(LocalID) on delete set null,
);

create table Aeronave(
AeronaveID int auto_increment primary key,
NumDeAssentos int,
NumCauda varchar(15),
Modelo varchar(30) not null,
Fabricante varchar(40) not null,
urlImagem text
);

create table Aeroporto(
Codigo varchar(10) primary key,
Descricao text,
Localizacao int,
urlImagem text,
foreign key (Localizacao) references Localidade(LocalID)
);

create table Voo(
VooID int auto_increment primary key,
NumPassageiros int,
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



