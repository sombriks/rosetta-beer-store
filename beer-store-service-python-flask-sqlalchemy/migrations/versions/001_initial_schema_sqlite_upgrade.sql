-- name: create-media-table
create table media (
  idmedia integer primary key autoincrement,
  creationdatemedia timestamp not null default CURRENT_TIMESTAMP,
  datamedia blob not null,
  nomemedia varchar(255) not null,
  mimemedia varchar(255) not null
);

-- name: create-beer-table
create table beer (
  idbeer integer primary key autoincrement,
  creationdatebeer timestamp not null default CURRENT_TIMESTAMP,
  titlebeer varchar(255) not null,
  descriptionbeer text,
  idmedia integer,

  foreign key (idmedia) references media(idmedia)
)
