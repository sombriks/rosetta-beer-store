-- https://github.com/rubenv/sql-migrate#writing-migrations

-- +migrate Up
create table if not exists media (
  idmedia integer primary key autoincrement,
  creationdatemedia timestamp not null default CURRENT_TIMESTAMP,
  datamedia blob not null,
  nomemedia varchar(255) not null,
  mimemedia varchar(255) not null
);

create table if not exists beer (
  idbeer integer primary key autoincrement,
  creationdatebeer timestamp not null default CURRENT_TIMESTAMP,
  titlebeer varchar(255),
  descriptionbeer text,
  idmedia integer,

  foreign key (idmedia) references media(idmedia)
);

-- +migrate Down
drop table if exists beer;
drop table if exists media;
