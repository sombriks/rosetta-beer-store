#!/usr/bin/env python
from migrate.versioning.shell import main
import os

DB_URL = 'sqlite:///beerstore.db'

if 'DB_URL' in os.environ:
    DB_URL = os.environ['DB_URL']

if __name__ == '__main__':
    main(repository='migrations', url=DB_URL, debug='False')
