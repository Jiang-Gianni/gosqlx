drop table if exists rdms;

create table if not exists rdms(
    id integer primary key,
    system text not null
);

insert into rdms(id, system) values (1, 'PostgreSQL');
insert into rdms(id, system) values (2, 'MySQL');
insert into rdms(id, system) values (3, 'SQL Server');
insert into rdms(id, system) values (4, 'Oracle');
