// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package user_dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/fzf-labs/fpkg/db/gen/example/postgres/user_model"
)

func newUserDemo(db *gorm.DB, opts ...gen.DOOption) userDemo {
	_userDemo := userDemo{}

	_userDemo.userDemoDo.UseDB(db, opts...)
	_userDemo.userDemoDo.UseModel(&user_model.UserDemo{})

	tableName := _userDemo.userDemoDo.TableName()
	_userDemo.ALL = field.NewAsterisk(tableName)
	_userDemo.ID = field.NewInt64(tableName, "id")
	_userDemo.UID = field.NewString(tableName, "uid")
	_userDemo.Username = field.NewString(tableName, "username")
	_userDemo.Password = field.NewString(tableName, "password")
	_userDemo.Nickname = field.NewString(tableName, "nickname")
	_userDemo.Remark = field.NewString(tableName, "remark")
	_userDemo.DeptID = field.NewInt64(tableName, "dept_id")
	_userDemo.PostIds = field.NewString(tableName, "post_ids")
	_userDemo.Email = field.NewString(tableName, "email")
	_userDemo.Mobile = field.NewString(tableName, "mobile")
	_userDemo.Sex = field.NewInt16(tableName, "sex")
	_userDemo.Avatar = field.NewString(tableName, "avatar")
	_userDemo.Status = field.NewInt16(tableName, "status")
	_userDemo.LoginIP = field.NewString(tableName, "login_ip")
	_userDemo.LoginDate = field.NewField(tableName, "login_date")
	_userDemo.TenantID = field.NewInt64(tableName, "tenant_id")
	_userDemo.CreatedAt = field.NewTime(tableName, "created_at")
	_userDemo.UpdatedAt = field.NewTime(tableName, "updated_at")
	_userDemo.DeletedAt = field.NewField(tableName, "deleted_at")

	_userDemo.fillFieldMap()

	return _userDemo
}

type userDemo struct {
	userDemoDo userDemoDo

	ALL       field.Asterisk
	ID        field.Int64  // ID
	UID       field.String // uid
	Username  field.String // 用户账号
	Password  field.String // 密码
	Nickname  field.String // 用户昵称
	Remark    field.String // 备注
	DeptID    field.Int64  // 部门ID
	PostIds   field.String // 岗位编号数组
	Email     field.String // 用户邮箱
	Mobile    field.String // 手机号码
	Sex       field.Int16  // 用户性别
	Avatar    field.String // 头像地址
	Status    field.Int16  // 帐号状态（0正常 -1停用）
	LoginIP   field.String // 最后登录IP
	LoginDate field.Field  // 最后登录时间
	TenantID  field.Int64  // 租户编号
	CreatedAt field.Time   // 创建时间
	UpdatedAt field.Time   // 更新时间
	DeletedAt field.Field  // 删除时间

	fieldMap map[string]field.Expr
}

func (u userDemo) Table(newTableName string) *userDemo {
	u.userDemoDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userDemo) As(alias string) *userDemo {
	u.userDemoDo.DO = *(u.userDemoDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userDemo) updateTableName(table string) *userDemo {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.UID = field.NewString(table, "uid")
	u.Username = field.NewString(table, "username")
	u.Password = field.NewString(table, "password")
	u.Nickname = field.NewString(table, "nickname")
	u.Remark = field.NewString(table, "remark")
	u.DeptID = field.NewInt64(table, "dept_id")
	u.PostIds = field.NewString(table, "post_ids")
	u.Email = field.NewString(table, "email")
	u.Mobile = field.NewString(table, "mobile")
	u.Sex = field.NewInt16(table, "sex")
	u.Avatar = field.NewString(table, "avatar")
	u.Status = field.NewInt16(table, "status")
	u.LoginIP = field.NewString(table, "login_ip")
	u.LoginDate = field.NewField(table, "login_date")
	u.TenantID = field.NewInt64(table, "tenant_id")
	u.CreatedAt = field.NewTime(table, "created_at")
	u.UpdatedAt = field.NewTime(table, "updated_at")
	u.DeletedAt = field.NewField(table, "deleted_at")

	u.fillFieldMap()

	return u
}

func (u *userDemo) WithContext(ctx context.Context) *userDemoDo { return u.userDemoDo.WithContext(ctx) }

func (u userDemo) TableName() string { return u.userDemoDo.TableName() }

func (u userDemo) Alias() string { return u.userDemoDo.Alias() }

func (u userDemo) Columns(cols ...field.Expr) gen.Columns { return u.userDemoDo.Columns(cols...) }

func (u *userDemo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userDemo) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 19)
	u.fieldMap["id"] = u.ID
	u.fieldMap["uid"] = u.UID
	u.fieldMap["username"] = u.Username
	u.fieldMap["password"] = u.Password
	u.fieldMap["nickname"] = u.Nickname
	u.fieldMap["remark"] = u.Remark
	u.fieldMap["dept_id"] = u.DeptID
	u.fieldMap["post_ids"] = u.PostIds
	u.fieldMap["email"] = u.Email
	u.fieldMap["mobile"] = u.Mobile
	u.fieldMap["sex"] = u.Sex
	u.fieldMap["avatar"] = u.Avatar
	u.fieldMap["status"] = u.Status
	u.fieldMap["login_ip"] = u.LoginIP
	u.fieldMap["login_date"] = u.LoginDate
	u.fieldMap["tenant_id"] = u.TenantID
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["deleted_at"] = u.DeletedAt
}

func (u userDemo) clone(db *gorm.DB) userDemo {
	u.userDemoDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userDemo) replaceDB(db *gorm.DB) userDemo {
	u.userDemoDo.ReplaceDB(db)
	return u
}

type userDemoDo struct{ gen.DO }

func (u userDemoDo) Debug() *userDemoDo {
	return u.withDO(u.DO.Debug())
}

func (u userDemoDo) WithContext(ctx context.Context) *userDemoDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userDemoDo) ReadDB() *userDemoDo {
	return u.Clauses(dbresolver.Read)
}

func (u userDemoDo) WriteDB() *userDemoDo {
	return u.Clauses(dbresolver.Write)
}

func (u userDemoDo) Session(config *gorm.Session) *userDemoDo {
	return u.withDO(u.DO.Session(config))
}

func (u userDemoDo) Clauses(conds ...clause.Expression) *userDemoDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userDemoDo) Returning(value interface{}, columns ...string) *userDemoDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userDemoDo) Not(conds ...gen.Condition) *userDemoDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userDemoDo) Or(conds ...gen.Condition) *userDemoDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userDemoDo) Select(conds ...field.Expr) *userDemoDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userDemoDo) Where(conds ...gen.Condition) *userDemoDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userDemoDo) Order(conds ...field.Expr) *userDemoDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userDemoDo) Distinct(cols ...field.Expr) *userDemoDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userDemoDo) Omit(cols ...field.Expr) *userDemoDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userDemoDo) Join(table schema.Tabler, on ...field.Expr) *userDemoDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userDemoDo) LeftJoin(table schema.Tabler, on ...field.Expr) *userDemoDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userDemoDo) RightJoin(table schema.Tabler, on ...field.Expr) *userDemoDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userDemoDo) Group(cols ...field.Expr) *userDemoDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userDemoDo) Having(conds ...gen.Condition) *userDemoDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userDemoDo) Limit(limit int) *userDemoDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userDemoDo) Offset(offset int) *userDemoDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userDemoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *userDemoDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userDemoDo) Unscoped() *userDemoDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userDemoDo) Create(values ...*user_model.UserDemo) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userDemoDo) CreateInBatches(values []*user_model.UserDemo, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userDemoDo) Save(values ...*user_model.UserDemo) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userDemoDo) First() (*user_model.UserDemo, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*user_model.UserDemo), nil
	}
}

func (u userDemoDo) Take() (*user_model.UserDemo, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*user_model.UserDemo), nil
	}
}

func (u userDemoDo) Last() (*user_model.UserDemo, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*user_model.UserDemo), nil
	}
}

func (u userDemoDo) Find() ([]*user_model.UserDemo, error) {
	result, err := u.DO.Find()
	return result.([]*user_model.UserDemo), err
}

func (u userDemoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*user_model.UserDemo, err error) {
	buf := make([]*user_model.UserDemo, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userDemoDo) FindInBatches(result *[]*user_model.UserDemo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userDemoDo) Attrs(attrs ...field.AssignExpr) *userDemoDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userDemoDo) Assign(attrs ...field.AssignExpr) *userDemoDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userDemoDo) Joins(fields ...field.RelationField) *userDemoDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userDemoDo) Preload(fields ...field.RelationField) *userDemoDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userDemoDo) FirstOrInit() (*user_model.UserDemo, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*user_model.UserDemo), nil
	}
}

func (u userDemoDo) FirstOrCreate() (*user_model.UserDemo, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*user_model.UserDemo), nil
	}
}

func (u userDemoDo) FindByPage(offset int, limit int) (result []*user_model.UserDemo, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userDemoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userDemoDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userDemoDo) Delete(models ...*user_model.UserDemo) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userDemoDo) withDO(do gen.Dao) *userDemoDo {
	u.DO = *do.(*gen.DO)
	return u
}
