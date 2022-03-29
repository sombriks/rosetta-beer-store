import { MigrationInterface, QueryRunner } from 'typeorm';

export class InitialData1648530670150 implements MigrationInterface {
  public async up(queryRunner: QueryRunner): Promise<void> {
    return queryRunner.query(`
        INSERT INTO "beer" 
            ("idbeer", "creationdatebeer", "titlebeer", "descriptionbeer") 
        VALUES 
            ('1', '2022-03-29 03:16:42', 'Brahma', 'A número 1!'),
            ('2', '2022-03-29 03:16:42', 'Antartica Original', 'Pilsen'),
            ('3', '2022-03-29 03:16:42', 'Itaipava', 'A cerveja do verão!'),
            ('4', '2022-03-29 03:16:42', 'Devassa', 'Tropical Lager'),
            ('5', '2022-03-29 03:16:42', 'Corona', 'Extra'),
            ('6', '2022-03-29 03:16:42', 'Therezópolis', 'Cerveja especial'),
            ('7', '2022-03-29 03:16:42', 'Budweiser', 'King of Beers'),
            ('8', '2022-03-29 03:16:42', 'Heineken', 'Premium quality'),
            ('9', '2022-03-29 03:16:42', 'Skol', 'A que desce redondo!'),
            ('10', '2022-03-29 03:16:42', 'Kaiser', 'Cerveja bem cervejada'),
            ('11', '2022-03-29 03:16:42', 'Eisenbahn', 'Cerveja puro malte'),
            ('12', '2022-03-29 03:16:42', 'Liefmans', 'Fruitesse'),
            ('13', '2022-03-29 03:16:42', 'Bohemia', 'Cerveja pilsen');
    `);
  }

  public async down(queryRunner: QueryRunner): Promise<void> {
    return queryRunner.query('delete from beer;');
  }
}
