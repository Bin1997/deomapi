
create database test;

use test;

create table users(
id int primary key ,
name varchar(20),
age int ,
gender varchar(10),
email varchar(20)
);

insert into user(id,name,age,gender,email) values (1,'tom',10,'boy','111@222');

insert into user(id,name,age,gender,email) values (2,'jerry',10,'boy','222@333');

insert into user(id,name,age,gender,email) values (3,'alice',10,'girl','333@444');