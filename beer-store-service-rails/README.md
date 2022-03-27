# beer-store-service-rails

This service implementation uses _ruby on rails_ to deliver the rest service.

## How to run this

Simply type:

```bash
bundle install
rails db:setup # runs both migrates and seeds
rails s
```

Project generated with ruby 3.1.1 and rails 7.0.2.3

## How easy is to develop with rails

The beer store experience is a little different from what the rails framework
tends to offer, however there is still space for flexibility.

For instance, the entire project skeleton is generated issuing one single
command:

```bash
rails new beer-store-service-rails --api
```

Model creation:

```bash
sombriks@hornet beer-store-service-rails % rails g model beer
      invoke  active_record
      create    db/migrate/20220327175011_create_beers.rb
      create    app/models/beer.rb
      invoke    test_unit
      create      test/models/beer_test.rb
      create      test/fixtures/beers.yml
```

Since we don`t use the expected table name (with plurals) we need to set table
name in the active record model.

Similar issue with our primary key column, a small custom fix and we're good to
go.

Code generation is very similar for controllers:

```bash
sombriks@hornet beer-store-service-rails % rails g controller beer
      create  app/controllers/beer_controller.rb
      invoke  test_unit
      create    test/controllers/beer_controller_test.rb
```

Templates created, it's just a matter of fill up the empty spaces and properly
wire it on `routes.rb`

```ruby
Rails.application.routes.draw do
  # Define your application routes per the DSL in https://guides.rubyonrails.org/routing.html
  scope "/", defaults: {format: :json} do
    get "beer/list", to: "beer#list"
    get "beer/:idbeer", to: "beer#find"
  end
end
```

One small difference is our seed script. In ruby,
**do not use migrations to seed the database**. Every time a new environment is
created, rails uses the generated version of migration files called `schema.rb`
and migrations not producing schema changes don't appear there.

Despite the huge amount of generated code, project structure is quite reasonable
and easy to navigate through. Rails version loses a few points regarding the
seed strategy. We seed it once and for good.

Besides that, is dead easy to deliver an API using rails.
