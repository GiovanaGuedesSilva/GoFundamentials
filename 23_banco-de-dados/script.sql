use devbook;

create table usuario(
	id int auto_increment primary key,
    nome varchar(50) not null,
    email varchar(50) not null
);