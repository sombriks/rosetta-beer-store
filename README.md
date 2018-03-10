# rosetta-beer-store

The very same beer store implemented in various stacks. Also, fontends and backends are interchangeable

## Why?

To compare stacks, analyze solutions, maybe give you a glimpse on what each approach can do.

Also because it's fun.

## What do we have so far?

- API
  - beers listing
  - beer detail
- Client
  - vue.js + browserify 
- Service
  - expressjs + knex
  - martini + sqlx

## How does it work?

Each project will have one README.md explaining how to fire the engines.

You will need at least one service and one client running in order to see a complete work.

We'll have clients and services. 

Services must provide the very same REST API so their clients shall connect on it in a transparent way.

Therefore all services must use port 3000. 

You shall run just one service at time.

The clients have not such limitation. 

However you must correctly point them to the service address.

This is not a simple language comparison, but tooling as well. 

Is there more than one way to set up a java service project? Let's have it!

More than one way to set up a javascript client project! Let's go!

## Tell me more about this Beer Store

- The Beer Store is a modern web app
  - It has a frontend and a backend
  - It has two roles: customer and manager
  - A default user called admin are available from the cold start
    - It has both customer and manager roles
  - New users can be created.
    - New users have only the customer role
    - The admin can grant the manager role for other users
    - An user can not change their own roles
  - **Customers** (i.e. users holding the customer role) can:
    - Register themselves into the store
    - Log in into the beer store
    - List **beers**
    - See beer details
    - Add a beer to the **cart**
    - Remove a beer from the cart
    - Change beer amount in the cart   
    - Checkout the cart (pay for the beers)
    - List past orders
  - **Managers** can:
    - Log in into the beer store
    - Add new beers
    - Change beer stock
    - List beer **orders**
    - See order details
    - Change order status
    - List customers
    - See customer details
- There must have a rest api able to perform all these actions
- Some kind of data schema must be set up to help to persist information
- A few initial data must be present when the store goes up
  - Some beers
  - The admin user with both profiles (customer, manager)
