# beer-store-service-express-knex

This sample repo shows how to build the beer store service using express and 
knex with bookshelf in a relational way.

[Express](https://expressjs.com/) does all the routing while 
[knex](http://knexjs.org/) manages databases migrations and 
[bookshelf](http://bookshelfjs.org/) maps tables.

We also use [nodemon](https://nodemon.io/) for automated process reload.

Although we use sqlite as database 'server', knex provides integration with 
mysql, postgresql, mssql and oracle.

## How do i run this?

Open a terminal in this folder and:

```bash
npm install
npm run dev
```

I'll assume that you have a modern nodejs/npm setup in place.
