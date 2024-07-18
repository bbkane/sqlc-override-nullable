package main

import (
	"context"
	"fmt"
	"os"

	"go.bbkane.com/sqlc-override-nullable/sqlite"
	"go.bbkane.com/sqlc-override-nullable/sqlite/sqlcgen"
)

const zeroTime = "0001-01-01T00:00:00Z"
const oneTime = "0001-01-01T01:00:00Z"

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	os.Remove("tmp.db")
	db, err := sqlite.Connect("tmp.db")
	panicIf(err)

	queries := sqlcgen.New(db)
	ctx := context.Background()

	// create
	_, err = queries.EnvCreate(ctx, zeroTime)
	panicIf(err)

	fmt.Println("create...")

	// list
	{
		envs, err := queries.EnvList(ctx)
		panicIf(err)
		for _, env := range envs {
			fmt.Printf("%#v\n", env)
		}
	}

	fmt.Println("empty update...")

	// empty update
	err = queries.EnvUpdate(ctx, sqlcgen.EnvUpdateParams{
		CreateTime: nil,
	})
	panicIf(err)

	// list
	{
		envs, err := queries.EnvList(ctx)
		panicIf(err)
		for _, env := range envs {
			fmt.Printf("%#v\n", env)
		}
	}

	fmt.Println("real update...")

	// real update
	oneTime := oneTime
	err = queries.EnvUpdate(ctx, sqlcgen.EnvUpdateParams{
		CreateTime: &oneTime,
	})
	panicIf(err)

	// list
	{
		envs, err := queries.EnvList(ctx)
		panicIf(err)
		for _, env := range envs {
			fmt.Printf("%#v\n", env)
		}
	}
}
