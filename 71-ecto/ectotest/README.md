# Ectotest

Some steps from [YT series](https://www.youtube.com/watch?v=IFKG4Hgt-zM&list=PLFhQVxlaKQElscjMvMmyMCaZ9mxf4XAw-) and from [ecto tutorial](https://hexdocs.pm/ecto/getting-started.html)

Test:

```sh
# skipped steps
# 1. install elixir
# 2. create new project with supervisor
# mix new ectotest --sup

# start and connect to postgres
cp .env.template .env
docker compose up -d
pgcli postgres://postgres:postgres@localhost:5432/postgres

# create database and schema
mix deps.get
mix ecto.create
# (done) mix ecto.gen.migration create_users
mix ecto.migrate

# run test script
iex -S mix
import_file("test.exs")

# cleanup
mix ecto.drop
docker compose down --volumes
```
