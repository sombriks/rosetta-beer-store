-- name: create-media-table
create table if not exists media (
  idmedia integer primary key autoincrement,
  creationdatemedia timestamp not null default CURRENT_TIMESTAMP,
  datamedia blob not null,
  nomemedia varchar(255) not null,
  mimemedia varchar(255) not null
);

-- name: create-beer-table
create table if not exists beer (
  idbeer integer primary key autoincrement,
  creationdatebeer timestamp not null default CURRENT_TIMESTAMP,
  titlebeer varchar(255),
  descriptionbeer text,
  idmedia integer,

  foreign key (idmedia) references media(idmedia)
)

-- name: count-beers
select count(*) from beer;

-- name: insert-a-few-beers
insert into 
  beer (idbeer,titlebeer,descriptionbeer)
values
  (1, 'Brahma', 'A número 1!'),
  (2, 'Antartica Original', 'Pilsen'),
  (3, 'Itaipava', 'A cerveja do verão!'),
  (4, 'Devassa', 'Tropical Lager'),
  (5, 'Corona', 'extra'),
  (6, 'therezópolis', 'Cerveja especial'),
  (7, 'Budweiser', 'King of Beers'),
  (8, 'Heinenken', 'Premium quality'),
  (9, 'Skol', 'A que desce redondo!'),
  (10, 'Kaiser', 'Cerveja bem cervejada!'),
  (11, 'Eisenbahn', 'Cerveja puro malte'),
  (12, 'Liefmans', 'Fruitesse'),
  (13, 'Bohemia', 'cerveja Pilsen');

-- name: find-beers
select * from beer where titlebeer like ? or descriptionbeer like ? limit ? offset ?
