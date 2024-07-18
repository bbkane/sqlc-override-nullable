package main

import "go.bbkane.com/sqlc-override-nullable/sqlite"

func main() {
	_, err := sqlite.Connect("tmp.db")
	if err != nil {
		panic(err)
	}
}
