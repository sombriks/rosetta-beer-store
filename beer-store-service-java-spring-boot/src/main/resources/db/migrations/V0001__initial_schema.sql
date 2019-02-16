-- name: create-media-table
create table if not exists media (
  idmedia int primary key auto_increment,
  creationdatemedia timestamp not null default CURRENT_TIMESTAMP(),
  datamedia blob not null,
  nomemedia varchar(255) not null,
  mimemedia varchar(255) not null
);

-- name: create-beer-table
create table if not exists beer (
  idbeer int primary key auto_increment,
  creationdatebeer timestamp not null default CURRENT_TIMESTAMP(),
  titlebeer varchar(255),
  descriptionbeer text,
  idmedia int,

  foreign key (idmedia) references media(idmedia)
)
