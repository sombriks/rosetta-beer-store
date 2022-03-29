import { MigrationInterface, QueryRunner, Table } from 'typeorm';

export class InitialSchema1648520518850 implements MigrationInterface {
  public async up(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.createTable(
      new Table({
        name: 'media',
        columns: [
          {
            name: 'idmedia',
            type: 'integer',
            isGenerated: true,
            isPrimary: true,
          },
          {
            name: 'creationdatemedia',
            type: 'timestamp',
            isNullable: false,
            default: 'CURRENT_TIMESTAMP',
          },
          {
            name: 'datamedia',
            type: 'blob',
            isNullable: false,
          },
          {
            name: 'nomemedia',
            type: 'string',
            isNullable: false,
          },
          {
            name: 'mimemedia',
            type: 'string',
            isNullable: false,
          },
        ],
      }),
    );
    return queryRunner.createTable(
      new Table({
        name: 'beer',
        columns: [
          {
            name: 'idbeer',
            type: 'integer',
            isGenerated: true,
            isPrimary: true,
          },
          {
            name: 'creationdatebeer',
            type: 'timestamp',
            isNullable: false,
            default: 'CURRENT_TIMESTAMP',
          },
          {
            name: 'titlebeer',
            type: 'string',
            isNullable: false,
          },
          {
            name: 'descriptionbeer',
            type: 'text',
            isNullable: false,
          },
          {
            name: 'idmedia',
            type: 'integer',
          },
        ],
        foreignKeys: [
          {
            columnNames: ['idmedia'],
            referencedTableName: 'media',
            referencedColumnNames: ['idmedia'],
          },
        ],
      }),
    );
  }

  public async down(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.dropTable('beer');
    return queryRunner.dropTable('media');
  }
}
