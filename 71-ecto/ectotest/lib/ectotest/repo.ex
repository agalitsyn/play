defmodule Ectotest.Repo do
   use Ecto.Repo,
    otp_app: :ectotest,
    adapter: Ecto.Adapters.Postgres
end
