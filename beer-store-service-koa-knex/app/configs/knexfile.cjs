const _cfg = {
    client: 'better-sqlite3',
    useNullAsDefault: true,
    connection: {
        filename: `${__dirname}/../../beerstore.sqlite3`,
    },
    pool: {
        min: 2,
        max: 10
    },
    migrations: {
        directory:`${__dirname}/migrations`,
        loadExtensions: [".mjs"],
    },
}
module.exports = {
    development: { ..._cfg},
    test: {
        ..._cfg,
        connection: {
            // filename: ':memory:',
            filename: `${__dirname}/../../beerstore-test.sqlite3`,
        }
    },
    production: {
        ..._cfg,
        client: 'postgresql',
        connection: process.env.PG_CONNECTION_URL
    }
}
