defmodule Ectotest.User do
    use Ecto.Schema

    schema "users" do
      field :email, :string

      field :full_name, :string, virtual: true

      timestamps()
    end

    def resolve_full_name(%Ectotest.User{id: id, email: email} = user) do
        %Ectotest.User{full_name: id <> " " <> email}
    end

    def get_user_by_id(id) do
        Ectotest.User
        |> Ectotest.Repo.get(id)
        |> case do
          %Ectotest.User{} = user ->
            {:ok, Ectotest.User.resolve_full_name(user)}
          nil ->
            {:error, :not_found}
        end
      end
end
