// Code generated by gen/repo. DO NOT EDIT.
// Code generated by gen/repo. DO NOT EDIT.
// Code generated by gen/repo. DO NOT EDIT.

package gorm_gen_repo

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/fzf-labs/fpkg/orm/gen/cache"
	"github.com/fzf-labs/fpkg/orm/gen/example/postgres/gorm_gen_dao"
	"github.com/fzf-labs/fpkg/orm/gen/example/postgres/gorm_gen_model"
	"github.com/fzf-labs/fpkg/orm/paginator"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ IUserDemoRepo = (*UserDemoRepo)(nil)

var (
	cacheUserDemoByIDPrefix        = "DBCache:gorm_gen:UserDemoByID"
	cacheUserDemoByUIDPrefix       = "DBCache:gorm_gen:UserDemoByUID"
	cacheUserDemoByUIDStatusPrefix = "DBCache:gorm_gen:UserDemoByUIDStatus"
)

type (
	IUserDemoRepo interface {
		// CreateOne 创建一条数据
		CreateOne(ctx context.Context, data *gorm_gen_model.UserDemo) error
		// CreateOneByTx 创建一条数据(事务)
		CreateOneByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.UserDemo) error
		// UpsertOne Upsert一条数据
		UpsertOne(ctx context.Context, data *gorm_gen_model.UserDemo) error
		// UpsertOneByTx Upsert一条数据(事务)
		UpsertOneByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.UserDemo) error
		// UpsertOneByFields Upsert一条数据，根据fields字段
		UpsertOneByFields(ctx context.Context, data *gorm_gen_model.UserDemo, fields []string) error
		// UpsertOneByFieldsTx Upsert一条数据，根据fields字段(事务)
		UpsertOneByFieldsTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.UserDemo, fields []string) error
		// CreateBatch 批量创建数据
		CreateBatch(ctx context.Context, data []*gorm_gen_model.UserDemo, batchSize int) error
		// UpdateOne 更新一条数据
		UpdateOne(ctx context.Context, data *gorm_gen_model.UserDemo) error
		// UpdateOne 更新一条数据(事务)
		UpdateOneByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.UserDemo) error
		// UpdateOneWithZero 更新一条数据,包含零值
		UpdateOneWithZero(ctx context.Context, data *gorm_gen_model.UserDemo) error
		// UpdateOneWithZero 更新一条数据,包含零值(事务)
		UpdateOneWithZeroByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.UserDemo) error
		// FindOneCacheByID 根据ID查询一条数据并设置缓存
		FindOneCacheByID(ctx context.Context, ID int64) (*gorm_gen_model.UserDemo, error)
		// FindOneByID 根据ID查询一条数据
		FindOneByID(ctx context.Context, ID int64) (*gorm_gen_model.UserDemo, error)
		// FindMultiCacheByIDS 根据IDS查询多条数据并设置缓存
		FindMultiCacheByIDS(ctx context.Context, IDS []int64) ([]*gorm_gen_model.UserDemo, error)
		// FindMultiByIDS 根据IDS查询多条数据
		FindMultiByIDS(ctx context.Context, IDS []int64) ([]*gorm_gen_model.UserDemo, error)
		// FindOneCacheByUID 根据UID查询一条数据并设置缓存
		FindOneCacheByUID(ctx context.Context, UID string) (*gorm_gen_model.UserDemo, error)
		// FindOneByUID 根据UID查询一条数据
		FindOneByUID(ctx context.Context, UID string) (*gorm_gen_model.UserDemo, error)
		// FindMultiCacheByUIDS 根据UIDS查询多条数据并设置缓存
		FindMultiCacheByUIDS(ctx context.Context, UIDS []string) ([]*gorm_gen_model.UserDemo, error)
		// FindMultiByUIDS 根据UIDS查询多条数据
		FindMultiByUIDS(ctx context.Context, UIDS []string) ([]*gorm_gen_model.UserDemo, error)
		// FindOneCacheByUIDStatus 根据UIDStatus查询一条数据并设置缓存
		FindOneCacheByUIDStatus(ctx context.Context, UID string, status int16) (*gorm_gen_model.UserDemo, error)
		// FindOneByUIDStatus 根据UIDStatus查询一条数据
		FindOneByUIDStatus(ctx context.Context, UID string, status int16) (*gorm_gen_model.UserDemo, error)
		// FindMultiByTenantIDDeptID 根据TenantIDDeptID查询多条数据
		FindMultiByTenantIDDeptID(ctx context.Context, tenantID int64, deptID int64) ([]*gorm_gen_model.UserDemo, error)
		// FindMultiByUsername 根据username查询多条数据
		FindMultiByUsername(ctx context.Context, username string) ([]*gorm_gen_model.UserDemo, error)
		// FindMultiByUsernames 根据usernames查询多条数据
		FindMultiByUsernames(ctx context.Context, usernames []string) ([]*gorm_gen_model.UserDemo, error)
		// FindMultiByTenantID 根据tenantID查询多条数据
		FindMultiByTenantID(ctx context.Context, tenantID int64) ([]*gorm_gen_model.UserDemo, error)
		// FindMultiByTenantIDS 根据tenantIDS查询多条数据
		FindMultiByTenantIDS(ctx context.Context, tenantIDS []int64) ([]*gorm_gen_model.UserDemo, error)
		// FindMultiByPaginator 查询分页数据(通用)
		FindMultiByPaginator(ctx context.Context, paginatorReq *paginator.Req) ([]*gorm_gen_model.UserDemo, *paginator.Reply, error)
		// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
		DeleteOneCacheByID(ctx context.Context, ID int64) error
		// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
		DeleteOneCacheByIDTx(ctx context.Context, tx *gorm_gen_dao.Query, ID int64) error
		// DeleteOneByID 根据ID删除一条数据
		DeleteOneByID(ctx context.Context, ID int64) error
		// DeleteOneByID 根据ID删除一条数据
		DeleteOneByIDTx(ctx context.Context, tx *gorm_gen_dao.Query, ID int64) error
		// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
		DeleteMultiCacheByIDS(ctx context.Context, IDS []int64) error
		// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
		DeleteMultiCacheByIDSTx(ctx context.Context, tx *gorm_gen_dao.Query, IDS []int64) error
		// DeleteMultiByIDS 根据IDS删除多条数据
		DeleteMultiByIDS(ctx context.Context, IDS []int64) error
		// DeleteMultiByIDS 根据IDS删除多条数据
		DeleteMultiByIDSTx(ctx context.Context, tx *gorm_gen_dao.Query, IDS []int64) error
		// DeleteOneCacheByUID 根据UID删除一条数据并清理缓存
		DeleteOneCacheByUID(ctx context.Context, UID string) error
		// DeleteOneCacheByUID 根据UID删除一条数据并清理缓存
		DeleteOneCacheByUIDTx(ctx context.Context, tx *gorm_gen_dao.Query, UID string) error
		// DeleteOneByUID 根据UID删除一条数据
		DeleteOneByUID(ctx context.Context, UID string) error
		// DeleteOneByUID 根据UID删除一条数据
		DeleteOneByUIDTx(ctx context.Context, tx *gorm_gen_dao.Query, UID string) error
		// DeleteMultiCacheByUIDS 根据UIDS删除多条数据并清理缓存
		DeleteMultiCacheByUIDS(ctx context.Context, UIDS []string) error
		// DeleteMultiCacheByUIDS 根据UIDS删除多条数据并清理缓存
		DeleteMultiCacheByUIDSTx(ctx context.Context, tx *gorm_gen_dao.Query, UIDS []string) error
		// DeleteMultiByUIDS 根据UIDS删除多条数据
		DeleteMultiByUIDS(ctx context.Context, UIDS []string) error
		// DeleteMultiByUIDS 根据UIDS删除多条数据
		DeleteMultiByUIDSTx(ctx context.Context, tx *gorm_gen_dao.Query, UIDS []string) error
		// DeleteOneCacheByUIDStatus 根据UIDStatus删除一条数据并清理缓存
		DeleteOneCacheByUIDStatus(ctx context.Context, UID string, status int16) error
		// DeleteOneCacheByUIDStatus 根据UIDStatus删除一条数据并清理缓存
		DeleteOneCacheByUIDStatusTx(ctx context.Context, tx *gorm_gen_dao.Query, UID string, status int16) error
		// DeleteOneByUIDStatus 根据UIDStatus删除一条数据
		DeleteOneByUIDStatus(ctx context.Context, UID string, status int16) error
		// DeleteOneByUIDStatus 根据UIDStatus删除一条数据
		DeleteOneByUIDStatusTx(ctx context.Context, tx *gorm_gen_dao.Query, UID string, status int16) error
		// DeleteUniqueIndexCache 删除唯一索引存在的缓存
		DeleteUniqueIndexCache(ctx context.Context, data []*gorm_gen_model.UserDemo) error
	}
	UserDemoRepo struct {
		db    *gorm.DB
		cache cache.IDBCache
	}
)

func NewUserDemoRepo(db *gorm.DB, cache cache.IDBCache) *UserDemoRepo {
	return &UserDemoRepo{
		db:    db,
		cache: cache,
	}
}

// CreateOne 创建一条数据
func (u *UserDemoRepo) CreateOne(ctx context.Context, data *gorm_gen_model.UserDemo) error {
	dao := gorm_gen_dao.Use(u.db).UserDemo
	err := dao.WithContext(ctx).Create(data)
	if err != nil {
		return err
	}
	return nil
}

// CreateOneByTx 创建一条数据(事务)
func (u *UserDemoRepo) CreateOneByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.UserDemo) error {
	dao := tx.UserDemo
	err := dao.WithContext(ctx).Create(data)
	if err != nil {
		return err
	}
	return nil
}

// UpsertOne Upsert一条数据
func (u *UserDemoRepo) UpsertOne(ctx context.Context, data *gorm_gen_model.UserDemo) error {
	dao := gorm_gen_dao.Use(u.db).UserDemo
	err := dao.WithContext(ctx).Save(data)
	if err != nil {
		return err
	}
	return nil
}

// UpsertOneByTx Upsert一条数据(事务)
func (u *UserDemoRepo) UpsertOneByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.UserDemo) error {
	dao := tx.UserDemo
	err := dao.WithContext(ctx).Save(data)
	if err != nil {
		return err
	}
	return nil
}

// UpsertOneByFields Upsert一条数据，根据fields字段
func (u *UserDemoRepo) UpsertOneByFields(ctx context.Context, data *gorm_gen_model.UserDemo, fields []string) error {
	if len(fields) == 0 {
		return errors.New("UpsertOneByFields fields is empty")
	}
	columns := make([]clause.Column, 0)
	for _, v := range fields {
		columns = append(columns, clause.Column{Name: v})
	}
	dao := gorm_gen_dao.Use(u.db).UserDemo
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
func (u *UserDemoRepo) UpsertOneByFieldsTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.UserDemo, fields []string) error {
	if len(fields) == 0 {
		return errors.New("UpsertOneByFieldsTx fields is empty")
	}
	columns := make([]clause.Column, 0)
	for _, v := range fields {
		columns = append(columns, clause.Column{Name: v})
	}
	dao := tx.UserDemo
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
func (u *UserDemoRepo) CreateBatch(ctx context.Context, data []*gorm_gen_model.UserDemo, batchSize int) error {
	dao := gorm_gen_dao.Use(u.db).UserDemo
	err := dao.WithContext(ctx).CreateInBatches(data, batchSize)
	if err != nil {
		return err
	}
	return nil
}

// UpdateOne 更新一条数据
func (u *UserDemoRepo) UpdateOne(ctx context.Context, data *gorm_gen_model.UserDemo) error {
	dao := gorm_gen_dao.Use(u.db).UserDemo
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).Updates(data)
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, []*gorm_gen_model.UserDemo{data})
	if err != nil {
		return err
	}
	return nil
}

// UpdateOneByTx 更新一条数据(事务)
func (u *UserDemoRepo) UpdateOneByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.UserDemo) error {
	dao := tx.UserDemo
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).Updates(data)
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, []*gorm_gen_model.UserDemo{data})
	if err != nil {
		return err
	}
	return err
}

// UpdateOneWithZero 更新一条数据,包含零值
func (u *UserDemoRepo) UpdateOneWithZero(ctx context.Context, data *gorm_gen_model.UserDemo) error {
	dao := gorm_gen_dao.Use(u.db).UserDemo
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).Select(dao.ALL.WithTable("")).Updates(data)
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, []*gorm_gen_model.UserDemo{data})
	if err != nil {
		return err
	}
	return nil
}

// UpdateOneWithZeroByTx 更新一条数据(事务),包含零值
func (u *UserDemoRepo) UpdateOneWithZeroByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.UserDemo) error {
	dao := tx.UserDemo
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).Select(dao.ALL.WithTable("")).Updates(data)
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, []*gorm_gen_model.UserDemo{data})
	if err != nil {
		return err
	}
	return err
}

// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
func (u *UserDemoRepo) DeleteOneCacheByID(ctx context.Context, ID int64) error {
	dao := gorm_gen_dao.Use(u.db).UserDemo
	result, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if result == nil {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.ID.Eq(ID)).Delete()
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, []*gorm_gen_model.UserDemo{result})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
func (u *UserDemoRepo) DeleteOneCacheByIDTx(ctx context.Context, tx *gorm_gen_dao.Query, ID int64) error {
	dao := tx.UserDemo
	result, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if result == nil {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.ID.Eq(ID)).Delete()
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, []*gorm_gen_model.UserDemo{result})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneByID 根据ID删除一条数据
func (u *UserDemoRepo) DeleteOneByID(ctx context.Context, ID int64) error {
	dao := gorm_gen_dao.Use(u.db).UserDemo
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneByID 根据ID删除一条数据
func (u *UserDemoRepo) DeleteOneByIDTx(ctx context.Context, tx *gorm_gen_dao.Query, ID int64) error {
	dao := tx.UserDemo
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
func (u *UserDemoRepo) DeleteMultiCacheByIDS(ctx context.Context, IDS []int64) error {
	dao := gorm_gen_dao.Use(u.db).UserDemo
	result, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Find()
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Delete()
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, result)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
func (u *UserDemoRepo) DeleteMultiCacheByIDSTx(ctx context.Context, tx *gorm_gen_dao.Query, IDS []int64) error {
	dao := tx.UserDemo
	result, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Find()
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Delete()
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, result)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByIDS 根据IDS删除多条数据
func (u *UserDemoRepo) DeleteMultiByIDS(ctx context.Context, IDS []int64) error {
	dao := gorm_gen_dao.Use(u.db).UserDemo
	_, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByIDS 根据IDS删除多条数据
func (u *UserDemoRepo) DeleteMultiByIDSTx(ctx context.Context, tx *gorm_gen_dao.Query, IDS []int64) error {
	dao := tx.UserDemo
	_, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneCacheByUID 根据UID删除一条数据并清理缓存
func (u *UserDemoRepo) DeleteOneCacheByUID(ctx context.Context, UID string) error {
	dao := gorm_gen_dao.Use(u.db).UserDemo
	result, err := dao.WithContext(ctx).Where(dao.UID.Eq(UID)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if result == nil {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.UID.Eq(UID)).Delete()
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, []*gorm_gen_model.UserDemo{result})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneCacheByUID 根据UID删除一条数据并清理缓存
func (u *UserDemoRepo) DeleteOneCacheByUIDTx(ctx context.Context, tx *gorm_gen_dao.Query, UID string) error {
	dao := tx.UserDemo
	result, err := dao.WithContext(ctx).Where(dao.UID.Eq(UID)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if result == nil {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.UID.Eq(UID)).Delete()
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, []*gorm_gen_model.UserDemo{result})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneByUID 根据UID删除一条数据
func (u *UserDemoRepo) DeleteOneByUID(ctx context.Context, UID string) error {
	dao := gorm_gen_dao.Use(u.db).UserDemo
	_, err := dao.WithContext(ctx).Where(dao.UID.Eq(UID)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneByUID 根据UID删除一条数据
func (u *UserDemoRepo) DeleteOneByUIDTx(ctx context.Context, tx *gorm_gen_dao.Query, UID string) error {
	dao := tx.UserDemo
	_, err := dao.WithContext(ctx).Where(dao.UID.Eq(UID)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByUIDS 根据UIDS删除多条数据并清理缓存
func (u *UserDemoRepo) DeleteMultiCacheByUIDS(ctx context.Context, UIDS []string) error {
	dao := gorm_gen_dao.Use(u.db).UserDemo
	result, err := dao.WithContext(ctx).Where(dao.UID.In(UIDS...)).Find()
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.UID.In(UIDS...)).Delete()
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, result)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByUIDS 根据UIDS删除多条数据并清理缓存
func (u *UserDemoRepo) DeleteMultiCacheByUIDSTx(ctx context.Context, tx *gorm_gen_dao.Query, UIDS []string) error {
	dao := tx.UserDemo
	result, err := dao.WithContext(ctx).Where(dao.UID.In(UIDS...)).Find()
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.UID.In(UIDS...)).Delete()
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, result)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByUIDS 根据UIDS删除多条数据
func (u *UserDemoRepo) DeleteMultiByUIDS(ctx context.Context, UIDS []string) error {
	dao := gorm_gen_dao.Use(u.db).UserDemo
	_, err := dao.WithContext(ctx).Where(dao.UID.In(UIDS...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByUIDS 根据UIDS删除多条数据
func (u *UserDemoRepo) DeleteMultiByUIDSTx(ctx context.Context, tx *gorm_gen_dao.Query, UIDS []string) error {
	dao := tx.UserDemo
	_, err := dao.WithContext(ctx).Where(dao.UID.In(UIDS...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneCacheByUIDStatus 根据UID删除一条数据并清理缓存
func (u *UserDemoRepo) DeleteOneCacheByUIDStatus(ctx context.Context, UID string, status int16) error {
	dao := gorm_gen_dao.Use(u.db).UserDemo
	result, err := dao.WithContext(ctx).Where(dao.UID.Eq(UID), dao.Status.Eq(status)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if result == nil {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.UID.Eq(UID), dao.Status.Eq(status)).Delete()
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, []*gorm_gen_model.UserDemo{result})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneCacheByUIDStatus 根据UID删除一条数据并清理缓存
func (u *UserDemoRepo) DeleteOneCacheByUIDStatusTx(ctx context.Context, tx *gorm_gen_dao.Query, UID string, status int16) error {
	dao := tx.UserDemo
	result, err := dao.WithContext(ctx).Where(dao.UID.Eq(UID), dao.Status.Eq(status)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if result == nil {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.UID.Eq(UID), dao.Status.Eq(status)).Delete()
	if err != nil {
		return err
	}
	err = u.DeleteUniqueIndexCache(ctx, []*gorm_gen_model.UserDemo{result})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneByUIDStatus 根据UID删除一条数据
func (u *UserDemoRepo) DeleteOneByUIDStatus(ctx context.Context, UID string, status int16) error {
	dao := gorm_gen_dao.Use(u.db).UserDemo
	_, err := dao.WithContext(ctx).Where(dao.UID.Eq(UID), dao.Status.Eq(status)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneByUIDStatus 根据UID删除一条数据
func (u *UserDemoRepo) DeleteOneByUIDStatusTx(ctx context.Context, tx *gorm_gen_dao.Query, UID string, status int16) error {
	dao := tx.UserDemo
	_, err := dao.WithContext(ctx).Where(dao.UID.Eq(UID), dao.Status.Eq(status)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteUniqueIndexCache 删除唯一索引存在的缓存
func (u *UserDemoRepo) DeleteUniqueIndexCache(ctx context.Context, data []*gorm_gen_model.UserDemo) error {
	keys := make([]string, 0)
	for _, v := range data {
		keys = append(keys, u.cache.Key(cacheUserDemoByIDPrefix, v.ID))
		keys = append(keys, u.cache.Key(cacheUserDemoByUIDPrefix, v.UID))
		keys = append(keys, u.cache.Key(cacheUserDemoByUIDStatusPrefix, v.UID, v.Status))

	}
	err := u.cache.DelBatch(ctx, keys)
	if err != nil {
		return err
	}
	return nil
}

// FindOneCacheByID 根据ID查询一条数据并设置缓存
func (u *UserDemoRepo) FindOneCacheByID(ctx context.Context, ID int64) (*gorm_gen_model.UserDemo, error) {
	resp := new(gorm_gen_model.UserDemo)
	cacheKey := u.cache.Key(cacheUserDemoByIDPrefix, ID)
	cacheValue, err := u.cache.Fetch(ctx, cacheKey, func() (string, error) {
		dao := gorm_gen_dao.Use(u.db).UserDemo
		result, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).First()
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return "", err
		}
		marshal, err := json.Marshal(result)
		if err != nil {
			return "", err
		}
		return string(marshal), nil
	})
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(cacheValue), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// FindOneByID 根据ID查询一条数据
func (u *UserDemoRepo) FindOneByID(ctx context.Context, ID int64) (*gorm_gen_model.UserDemo, error) {
	dao := gorm_gen_dao.Use(u.db).UserDemo
	result, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return result, nil
}

// FindMultiCacheByIDS 根据IDS查询多条数据并设置缓存
func (u *UserDemoRepo) FindMultiCacheByIDS(ctx context.Context, IDS []int64) ([]*gorm_gen_model.UserDemo, error) {
	resp := make([]*gorm_gen_model.UserDemo, 0)
	cacheKeys := make([]string, 0)
	keyToParam := make(map[string]int64)
	for _, v := range IDS {
		cacheKey := u.cache.Key(cacheUserDemoByIDPrefix, v)
		cacheKeys = append(cacheKeys, cacheKey)
		keyToParam[cacheKey] = v
	}
	cacheValue, err := u.cache.FetchBatch(ctx, cacheKeys, func(miss []string) (map[string]string, error) {
		parameters := make([]int64, 0)
		for _, v := range miss {
			parameters = append(parameters, keyToParam[v])
		}
		dao := gorm_gen_dao.Use(u.db).UserDemo
		result, err := dao.WithContext(ctx).Where(dao.ID.In(parameters...)).Find()
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		value := make(map[string]string)
		for _, v := range miss {
			value[v] = ""
		}
		for _, v := range result {
			marshal, err := json.Marshal(v)
			if err != nil {
				return nil, err
			}
			value[u.cache.Key(cacheUserDemoByIDPrefix, v.ID)] = string(marshal)
		}
		return value, nil
	})
	if err != nil {
		return nil, err
	}
	for _, v := range cacheValue {
		tmp := new(gorm_gen_model.UserDemo)
		err := json.Unmarshal([]byte(v), tmp)
		if err != nil {
			return nil, err
		}
		resp = append(resp, tmp)
	}
	return resp, nil
}

// FindMultiByIDS 根据IDS查询多条数据
func (u *UserDemoRepo) FindMultiByIDS(ctx context.Context, IDS []int64) ([]*gorm_gen_model.UserDemo, error) {
	dao := gorm_gen_dao.Use(u.db).UserDemo
	result, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindOneCacheByUID 根据UID查询一条数据并设置缓存
func (u *UserDemoRepo) FindOneCacheByUID(ctx context.Context, UID string) (*gorm_gen_model.UserDemo, error) {
	resp := new(gorm_gen_model.UserDemo)
	cacheKey := u.cache.Key(cacheUserDemoByUIDPrefix, UID)
	cacheValue, err := u.cache.Fetch(ctx, cacheKey, func() (string, error) {
		dao := gorm_gen_dao.Use(u.db).UserDemo
		result, err := dao.WithContext(ctx).Where(dao.UID.Eq(UID)).First()
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return "", err
		}
		marshal, err := json.Marshal(result)
		if err != nil {
			return "", err
		}
		return string(marshal), nil
	})
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(cacheValue), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// FindOneByUID 根据UID查询一条数据
func (u *UserDemoRepo) FindOneByUID(ctx context.Context, UID string) (*gorm_gen_model.UserDemo, error) {
	dao := gorm_gen_dao.Use(u.db).UserDemo
	result, err := dao.WithContext(ctx).Where(dao.UID.Eq(UID)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return result, nil
}

// FindMultiCacheByUIDS 根据UIDS查询多条数据并设置缓存
func (u *UserDemoRepo) FindMultiCacheByUIDS(ctx context.Context, UIDS []string) ([]*gorm_gen_model.UserDemo, error) {
	resp := make([]*gorm_gen_model.UserDemo, 0)
	cacheKeys := make([]string, 0)
	keyToParam := make(map[string]string)
	for _, v := range UIDS {
		cacheKey := u.cache.Key(cacheUserDemoByUIDPrefix, v)
		cacheKeys = append(cacheKeys, cacheKey)
		keyToParam[cacheKey] = v
	}
	cacheValue, err := u.cache.FetchBatch(ctx, cacheKeys, func(miss []string) (map[string]string, error) {
		parameters := make([]string, 0)
		for _, v := range miss {
			parameters = append(parameters, keyToParam[v])
		}
		dao := gorm_gen_dao.Use(u.db).UserDemo
		result, err := dao.WithContext(ctx).Where(dao.UID.In(parameters...)).Find()
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		value := make(map[string]string)
		for _, v := range miss {
			value[v] = ""
		}
		for _, v := range result {
			marshal, err := json.Marshal(v)
			if err != nil {
				return nil, err
			}
			value[u.cache.Key(cacheUserDemoByUIDPrefix, v.UID)] = string(marshal)
		}
		return value, nil
	})
	if err != nil {
		return nil, err
	}
	for _, v := range cacheValue {
		tmp := new(gorm_gen_model.UserDemo)
		err := json.Unmarshal([]byte(v), tmp)
		if err != nil {
			return nil, err
		}
		resp = append(resp, tmp)
	}
	return resp, nil
}

// FindMultiByUIDS 根据UIDS查询多条数据
func (u *UserDemoRepo) FindMultiByUIDS(ctx context.Context, UIDS []string) ([]*gorm_gen_model.UserDemo, error) {
	dao := gorm_gen_dao.Use(u.db).UserDemo
	result, err := dao.WithContext(ctx).Where(dao.UID.In(UIDS...)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindOneCacheByUIDStatus 根据UIDStatus查询一条数据并设置缓存
func (u *UserDemoRepo) FindOneCacheByUIDStatus(ctx context.Context, UID string, status int16) (*gorm_gen_model.UserDemo, error) {
	resp := new(gorm_gen_model.UserDemo)
	cacheKey := u.cache.Key(cacheUserDemoByUIDStatusPrefix, UID, status)
	cacheValue, err := u.cache.Fetch(ctx, cacheKey, func() (string, error) {
		dao := gorm_gen_dao.Use(u.db).UserDemo
		result, err := dao.WithContext(ctx).Where(dao.UID.Eq(UID), dao.Status.Eq(status)).First()
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return "", err
		}
		marshal, err := json.Marshal(result)
		if err != nil {
			return "", err
		}
		return string(marshal), nil
	})
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(cacheValue), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// FindOneByUIDStatus 根据UIDStatus查询一条数据
func (u *UserDemoRepo) FindOneByUIDStatus(ctx context.Context, UID string, status int16) (*gorm_gen_model.UserDemo, error) {
	dao := gorm_gen_dao.Use(u.db).UserDemo
	result, err := dao.WithContext(ctx).Where(dao.UID.Eq(UID), dao.Status.Eq(status)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return result, nil
}

// FindMultiByTenantIDDeptID 根据TenantIDDeptID查询多条数据
func (u *UserDemoRepo) FindMultiByTenantIDDeptID(ctx context.Context, tenantID int64, deptID int64) ([]*gorm_gen_model.UserDemo, error) {
	dao := gorm_gen_dao.Use(u.db).UserDemo
	result, err := dao.WithContext(ctx).Where(dao.TenantID.Eq(tenantID), dao.DeptID.Eq(deptID)).Find()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return result, nil
}

// FindMultiByUsername 根据username查询多条数据
func (u *UserDemoRepo) FindMultiByUsername(ctx context.Context, username string) ([]*gorm_gen_model.UserDemo, error) {
	dao := gorm_gen_dao.Use(u.db).UserDemo
	result, err := dao.WithContext(ctx).Where(dao.Username.Eq(username)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindMultiByUsernames 根据usernames查询多条数据
func (u *UserDemoRepo) FindMultiByUsernames(ctx context.Context, usernames []string) ([]*gorm_gen_model.UserDemo, error) {
	dao := gorm_gen_dao.Use(u.db).UserDemo
	result, err := dao.WithContext(ctx).Where(dao.Username.In(usernames...)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindMultiByTenantID 根据tenantID查询多条数据
func (u *UserDemoRepo) FindMultiByTenantID(ctx context.Context, tenantID int64) ([]*gorm_gen_model.UserDemo, error) {
	dao := gorm_gen_dao.Use(u.db).UserDemo
	result, err := dao.WithContext(ctx).Where(dao.TenantID.Eq(tenantID)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindMultiByTenantIDS 根据tenantIDS查询多条数据
func (u *UserDemoRepo) FindMultiByTenantIDS(ctx context.Context, tenantIDS []int64) ([]*gorm_gen_model.UserDemo, error) {
	dao := gorm_gen_dao.Use(u.db).UserDemo
	result, err := dao.WithContext(ctx).Where(dao.TenantID.In(tenantIDS...)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindMultiByPaginator 查询分页数据(通用)
func (u *UserDemoRepo) FindMultiByPaginator(ctx context.Context, paginatorReq *paginator.Req) ([]*gorm_gen_model.UserDemo, *paginator.Reply, error) {
	result := make([]*gorm_gen_model.UserDemo, 0)
	var total int64
	whereExpressions, orderExpressions, err := paginatorReq.ConvertToGormExpression(gorm_gen_model.UserDemo{})
	if err != nil {
		return result, nil, err
	}
	err = u.db.WithContext(ctx).Model(&gorm_gen_model.UserDemo{}).Select([]string{"*"}).Clauses(whereExpressions...).Count(&total).Error
	if err != nil {
		return result, nil, err
	}
	if total == 0 {
		return result, nil, nil
	}
	paginatorReply, err := paginatorReq.ConvertToPage(int(total))
	if err != nil {
		return result, nil, err
	}
	query := u.db.WithContext(ctx).Model(&gorm_gen_model.UserDemo{}).Clauses(whereExpressions...).Clauses(orderExpressions...)
	if paginatorReply.Offset != 0 {
		query = query.Offset(paginatorReply.Offset)
	}
	if paginatorReply.Limit != 0 {
		query = query.Limit(paginatorReply.Limit)
	}
	err = query.Find(&result).Error
	if err != nil {
		return result, nil, err
	}
	return result, paginatorReply, err
}
