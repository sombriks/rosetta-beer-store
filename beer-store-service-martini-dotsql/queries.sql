-- name: create-beer-table
create table if not exists beer (
  idbeer integer primary key autoincrement,
  titlebeer varchar(255) not null
)