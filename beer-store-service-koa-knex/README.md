# beer-store-service-koa-knex

This is the beer store service implementation using updated modern node and koa,
the spiritual successor of express.

Knex still present, but it's the latest and modern version, so expect all
benefits from modern node development.

We're also sampling dotenv-flow, so we demonstrate how to handle different
environments with a single codebase

For our test cases we're using mocha/chai and for coverage c8, a modern
replacement for istanbul.

- [koa](https://koajs.com/)
- [knex](https://knexjs.org/)
- [dotenv-flow](https://github.com/kerimdzhanov/dotenv-flow)
- [mocha](https://mochajs.org/)
- [chai](https://www.chaijs.com/)
- [c8](https://github.com/bcoe/c8)

Besides that, this is the good old sqlite beer store service running in
development mode.

## How to run this

Intall dependencies:

````bash
npm install
````

Then call 'dev' npm task:

````bash
npm run dev
````

You also can run tests:

````bash
npm run test:coverage
````

