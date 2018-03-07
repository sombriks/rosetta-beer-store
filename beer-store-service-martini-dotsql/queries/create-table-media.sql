-- name: create-media-table
create table if not exists media (
  idmedia integer primary key autoincrement,
  creationdatemedia timestamp not null default CURRENT_TIMESTAMP,
  datamedia blob not null,
  nomemedia varchar(255) not null,
  mimemedia varchar(255) not null
);
