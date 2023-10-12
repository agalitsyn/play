# Ectotest

[YT series](https://www.youtube.com/watch?v=IFKG4Hgt-zM&list=PLFhQVxlaKQElscjMvMmyMCaZ9mxf4XAw-)

Test:

```sh
docker compose up -d

mix deps.get
mix ecto.create
mix ecto.migrate

pgcli postgres://postgres:postgres@localhost:5432/postgres
\d users


iex -S mix
import_file("test.exs")


mix ecto.drop
docker compose down --volumes
```
