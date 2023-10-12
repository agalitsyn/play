use Mix.Config

config :ectotest, :ecto_repos, [Ectotest.Repo]

config :ectotest, Ectotest.Repo,
  username: "postgres",
  password: "postgres",
  database: "postgres",
  hostname: "localhost",
  port: "5432"
