// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gorm_gen_dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/fzf-labs/fpkg/orm/gen/example/postgres/gorm_gen_model"
)

func newAdminToRoleDemo(db *gorm.DB, opts ...gen.DOOption) adminToRoleDemo {
	_adminToRoleDemo := adminToRoleDemo{}

	_adminToRoleDemo.adminToRoleDemoDo.UseDB(db, opts...)
	_adminToRoleDemo.adminToRoleDemoDo.UseModel(&gorm_gen_model.AdminToRoleDemo{})

	tableName := _adminToRoleDemo.adminToRoleDemoDo.TableName()
	_adminToRoleDemo.ALL = field.NewAsterisk(tableName)
	_adminToRoleDemo.AdminID = field.NewString(tableName, "admin_id")
	_adminToRoleDemo.RoleID = field.NewString(tableName, "role_id")

	_adminToRoleDemo.fillFieldMap()

	return _adminToRoleDemo
}

type adminToRoleDemo struct {
	adminToRoleDemoDo adminToRoleDemoDo

	ALL     field.Asterisk
	AdminID field.String
	RoleID  field.String

	fieldMap map[string]field.Expr
}

func (a adminToRoleDemo) Table(newTableName string) *adminToRoleDemo {
	a.adminToRoleDemoDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a adminToRoleDemo) As(alias string) *adminToRoleDemo {
	a.adminToRoleDemoDo.DO = *(a.adminToRoleDemoDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *adminToRoleDemo) updateTableName(table string) *adminToRoleDemo {
	a.ALL = field.NewAsterisk(table)
	a.AdminID = field.NewString(table, "admin_id")
	a.RoleID = field.NewString(table, "role_id")

	a.fillFieldMap()

	return a
}

func (a *adminToRoleDemo) WithContext(ctx context.Context) *adminToRoleDemoDo {
	return a.adminToRoleDemoDo.WithContext(ctx)
}

func (a adminToRoleDemo) TableName() string { return a.adminToRoleDemoDo.TableName() }

func (a adminToRoleDemo) Alias() string { return a.adminToRoleDemoDo.Alias() }

func (a adminToRoleDemo) Columns(cols ...field.Expr) gen.Columns {
	return a.adminToRoleDemoDo.Columns(cols...)
}

func (a *adminToRoleDemo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *adminToRoleDemo) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 2)
	a.fieldMap["admin_id"] = a.AdminID
	a.fieldMap["role_id"] = a.RoleID
}

func (a adminToRoleDemo) clone(db *gorm.DB) adminToRoleDemo {
	a.adminToRoleDemoDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a adminToRoleDemo) replaceDB(db *gorm.DB) adminToRoleDemo {
	a.adminToRoleDemoDo.ReplaceDB(db)
	return a
}

type adminToRoleDemoDo struct{ gen.DO }

func (a adminToRoleDemoDo) Debug() *adminToRoleDemoDo {
	return a.withDO(a.DO.Debug())
}

func (a adminToRoleDemoDo) WithContext(ctx context.Context) *adminToRoleDemoDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a adminToRoleDemoDo) ReadDB() *adminToRoleDemoDo {
	return a.Clauses(dbresolver.Read)
}

func (a adminToRoleDemoDo) WriteDB() *adminToRoleDemoDo {
	return a.Clauses(dbresolver.Write)
}

func (a adminToRoleDemoDo) Session(config *gorm.Session) *adminToRoleDemoDo {
	return a.withDO(a.DO.Session(config))
}

func (a adminToRoleDemoDo) Clauses(conds ...clause.Expression) *adminToRoleDemoDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a adminToRoleDemoDo) Returning(value interface{}, columns ...string) *adminToRoleDemoDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a adminToRoleDemoDo) Not(conds ...gen.Condition) *adminToRoleDemoDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a adminToRoleDemoDo) Or(conds ...gen.Condition) *adminToRoleDemoDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a adminToRoleDemoDo) Select(conds ...field.Expr) *adminToRoleDemoDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a adminToRoleDemoDo) Where(conds ...gen.Condition) *adminToRoleDemoDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a adminToRoleDemoDo) Order(conds ...field.Expr) *adminToRoleDemoDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a adminToRoleDemoDo) Distinct(cols ...field.Expr) *adminToRoleDemoDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a adminToRoleDemoDo) Omit(cols ...field.Expr) *adminToRoleDemoDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a adminToRoleDemoDo) Join(table schema.Tabler, on ...field.Expr) *adminToRoleDemoDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a adminToRoleDemoDo) LeftJoin(table schema.Tabler, on ...field.Expr) *adminToRoleDemoDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a adminToRoleDemoDo) RightJoin(table schema.Tabler, on ...field.Expr) *adminToRoleDemoDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a adminToRoleDemoDo) Group(cols ...field.Expr) *adminToRoleDemoDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a adminToRoleDemoDo) Having(conds ...gen.Condition) *adminToRoleDemoDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a adminToRoleDemoDo) Limit(limit int) *adminToRoleDemoDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a adminToRoleDemoDo) Offset(offset int) *adminToRoleDemoDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a adminToRoleDemoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *adminToRoleDemoDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a adminToRoleDemoDo) Unscoped() *adminToRoleDemoDo {
	return a.withDO(a.DO.Unscoped())
}

func (a adminToRoleDemoDo) Create(values ...*gorm_gen_model.AdminToRoleDemo) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a adminToRoleDemoDo) CreateInBatches(values []*gorm_gen_model.AdminToRoleDemo, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a adminToRoleDemoDo) Save(values ...*gorm_gen_model.AdminToRoleDemo) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a adminToRoleDemoDo) First() (*gorm_gen_model.AdminToRoleDemo, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*gorm_gen_model.AdminToRoleDemo), nil
	}
}

func (a adminToRoleDemoDo) Take() (*gorm_gen_model.AdminToRoleDemo, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*gorm_gen_model.AdminToRoleDemo), nil
	}
}

func (a adminToRoleDemoDo) Last() (*gorm_gen_model.AdminToRoleDemo, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*gorm_gen_model.AdminToRoleDemo), nil
	}
}

func (a adminToRoleDemoDo) Find() ([]*gorm_gen_model.AdminToRoleDemo, error) {
	result, err := a.DO.Find()
	return result.([]*gorm_gen_model.AdminToRoleDemo), err
}

func (a adminToRoleDemoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*gorm_gen_model.AdminToRoleDemo, err error) {
	buf := make([]*gorm_gen_model.AdminToRoleDemo, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a adminToRoleDemoDo) FindInBatches(result *[]*gorm_gen_model.AdminToRoleDemo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a adminToRoleDemoDo) Attrs(attrs ...field.AssignExpr) *adminToRoleDemoDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a adminToRoleDemoDo) Assign(attrs ...field.AssignExpr) *adminToRoleDemoDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a adminToRoleDemoDo) Joins(fields ...field.RelationField) *adminToRoleDemoDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a adminToRoleDemoDo) Preload(fields ...field.RelationField) *adminToRoleDemoDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a adminToRoleDemoDo) FirstOrInit() (*gorm_gen_model.AdminToRoleDemo, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*gorm_gen_model.AdminToRoleDemo), nil
	}
}

func (a adminToRoleDemoDo) FirstOrCreate() (*gorm_gen_model.AdminToRoleDemo, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*gorm_gen_model.AdminToRoleDemo), nil
	}
}

func (a adminToRoleDemoDo) FindByPage(offset int, limit int) (result []*gorm_gen_model.AdminToRoleDemo, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a adminToRoleDemoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a adminToRoleDemoDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a adminToRoleDemoDo) Delete(models ...*gorm_gen_model.AdminToRoleDemo) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *adminToRoleDemoDo) withDO(do gen.Dao) *adminToRoleDemoDo {
	a.DO = *do.(*gen.DO)
	return a
}