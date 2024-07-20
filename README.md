# 2024-07-20 sqlc codegen bug

TODO: put issue link here

Build:

```bash
go generate ./...
```

Run `using-text`, which works as expected

```bash
$ go run ./using-text
create...
sqlcgen.Env{ID:1, CreateTime:"0001-01-01T00:00:00Z"}

empty update...
sqlcgen.Env{ID:1, CreateTime:"0001-01-01T00:00:00Z"}

real update...
sqlcgen.Env{ID:1, CreateTime:"0001-01-01T11:11:11Z"}
```

Run `using-override`, which is the same as `using-text`, except using a `type SQLiteTime string` as a custom type override:

```bash
$ go run ./using-override
create...
sqlcgen.Env{ID:1, CreateTime:"0001-01-01T00:00:00Z"}

empty update...
sqlcgen.Env{ID:1, CreateTime:"???"}

real update...
sqlcgen.Env{ID:1, CreateTime:"0001-01-01T11:11:11Z"}
```

As the type generated is not a pointer, you can't update it with NULL.