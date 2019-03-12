# beer-store-service-python-flask-sqlalchemy

This project shows how the beerservice can be implemented using the python
ecosystem

The flask microframework provides the rest api

The sqlalchemy handles the database access and schema migrations

## How do i run this

You **must have** to install [pipenv](https://github.com/pypa/pipenv).

Then open a terminal on the folder containing this readme and

```bash
pipenv install
pipenv run dev
```

To manage database migrations you can use this convenience shortcut (see the 
Pipfile):

```bash
pipenv run db
```

Pretty much like the node/knex version, this one has a full featured migration
system able to upgrade, downgrade and also can either do sql or python migration
unit scripts.
