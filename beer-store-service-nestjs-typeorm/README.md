# beer-store-service-nestjs-typeorm

Beer service implementation using typescript and the best of what it have to
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

While nestjs isn't exactly a form-free framework, it's quite more flexible than
other implementations seen here in this experiment.

With node ecosystem at hand and type-checking thanks to typescript, it has a
good dependency injection approach and a cli tool that does not delivers an
overbloated project template:

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

Regarding typeorm i find quite elegant it's migration strategy. See! Let's
create one migration file:

```bash
npm run typeorm migration:create -- -n InitialSchema
```

Now we complete the file:

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

A few tweaks where made in order to both cli and app migrations could properly
work, mostly because typescript. key change was the cli being called from npm
scripts due the need of proper typescript interpretation. Also a small
unintuitive change on app migrations dir because we need to use the real runtime
content when doing this.

One caveat is when you decide to manage migrations entirely from command line.
You will need to run the application at least once to "compile" migrations.

```bash
npm run typeorm migration:create -- -n pto 
# fill up migration file, then
npm run dev
# migration gets compiled into js now migrate:show, migrate:revert and 
# migrate:run will work as expected.
```
