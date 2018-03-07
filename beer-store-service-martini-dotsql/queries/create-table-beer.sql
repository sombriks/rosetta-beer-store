-- name: create-beer-table
create table if not exists beer (
  idbeer integer primary key autoincrement,
  creationdatebeer timestamp not null default CURRENT_TIMESTAMP,
  titlebeer varchar(255),
  descriptionbeer text,
  idmedia integer,

  foreign key (idmedia) references media(idmedia)
)