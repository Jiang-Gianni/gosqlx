drop table if exists connection;

create table if not exists connection(
    id integer primary key,
    id_rdms integer not null,
    ssl boolean default false,
    name text not null,
    host text not null,
    port integer not null,
    datasource text not null,
    user text not null,
    password text not null
);

insert into connection(id, id_rdms, ssl, name, host, port, datasource, user, password)
values (1, 1, false, 'PostgreSQL Docker', 'localhost', 5432, 'mydb', 'root', 'my-secret-pw');

insert into connection(id, id_rdms, ssl, name, host, port, datasource, user, password)
values (2, 2, false, 'MySQL Docker', 'localhost', 3306, 'mydb', 'root', 'my-secret-pw');
