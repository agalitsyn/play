import Ecto.Query
alias Ecto.Adapter.SQL
alias Ectotest.Repo
alias Ectotest.User

to_insert = [
    [
        email: "a@b.com",
        inserted_at: DateTime.utc_now(),
        updated_at: DateTime.utc_now(),
    ],
    [
        email: "b@b.com",
        inserted_at: DateTime.utc_now(),
        updated_at: DateTime.utc_now(),
    ],
]

Repo.insert_all "users", to_insert, returning: [:id, :email]

Repo.query("select email from users")
Repo.one(from u in "users", where: u.email == "a@b.com", select: u.id)
Repo.all(from u in "users", select: [u.id, u.email, u.inserted_at])
