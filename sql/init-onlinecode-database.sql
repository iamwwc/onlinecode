create database if not exists onlinecode;
use onlinecode;
create table if not exists share (hashkey varchar(32) primary key, codesvalue varchar(10000) not null) default charset utf8MB4;