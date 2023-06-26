package repo

import (
	_ "embed"
)

//go:embed tpl/pkg.tpl
var Pkg string

//go:embed tpl/import.tpl
var Import string

//go:embed tpl/var.tpl
var Vars string

//go:embed tpl/types.tpl
var Types string

// TableRepo 表结构
type TableRepo struct {
	dbName         string
	lowerDbName    string
	upperDbName    string
	tableName      string
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

type FieldRepo struct {
	lowerField        string //字段小写
	upperField        string //大写字段
	lowerFieldComplex string //字段小写复数
	upperFieldComplex string //字段大写复数
	dataType          string //字段类型
}
