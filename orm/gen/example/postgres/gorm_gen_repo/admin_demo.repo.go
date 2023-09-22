// Code generated by gen/repo. DO NOT EDIT.
// Code generated by gen/repo. DO NOT EDIT.
// Code generated by gen/repo. DO NOT EDIT.

package gorm_gen_repo

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/fzf-labs/fpkg/orm"
	"github.com/fzf-labs/fpkg/orm/gen/example/postgres/gorm_gen_dao"
	"github.com/fzf-labs/fpkg/orm/gen/example/postgres/gorm_gen_model"
	"gorm.io/gorm"
)

var _ IAdminDemoRepo = (*AdminDemoRepo)(nil)

var (
	cacheAdminDemoByUsernamePrefix = "DBCache:gorm_gen:AdminDemoByUsername"
	cacheAdminDemoByIDPrefix       = "DBCache:gorm_gen:AdminDemoByID"
)

type (
	IAdminDemoRepo interface {
		// CreateOne 创建一条数据
		CreateOne(ctx context.Context, data *gorm_gen_model.AdminDemo) error
		// CreateOneByTx 创建一条数据(事务)
		CreateOneByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminDemo) error
		// CreateBatch 批量创建数据
		CreateBatch(ctx context.Context, data []*gorm_gen_model.AdminDemo, batchSize int) error
		// UpdateOne 更新一条数据
		UpdateOne(ctx context.Context, data *gorm_gen_model.AdminDemo) error
		// UpdateOne 更新一条数据(事务)
		UpdateOneByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminDemo) error
		// FindOneCacheByUsername 根据username查询一条数据并设置缓存
		FindOneCacheByUsername(ctx context.Context, username string) (*gorm_gen_model.AdminDemo, error)
		// FindOneByUsername 根据username查询一条数据
		FindOneByUsername(ctx context.Context, username string) (*gorm_gen_model.AdminDemo, error)
		// FindMultiCacheByUsernames 根据usernames查询多条数据并设置缓存
		FindMultiCacheByUsernames(ctx context.Context, usernames []string) ([]*gorm_gen_model.AdminDemo, error)
		// FindMultiByUsernames 根据usernames查询多条数据
		FindMultiByUsernames(ctx context.Context, usernames []string) ([]*gorm_gen_model.AdminDemo, error)
		// FindOneCacheByID 根据ID查询一条数据并设置缓存
		FindOneCacheByID(ctx context.Context, ID string) (*gorm_gen_model.AdminDemo, error)
		// FindOneByID 根据ID查询一条数据
		FindOneByID(ctx context.Context, ID string) (*gorm_gen_model.AdminDemo, error)
		// FindMultiCacheByIDS 根据IDS查询多条数据并设置缓存
		FindMultiCacheByIDS(ctx context.Context, IDS []string) ([]*gorm_gen_model.AdminDemo, error)
		// FindMultiByIDS 根据IDS查询多条数据
		FindMultiByIDS(ctx context.Context, IDS []string) ([]*gorm_gen_model.AdminDemo, error)
		// FindMultiByPaginator 查询分页数据(通用)
		FindMultiByPaginator(ctx context.Context, paginatorReq *orm.PaginatorReq) ([]*gorm_gen_model.AdminDemo, *orm.PaginatorReply, error)
		// DeleteOneCacheByUsername 根据username删除一条数据并清理缓存
		DeleteOneCacheByUsername(ctx context.Context, username string) error
		// DeleteOneCacheByUsername 根据username删除一条数据并清理缓存
		DeleteOneCacheByUsernameTx(ctx context.Context, tx *gorm_gen_dao.Query, username string) error
		// DeleteOneByUsername 根据username删除一条数据
		DeleteOneByUsername(ctx context.Context, username string) error
		// DeleteOneByUsername 根据username删除一条数据
		DeleteOneByUsernameTx(ctx context.Context, tx *gorm_gen_dao.Query, username string) error
		// DeleteMultiCacheByUsernames 根据Usernames删除多条数据并清理缓存
		DeleteMultiCacheByUsernames(ctx context.Context, usernames []string) error
		// DeleteMultiCacheByUsernames 根据Usernames删除多条数据并清理缓存
		DeleteMultiCacheByUsernamesTx(ctx context.Context, tx *gorm_gen_dao.Query, usernames []string) error
		// DeleteMultiByUsernames 根据Usernames删除多条数据
		DeleteMultiByUsernames(ctx context.Context, usernames []string) error
		// DeleteMultiByUsernames 根据Usernames删除多条数据
		DeleteMultiByUsernamesTx(ctx context.Context, tx *gorm_gen_dao.Query, usernames []string) error
		// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
		DeleteOneCacheByID(ctx context.Context, ID string) error
		// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
		DeleteOneCacheByIDTx(ctx context.Context, tx *gorm_gen_dao.Query, ID string) error
		// DeleteOneByID 根据ID删除一条数据
		DeleteOneByID(ctx context.Context, ID string) error
		// DeleteOneByID 根据ID删除一条数据
		DeleteOneByIDTx(ctx context.Context, tx *gorm_gen_dao.Query, ID string) error
		// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
		DeleteMultiCacheByIDS(ctx context.Context, IDS []string) error
		// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
		DeleteMultiCacheByIDSTx(ctx context.Context, tx *gorm_gen_dao.Query, IDS []string) error
		// DeleteMultiByIDS 根据IDS删除多条数据
		DeleteMultiByIDS(ctx context.Context, IDS []string) error
		// DeleteMultiByIDS 根据IDS删除多条数据
		DeleteMultiByIDSTx(ctx context.Context, tx *gorm_gen_dao.Query, IDS []string) error
		// DeleteUniqueIndexCache 删除唯一索引存在的缓存
		DeleteUniqueIndexCache(ctx context.Context, data []*gorm_gen_model.AdminDemo) error
	}
	IAdminDemoCache interface {
		Key(fields ...any) string
		Fetch(ctx context.Context, key string, fn func() (string, error)) (string, error)
		FetchBatch(ctx context.Context, keys []string, fn func(miss []string) (map[string]string, error)) (map[string]string, error)
		Del(ctx context.Context, key string) error
		DelBatch(ctx context.Context, keys []string) error
	}
	AdminDemoRepo struct {
		db    *gorm.DB
		cache IAdminDemoCache
	}
)

func NewAdminDemoRepo(db *gorm.DB, cache IAdminDemoCache) *AdminDemoRepo {
	return &AdminDemoRepo{
		db:    db,
		cache: cache,
	}
}

// CreateOne 创建一条数据
func (a *AdminDemoRepo) CreateOne(ctx context.Context, data *gorm_gen_model.AdminDemo) error {
	dao := gorm_gen_dao.Use(a.db).AdminDemo
	err := dao.WithContext(ctx).Create(data)
	if err != nil {
		return err
	}
	return nil
}

// CreateOneByTx 创建一条数据(事务)
func (a *AdminDemoRepo) CreateOneByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminDemo) error {
	dao := tx.AdminDemo
	err := dao.WithContext(ctx).Create(data)
	if err != nil {
		return err
	}
	return nil
}

// CreateBatch 批量创建数据
func (a *AdminDemoRepo) CreateBatch(ctx context.Context, data []*gorm_gen_model.AdminDemo, batchSize int) error {
	dao := gorm_gen_dao.Use(a.db).AdminDemo
	err := dao.WithContext(ctx).CreateInBatches(data, batchSize)
	if err != nil {
		return err
	}
	return nil
}

// UpdateOne 更新一条数据
func (a *AdminDemoRepo) UpdateOne(ctx context.Context, data *gorm_gen_model.AdminDemo) error {
	dao := gorm_gen_dao.Use(a.db).AdminDemo
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).Updates(data)
	if err != nil {
		return err
	}
	err = a.DeleteUniqueIndexCache(ctx, []*gorm_gen_model.AdminDemo{data})
	if err != nil {
		return err
	}
	return nil
}

// UpdateOneByTx 更新一条数据(事务)
func (a *AdminDemoRepo) UpdateOneByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.AdminDemo) error {
	dao := tx.AdminDemo
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).Updates(data)
	if err != nil {
		return err
	}
	err = a.DeleteUniqueIndexCache(ctx, []*gorm_gen_model.AdminDemo{data})
	if err != nil {
		return err
	}
	return err
}

// DeleteOneCacheByUsername 根据username删除一条数据并清理缓存
func (a *AdminDemoRepo) DeleteOneCacheByUsername(ctx context.Context, username string) error {
	dao := gorm_gen_dao.Use(a.db).AdminDemo
	result, err := dao.WithContext(ctx).Where(dao.Username.Eq(username)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if result == nil {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.Username.Eq(username)).Delete()
	if err != nil {
		return err
	}
	err = a.DeleteUniqueIndexCache(ctx, []*gorm_gen_model.AdminDemo{result})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneCacheByUsername 根据username删除一条数据并清理缓存
func (a *AdminDemoRepo) DeleteOneCacheByUsernameTx(ctx context.Context, tx *gorm_gen_dao.Query, username string) error {
	dao := tx.AdminDemo
	result, err := dao.WithContext(ctx).Where(dao.Username.Eq(username)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if result == nil {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.Username.Eq(username)).Delete()
	if err != nil {
		return err
	}
	err = a.DeleteUniqueIndexCache(ctx, []*gorm_gen_model.AdminDemo{result})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneByUsername 根据username删除一条数据
func (a *AdminDemoRepo) DeleteOneByUsername(ctx context.Context, username string) error {
	dao := gorm_gen_dao.Use(a.db).AdminDemo
	_, err := dao.WithContext(ctx).Where(dao.Username.Eq(username)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneByUsername 根据username删除一条数据
func (a *AdminDemoRepo) DeleteOneByUsernameTx(ctx context.Context, tx *gorm_gen_dao.Query, username string) error {
	dao := tx.AdminDemo
	_, err := dao.WithContext(ctx).Where(dao.Username.Eq(username)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByUsernames 根据usernames删除多条数据并清理缓存
func (a *AdminDemoRepo) DeleteMultiCacheByUsernames(ctx context.Context, usernames []string) error {
	dao := gorm_gen_dao.Use(a.db).AdminDemo
	result, err := dao.WithContext(ctx).Where(dao.Username.In(usernames...)).Find()
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.Username.In(usernames...)).Delete()
	if err != nil {
		return err
	}
	err = a.DeleteUniqueIndexCache(ctx, result)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByUsernames 根据usernames删除多条数据并清理缓存
func (a *AdminDemoRepo) DeleteMultiCacheByUsernamesTx(ctx context.Context, tx *gorm_gen_dao.Query, usernames []string) error {
	dao := tx.AdminDemo
	result, err := dao.WithContext(ctx).Where(dao.Username.In(usernames...)).Find()
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.Username.In(usernames...)).Delete()
	if err != nil {
		return err
	}
	err = a.DeleteUniqueIndexCache(ctx, result)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByUsernames 根据usernames删除多条数据
func (a *AdminDemoRepo) DeleteMultiByUsernames(ctx context.Context, usernames []string) error {
	dao := gorm_gen_dao.Use(a.db).AdminDemo
	_, err := dao.WithContext(ctx).Where(dao.Username.In(usernames...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByUsernames 根据usernames删除多条数据
func (a *AdminDemoRepo) DeleteMultiByUsernamesTx(ctx context.Context, tx *gorm_gen_dao.Query, usernames []string) error {
	dao := tx.AdminDemo
	_, err := dao.WithContext(ctx).Where(dao.Username.In(usernames...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
func (a *AdminDemoRepo) DeleteOneCacheByID(ctx context.Context, ID string) error {
	dao := gorm_gen_dao.Use(a.db).AdminDemo
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
	err = a.DeleteUniqueIndexCache(ctx, []*gorm_gen_model.AdminDemo{result})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
func (a *AdminDemoRepo) DeleteOneCacheByIDTx(ctx context.Context, tx *gorm_gen_dao.Query, ID string) error {
	dao := tx.AdminDemo
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
	err = a.DeleteUniqueIndexCache(ctx, []*gorm_gen_model.AdminDemo{result})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneByID 根据ID删除一条数据
func (a *AdminDemoRepo) DeleteOneByID(ctx context.Context, ID string) error {
	dao := gorm_gen_dao.Use(a.db).AdminDemo
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneByID 根据ID删除一条数据
func (a *AdminDemoRepo) DeleteOneByIDTx(ctx context.Context, tx *gorm_gen_dao.Query, ID string) error {
	dao := tx.AdminDemo
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
func (a *AdminDemoRepo) DeleteMultiCacheByIDS(ctx context.Context, IDS []string) error {
	dao := gorm_gen_dao.Use(a.db).AdminDemo
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
	err = a.DeleteUniqueIndexCache(ctx, result)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
func (a *AdminDemoRepo) DeleteMultiCacheByIDSTx(ctx context.Context, tx *gorm_gen_dao.Query, IDS []string) error {
	dao := tx.AdminDemo
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
	err = a.DeleteUniqueIndexCache(ctx, result)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByIDS 根据IDS删除多条数据
func (a *AdminDemoRepo) DeleteMultiByIDS(ctx context.Context, IDS []string) error {
	dao := gorm_gen_dao.Use(a.db).AdminDemo
	_, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByIDS 根据IDS删除多条数据
func (a *AdminDemoRepo) DeleteMultiByIDSTx(ctx context.Context, tx *gorm_gen_dao.Query, IDS []string) error {
	dao := tx.AdminDemo
	_, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteUniqueIndexCache 删除唯一索引存在的缓存
func (a *AdminDemoRepo) DeleteUniqueIndexCache(ctx context.Context, data []*gorm_gen_model.AdminDemo) error {
	keys := make([]string, 0)
	for _, v := range data {
		keys = append(keys, a.cache.Key(cacheAdminDemoByUsernamePrefix, v.Username))
		keys = append(keys, a.cache.Key(cacheAdminDemoByIDPrefix, v.ID))

	}
	err := a.cache.DelBatch(ctx, keys)
	if err != nil {
		return err
	}
	return nil
}

// FindOneCacheByUsername 根据username查询一条数据并设置缓存
func (a *AdminDemoRepo) FindOneCacheByUsername(ctx context.Context, username string) (*gorm_gen_model.AdminDemo, error) {
	resp := new(gorm_gen_model.AdminDemo)
	cacheKey := a.cache.Key(cacheAdminDemoByUsernamePrefix, username)
	cacheValue, err := a.cache.Fetch(ctx, cacheKey, func() (string, error) {
		dao := gorm_gen_dao.Use(a.db).AdminDemo
		result, err := dao.WithContext(ctx).Where(dao.Username.Eq(username)).First()
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

// FindOneByUsername 根据username查询一条数据
func (a *AdminDemoRepo) FindOneByUsername(ctx context.Context, username string) (*gorm_gen_model.AdminDemo, error) {
	dao := gorm_gen_dao.Use(a.db).AdminDemo
	result, err := dao.WithContext(ctx).Where(dao.Username.Eq(username)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return result, nil
}

// FindMultiCacheByUsernames 根据usernames查询多条数据并设置缓存
func (a *AdminDemoRepo) FindMultiCacheByUsernames(ctx context.Context, usernames []string) ([]*gorm_gen_model.AdminDemo, error) {
	resp := make([]*gorm_gen_model.AdminDemo, 0)
	cacheKeys := make([]string, 0)
	keyToParam := make(map[string]string)
	for _, v := range usernames {
		cacheKey := a.cache.Key(cacheAdminDemoByUsernamePrefix, v)
		cacheKeys = append(cacheKeys, cacheKey)
		keyToParam[cacheKey] = v
	}
	cacheValue, err := a.cache.FetchBatch(ctx, cacheKeys, func(miss []string) (map[string]string, error) {
		parameters := make([]string, 0)
		for _, v := range miss {
			parameters = append(parameters, keyToParam[v])
		}
		dao := gorm_gen_dao.Use(a.db).AdminDemo
		result, err := dao.WithContext(ctx).Where(dao.Username.In(parameters...)).Find()
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
			value[a.cache.Key(cacheAdminDemoByUsernamePrefix, v.Username)] = string(marshal)
		}
		return value, nil
	})
	if err != nil {
		return nil, err
	}
	for _, v := range cacheValue {
		tmp := new(gorm_gen_model.AdminDemo)
		err := json.Unmarshal([]byte(v), tmp)
		if err != nil {
			return nil, err
		}
		resp = append(resp, tmp)
	}
	return resp, nil
}

// FindMultiByUsernames 根据usernames查询多条数据
func (a *AdminDemoRepo) FindMultiByUsernames(ctx context.Context, usernames []string) ([]*gorm_gen_model.AdminDemo, error) {
	dao := gorm_gen_dao.Use(a.db).AdminDemo
	result, err := dao.WithContext(ctx).Where(dao.Username.In(usernames...)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindOneCacheByID 根据ID查询一条数据并设置缓存
func (a *AdminDemoRepo) FindOneCacheByID(ctx context.Context, ID string) (*gorm_gen_model.AdminDemo, error) {
	resp := new(gorm_gen_model.AdminDemo)
	cacheKey := a.cache.Key(cacheAdminDemoByIDPrefix, ID)
	cacheValue, err := a.cache.Fetch(ctx, cacheKey, func() (string, error) {
		dao := gorm_gen_dao.Use(a.db).AdminDemo
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
func (a *AdminDemoRepo) FindOneByID(ctx context.Context, ID string) (*gorm_gen_model.AdminDemo, error) {
	dao := gorm_gen_dao.Use(a.db).AdminDemo
	result, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return result, nil
}

// FindMultiCacheByIDS 根据IDS查询多条数据并设置缓存
func (a *AdminDemoRepo) FindMultiCacheByIDS(ctx context.Context, IDS []string) ([]*gorm_gen_model.AdminDemo, error) {
	resp := make([]*gorm_gen_model.AdminDemo, 0)
	cacheKeys := make([]string, 0)
	keyToParam := make(map[string]string)
	for _, v := range IDS {
		cacheKey := a.cache.Key(cacheAdminDemoByIDPrefix, v)
		cacheKeys = append(cacheKeys, cacheKey)
		keyToParam[cacheKey] = v
	}
	cacheValue, err := a.cache.FetchBatch(ctx, cacheKeys, func(miss []string) (map[string]string, error) {
		parameters := make([]string, 0)
		for _, v := range miss {
			parameters = append(parameters, keyToParam[v])
		}
		dao := gorm_gen_dao.Use(a.db).AdminDemo
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
			value[a.cache.Key(cacheAdminDemoByIDPrefix, v.ID)] = string(marshal)
		}
		return value, nil
	})
	if err != nil {
		return nil, err
	}
	for _, v := range cacheValue {
		tmp := new(gorm_gen_model.AdminDemo)
		err := json.Unmarshal([]byte(v), tmp)
		if err != nil {
			return nil, err
		}
		resp = append(resp, tmp)
	}
	return resp, nil
}

// FindMultiByIDS 根据IDS查询多条数据
func (a *AdminDemoRepo) FindMultiByIDS(ctx context.Context, IDS []string) ([]*gorm_gen_model.AdminDemo, error) {
	dao := gorm_gen_dao.Use(a.db).AdminDemo
	result, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindMultiByPaginator 查询分页数据(通用)
func (a *AdminDemoRepo) FindMultiByPaginator(ctx context.Context, paginatorReq *orm.PaginatorReq) ([]*gorm_gen_model.AdminDemo, *orm.PaginatorReply, error) {
	result := make([]*gorm_gen_model.AdminDemo, 0)
	var total int64
	queryStr, args, err := paginatorReq.ConvertToGormConditions()
	if err != nil {
		return nil, nil, err
	}
	err = a.db.WithContext(ctx).Model(&gorm_gen_model.AdminDemo{}).Select([]string{"id"}).Where(queryStr, args...).Count(&total).Error
	if err != nil {
		return nil, nil, err
	}
	if total == 0 {
		return nil, nil, nil
	}
	query := a.db.WithContext(ctx)
	order := paginatorReq.ConvertToOrder()
	if order != "" {
		query = query.Order(order)
	}
	paginatorReply := paginatorReq.ConvertToPage(int(total))
	err = query.Limit(paginatorReply.Limit).Offset(paginatorReply.Offset).Where(queryStr, args...).Find(&result).Error
	if err != nil {
		return nil, nil, err
	}
	return result, paginatorReply, err
}
