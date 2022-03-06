drop database if exists db_maps;
create database db_maps;

use db_maps;


create table maps (
	id int unsigned auto_increment primary key,
    street varchar(255) not null,
    district varchar(255) not null,
    city varchar(255) not null,
    country varchar(255) not null,
    state varchar(255) not null,
    number varchar(255) not null,
    `created_at` datetime default now()
);
