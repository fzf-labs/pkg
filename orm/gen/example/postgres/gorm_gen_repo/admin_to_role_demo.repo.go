// Code generated by gen/repo. DO NOT EDIT.
// Code generated by gen/repo. DO NOT EDIT.
// Code generated by gen/repo. DO NOT EDIT.

package gorm_gen_repo

import (
	"context"
	"errors"

	"github.com/fzf-labs/fpkg/orm/custom"
	"github.com/fzf-labs/fpkg/orm/gen/cache"
	"github.com/fzf-labs/fpkg/orm/gen/example/postgres/gorm_gen_dao"
	"github.com/fzf-labs/fpkg/orm/gen/example/postgres/gorm_gen_model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ IAdminToRoleDemoRepo = (*AdminToRoleDemoRepo)(nil)

type (
	IAdminToRoleDemoRepo interface {
		// CreateOne 创建一条数据
		CreateOne(ctx context.Context, data *gorm_gen_model.AdminToRoleDemo) error
		// CreateOneByTx 创建一条数据(事务)
		CreateOneByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminToRoleDemo) error
		// UpsertOne Upsert一条数据
		UpsertOne(ctx context.Context, data *gorm_gen_model.AdminToRoleDemo) error
		// UpsertOneByTx Upsert一条数据(事务)
		UpsertOneByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminToRoleDemo) error
		// UpsertOneByFields Upsert一条数据，根据fields字段
		UpsertOneByFields(ctx context.Context, data *gorm_gen_model.AdminToRoleDemo, fields []string) error
		// UpsertOneByFieldsTx Upsert一条数据，根据fields字段(事务)
		UpsertOneByFieldsTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminToRoleDemo, fields []string) error
		// CreateBatch 批量创建数据
		CreateBatch(ctx context.Context, data []*gorm_gen_model.AdminToRoleDemo, batchSize int) error
		// FindMultiByCustom 自定义查询数据(通用)
		FindMultiByCustom(ctx context.Context, customReq *custom.PaginatorReq) ([]*gorm_gen_model.AdminToRoleDemo, *custom.PaginatorReply, error)
	}
	AdminToRoleDemoRepo struct {
		db    *gorm.DB
		cache cache.IDBCache
	}
)

func NewAdminToRoleDemoRepo(db *gorm.DB, cache cache.IDBCache) *AdminToRoleDemoRepo {
	return &AdminToRoleDemoRepo{
		db:    db,
		cache: cache,
	}
}

// CreateOne 创建一条数据
func (a *AdminToRoleDemoRepo) CreateOne(ctx context.Context, data *gorm_gen_model.AdminToRoleDemo) error {
	dao := gorm_gen_dao.Use(a.db).AdminToRoleDemo
	err := dao.WithContext(ctx).Create(data)
	if err != nil {
		return err
	}
	return nil
}

// CreateOneByTx 创建一条数据(事务)
func (a *AdminToRoleDemoRepo) CreateOneByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminToRoleDemo) error {
	dao := tx.AdminToRoleDemo
	err := dao.WithContext(ctx).Create(data)
	if err != nil {
		return err
	}
	return nil
}

// UpsertOne Upsert一条数据
func (a *AdminToRoleDemoRepo) UpsertOne(ctx context.Context, data *gorm_gen_model.AdminToRoleDemo) error {
	dao := gorm_gen_dao.Use(a.db).AdminToRoleDemo
	err := dao.WithContext(ctx).Save(data)
	if err != nil {
		return err
	}
	return nil
}

// UpsertOneByTx Upsert一条数据(事务)
func (a *AdminToRoleDemoRepo) UpsertOneByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminToRoleDemo) error {
	dao := tx.AdminToRoleDemo
	err := dao.WithContext(ctx).Save(data)
	if err != nil {
		return err
	}
	return nil
}

// UpsertOneByFields Upsert一条数据，根据fields字段
func (a *AdminToRoleDemoRepo) UpsertOneByFields(ctx context.Context, data *gorm_gen_model.AdminToRoleDemo, fields []string) error {
	if len(fields) == 0 {
		return errors.New("UpsertOneByFields fields is empty")
	}
	columns := make([]clause.Column, 0)
	for _, v := range fields {
		columns = append(columns, clause.Column{Name: v})
	}
	dao := gorm_gen_dao.Use(a.db).AdminToRoleDemo
	err := dao.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   columns,
		UpdateAll: true,
	}).Create(data)
	if err != nil {
		return err
	}
	return nil
}

// UpsertOneByFieldsTx Upsert一条数据，根据fields字段(事务)
func (a *AdminToRoleDemoRepo) UpsertOneByFieldsTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminToRoleDemo, fields []string) error {
	if len(fields) == 0 {
		return errors.New("UpsertOneByFieldsTx fields is empty")
	}
	columns := make([]clause.Column, 0)
	for _, v := range fields {
		columns = append(columns, clause.Column{Name: v})
	}
	dao := tx.AdminToRoleDemo
	err := dao.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   columns,
		UpdateAll: true,
	}).Create(data)
	if err != nil {
		return err
	}
	return nil
}

// CreateBatch 批量创建数据
func (a *AdminToRoleDemoRepo) CreateBatch(ctx context.Context, data []*gorm_gen_model.AdminToRoleDemo, batchSize int) error {
	dao := gorm_gen_dao.Use(a.db).AdminToRoleDemo
	err := dao.WithContext(ctx).CreateInBatches(data, batchSize)
	if err != nil {
		return err
	}
	return nil
}

// FindMultiByCustom 自定义查询数据(通用)
func (a *AdminToRoleDemoRepo) FindMultiByCustom(ctx context.Context, customReq *custom.PaginatorReq) ([]*gorm_gen_model.AdminToRoleDemo, *custom.PaginatorReply, error) {
	result := make([]*gorm_gen_model.AdminToRoleDemo, 0)
	var total int64
	whereExpressions, orderExpressions, err := customReq.ConvertToGormExpression(gorm_gen_model.AdminToRoleDemo{})
	if err != nil {
		return result, nil, err
	}
	err = a.db.WithContext(ctx).Model(&gorm_gen_model.AdminToRoleDemo{}).Select([]string{"*"}).Clauses(whereExpressions...).Count(&total).Error
	if err != nil {
		return result, nil, err
	}
	if total == 0 {
		return result, nil, nil
	}
	customReply, err := customReq.ConvertToPage(int32(total))
	if err != nil {
		return result, nil, err
	}
	query := a.db.WithContext(ctx).Model(&gorm_gen_model.AdminToRoleDemo{}).Clauses(whereExpressions...).Clauses(orderExpressions...)
	if customReply.Page != 0 && customReply.PageSize != 0 {
		query = query.Offset(int((customReply.Page - 1) * customReply.PageSize))
		query = query.Limit(int(customReply.PageSize))
	}
	err = query.Find(&result).Error
	if err != nil {
		return result, nil, err
	}
	return result, customReply, err
}
