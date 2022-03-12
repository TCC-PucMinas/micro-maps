-- drop database if exists db_maps;
create database db_maps;

use db_maps;


create table maps (
	id int unsigned auto_increment primary key,
	inline varchar(255) not null,
    street varchar(255) not null,
    district varchar(255) not null,
    city varchar(255) not null,
    country varchar(255) not null,
    state varchar(255) not null,
    zipCode varchar(255) not null,
    number varchar(255) not null,
    lat varchar(255) not null,
    lng varchar(255) not null,
    `created_at` datetime default now()
);

create table calculates(
    id int unsigned auto_increment primary key,
    origin json not null,
    destiny json not null,
    humanReadble varchar(255) not null,
    meters int not null,
    `created_at` datetime default now()
)