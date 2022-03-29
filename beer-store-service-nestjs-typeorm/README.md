# beer-store-service-nestjs-typeorm

Beer service implementation using typescript and the 'best' of what it have to
offer.

- [nestjs](https://docs.nestjs.com/first-steps) for controllers, services, etc
- [typeorm](https://typeorm.io/) form data access and database migrations

## how do i run this

Install dependencies and then call the service via npm script:

```bash
npm install
npm run start:dev
```

## trivia

at first, it's quite simple to start working with nest. install it's cli tool:

```bash
npm i -g @nestjs/cli
```

Then create your project:

```bash
nest new beer-store-service-nestjs-typeorm 
```

It didn't came with typeorm out of the box, but it's simple to add:

```bash
npm install @nestjs/typeorm sqlite3 typeorm@0.2
```

Older version of typeorm is needed due to compatibility issues with 0.3 version.

One reason for this is the amount of
[alternatives available](https://docs.nestjs.com/techniques/database).

At first, typeorm seems a reasonable tool, but it's filled with caveats. For
example:

```bash
npm run typeorm migration:create -- -n InitialSchema
```

You will need this custom npm script to properly use the cli tool.

Also the `ormconfig.json` file need fine tuning to allow both cli and app behave
nicely each other:

```json
{
  "type": "sqlite",
  "database": "beerstore.sqlite3",
  "autoLoadEntities": true,
  "migrationsRun": true,
  "synchronize": false,
  "migrations": [
    "dist/migrations/*.js"
  ],
  "entities": [
    "dist/**/*.entity{.ts,.js}"
  ],
  "cli": {
    "migrationsDir": "src/migrations"
  }
}
```

what works for cli does not work for application.

The migration file structure looks a little scary at first, but it's typescript,
there is a lot of good documentation out there:

```ts
import { MigrationInterface, QueryRunner, Table } from 'typeorm';

export class InitialSchema1648520518850 implements MigrationInterface {
  public async up(queryRunner: QueryRunner): Promise<void> {
    return queryRunner.createTable(
      new Table({
        name: 'beer',
        columns: [
          {
            name: 'idbeer',
            type: 'integer',
            isGenerated: true,
            isPrimary: true,
          }, // pretend i put more columns
        ],
      }),
    );
  }

  public async down(queryRunner: QueryRunner): Promise<void> {
    return queryRunner.dropTable('beer');
  }
}
```

One caveat is when you decide to manage migrations entirely from command line.
You will need to run the application at least once to "compile" migrations.

```bash
npm run typeorm migration:create -- -n pto 
# fill up migration file, then
npm run dev
# migration gets compiled into js now migrate:show, migrate:revert and 
# migrate:run will work as expected.
```

Another concerning behavior is, sometimes, nest development mode will swallow
some errors. Delete the dist folder times to times so you don't get stuck into
imaginary issues.

## but is it any good

Besides those drawbacks, the module architecture and dependency injection
strategy are quite good, and this solution delivers what is needed.
