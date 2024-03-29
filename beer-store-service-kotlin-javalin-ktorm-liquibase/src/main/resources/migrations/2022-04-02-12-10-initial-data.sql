-- liquibase formatted sql
-- changeset sombriks:2022-04-02-12-10-initial-data.sql

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

-- rollback delete from beer;
