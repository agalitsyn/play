package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgxutil"
)

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://postgres:postgres@localhost:5432/postgres")
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	if err := sample(ctx, conn); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

type Foo struct {
	ID       int64
	Title    string
	Subtitle string
}

func sample(ctx context.Context, db *pgx.Conn) error {
	q := `
DROP TABLE foo;

CREATE TABLE foo (
	id serial primary key,
	title text not null,
	subtitle text not null default ''
);

INSERT INTO foo (title) VALUES ('foo'), ('bar'), ('baz');
`
	_, err := db.Exec(ctx, q)
	if err != nil {
		return err
	}

	q = `SELECT id, title from foo;`
	rows, err := db.Query(ctx, q)
	if err != nil {
		return fmt.Errorf("cound not perform query: %w", err)
	}
	foos, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Foo])
	if err != nil {
		return fmt.Errorf("CollectRows fail: %w", err)
	}
	fmt.Printf("(foos): %#v\n", foos)

	q = `SELECT * from foo;`
	foos2, err := pgxutil.Select(ctx, db, q, nil, pgx.RowToStructByPos[Foo])
	if err != nil {
		return err
	}
	fmt.Printf("(foos2): %#v\n", foos2)

	return nil
}
