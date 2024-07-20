package main

import (
	"context"
	"fmt"
	"os"

	"go.bbkane.com/sqlc-override-nullable/using-override/sqlite"
	"go.bbkane.com/sqlc-override-nullable/using-override/sqlite/sqlcgen"
)

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func printEnv(ctx context.Context, title string, queries *sqlcgen.Queries, id int64) {
	env, err := queries.EnvGet(ctx, id)
	panicIf(err)
	fmt.Printf("%s\n%#v\n\n", title, env)
}

func main() {

	zeroTime := "0001-01-01T00:00:00Z"
	oneTime := "0001-01-01T11:11:11Z"
	var id int64 = 1

	os.Remove("tmp.db")
	db, err := sqlite.Connect("tmp.db")
	panicIf(err)

	queries := sqlcgen.New(db)
	ctx := context.Background()

	// create
	_, err = queries.EnvCreate(ctx, sqlcgen.EnvCreateParams{
		ID:         int64(id),
		CreateTime: sqlite.SQLiteTime(zeroTime),
	})
	panicIf(err)
	printEnv(ctx, "create...", queries, id)

	// empty update
	err = queries.EnvUpdate(ctx, sqlcgen.EnvUpdateParams{
		CreateTime: sqlite.SQLiteTime("???"), // BUG: this should be a pointer type
		ID:         id,
	})
	panicIf(err)
	printEnv(ctx, "empty update...", queries, id)

	// real update
	err = queries.EnvUpdate(ctx, sqlcgen.EnvUpdateParams{
		CreateTime: sqlite.SQLiteTime(oneTime), // BUG: this should be a pointer type
		ID:         id,
	})
	panicIf(err)
	printEnv(ctx, "real update...", queries, id)

}
