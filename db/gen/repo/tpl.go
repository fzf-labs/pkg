package repo

import (
	_ "embed"
)

//go:embed tpl/pkg.tpl
var Pkg string

//go:embed tpl/import.tpl
var Imports string

//go:embed tpl/var.tpl
var Vars string

//go:embed tpl/types.tpl
var Types string

//go:embed tpl/find-one-by-field.tpl
var FindOneByField string

//go:embed tpl/interface-find-one-by-field.tpl
var FindOneByFieldMethod string

type Repo struct {
	lowerDbName    string
	upperDbName    string
	lowerTableName string
	upperTableName string
	cacheKeys      []string
	methods        []string
}

type FuncRepo struct {
	lowerField string
	upperField string
	in         string
}
