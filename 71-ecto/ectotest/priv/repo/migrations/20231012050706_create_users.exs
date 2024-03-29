defmodule Ectotest.Repo.Migrations.CreateUserss do
  use Ecto.Migration

  def change do
    create table(:users) do
        add :email, :string, null: false

        timestamps()
    end

    create unique_index(:users, [:email])
  end
end
