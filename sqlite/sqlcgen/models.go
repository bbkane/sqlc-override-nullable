// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package sqlcgen

type Env struct {
	EnvID      int64
	Name       string
	Comment    string
	CreateTime string
	UpdateTime string
}

type EnvRef struct {
	EnvRefID   int64
	EnvID      int64
	Name       string
	Comment    string
	CreateTime string
	UpdateTime string
	EnvVarID   int64
}

type EnvVar struct {
	EnvVarID   int64
	EnvID      int64
	Name       string
	Comment    string
	CreateTime string
	UpdateTime string
	Value      string
}

type KeyringEntry struct {
	KeyringEntryID int64
	Name           string
	Comment        string
	CreateTime     string
	UpdateTime     string
}

type VwEnvEnvRefEnvVarUniqueName struct {
	EnvID int64
	Name  string
}

type VwEnvRefReferencedName struct {
	EnvRefID   int64
	EnvID      int64
	EnvName    string
	Name       string
	EnvVarID   int64
	RefVarName string
	RefEnvName string
	Comment    string
	CreateTime string
	UpdateTime string
}

type VwEnvVarReferencedName struct {
	EnvVarID   int64
	EnvID      int64
	EnvName    string
	Name       string
	Value      string
	Comment    string
	CreateTime string
	UpdateTime string
}