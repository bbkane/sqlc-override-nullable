// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package sqlcgen

import (
	"go.bbkane.com/sqlc-override-nullable/using-override/sqlite"
)

type Env struct {
	ID         int64
	CreateTime sqlite.SQLiteTime
}
