// Code generated by gen/repo. DO NOT EDIT.
// Code generated by gen/repo. DO NOT EDIT.
// Code generated by gen/repo. DO NOT EDIT.

package gorm_gen_repo

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/fzf-labs/fpkg/orm/gen/example/postgres/gorm_gen_dao"
	"github.com/fzf-labs/fpkg/orm/gen/example/postgres/gorm_gen_model"
	"gorm.io/gorm"
)

var _ ISysLogRepo = (*SysLogRepo)(nil)

var (
	cacheSysLogByIDPrefix = "DBCache:gorm_gen:SysLogByID"
)

type (
	ISysLogRepo interface {
		// CreateOne 创建一条数据
		CreateOne(ctx context.Context, data *gorm_gen_model.SysLog) error
		// CreateOneByTx 创建一条数据(事务)
		CreateOneByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.SysLog) error
		// CreateBatch 批量创建数据
		CreateBatch(ctx context.Context, data []*gorm_gen_model.SysLog, batchSize int) error
		// UpdateOne 更新一条数据
		UpdateOne(ctx context.Context, data *gorm_gen_model.SysLog) error
		// UpdateOne 更新一条数据(事务)
		UpdateOneByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.SysLog) error
		// FindOneCacheByID 根据ID查询一条数据并设置缓存
		FindOneCacheByID(ctx context.Context, ID string) (*gorm_gen_model.SysLog, error)
		// FindOneByID 根据ID查询一条数据
		FindOneByID(ctx context.Context, ID string) (*gorm_gen_model.SysLog, error)
		// FindMultiCacheByIDS 根据IDS查询多条数据并设置缓存
		FindMultiCacheByIDS(ctx context.Context, IDS []string) ([]*gorm_gen_model.SysLog, error)
		// FindMultiByIDS 根据IDS查询多条数据
		FindMultiByIDS(ctx context.Context, IDS []string) ([]*gorm_gen_model.SysLog, error)
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
		DeleteUniqueIndexCache(ctx context.Context, data []*gorm_gen_model.SysLog) error
	}
	ISysLogCache interface {
		Key(fields ...any) string
		Fetch(ctx context.Context, key string, fn func() (string, error)) (string, error)
		FetchBatch(ctx context.Context, keys []string, fn func(miss []string) (map[string]string, error)) (map[string]string, error)
		Del(ctx context.Context, key string) error
		DelBatch(ctx context.Context, keys []string) error
	}
	SysLogRepo struct {
		db    *gorm.DB
		cache ISysLogCache
	}
)

func NewSysLogRepo(db *gorm.DB, cache ISysLogCache) *SysLogRepo {
	return &SysLogRepo{
		db:    db,
		cache: cache,
	}
}

// CreateOne 创建一条数据
func (s *SysLogRepo) CreateOne(ctx context.Context, data *gorm_gen_model.SysLog) error {
	dao := gorm_gen_dao.Use(s.db).SysLog
	err := dao.WithContext(ctx).Create(data)
	if err != nil {
		return err
	}
	return nil
}

// CreateOneByTx 创建一条数据(事务)
func (s *SysLogRepo) CreateOneByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.SysLog) error {
	dao := tx.SysLog
	err := dao.WithContext(ctx).Create(data)
	if err != nil {
		return err
	}
	return nil
}

// CreateBatch 批量创建数据
func (s *SysLogRepo) CreateBatch(ctx context.Context, data []*gorm_gen_model.SysLog, batchSize int) error {
	dao := gorm_gen_dao.Use(s.db).SysLog
	err := dao.WithContext(ctx).CreateInBatches(data, batchSize)
	if err != nil {
		return err
	}
	return nil
}

// UpdateOne 更新一条数据
func (s *SysLogRepo) UpdateOne(ctx context.Context, data *gorm_gen_model.SysLog) error {
	dao := gorm_gen_dao.Use(s.db).SysLog
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).Updates(data)
	if err != nil {
		return err
	}
	err = s.DeleteUniqueIndexCache(ctx, []*gorm_gen_model.SysLog{data})
	if err != nil {
		return err
	}
	return nil
}

// UpdateOneByTx 更新一条数据(事务)
func (s *SysLogRepo) UpdateOneByTx(ctx context.Context, tx *gorm_gen_dao.Query, data *gorm_gen_model.SysLog) error {
	dao := tx.SysLog
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).Updates(data)
	if err != nil {
		return err
	}
	err = s.DeleteUniqueIndexCache(ctx, []*gorm_gen_model.SysLog{data})
	if err != nil {
		return err
	}
	return err
}

// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
func (s *SysLogRepo) DeleteOneCacheByID(ctx context.Context, ID string) error {
	dao := gorm_gen_dao.Use(s.db).SysLog
	first, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if first == nil {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.ID.Eq(ID)).Delete()
	if err != nil {
		return err
	}
	err = s.DeleteUniqueIndexCache(ctx, []*gorm_gen_model.SysLog{first})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneCacheByID 根据ID删除一条数据并清理缓存
func (s *SysLogRepo) DeleteOneCacheByIDTx(ctx context.Context, tx *gorm_gen_dao.Query, ID string) error {
	dao := tx.SysLog
	first, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if first == nil {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.ID.Eq(ID)).Delete()
	if err != nil {
		return err
	}
	err = s.DeleteUniqueIndexCache(ctx, []*gorm_gen_model.SysLog{first})
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneByID 根据ID删除一条数据
func (s *SysLogRepo) DeleteOneByID(ctx context.Context, ID string) error {
	dao := gorm_gen_dao.Use(s.db).SysLog
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteOneByID 根据ID删除一条数据
func (s *SysLogRepo) DeleteOneByIDTx(ctx context.Context, tx *gorm_gen_dao.Query, ID string) error {
	dao := tx.SysLog
	_, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
func (s *SysLogRepo) DeleteMultiCacheByIDS(ctx context.Context, IDS []string) error {
	dao := gorm_gen_dao.Use(s.db).SysLog
	list, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Find()
	if err != nil {
		return err
	}
	if len(list) == 0 {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Delete()
	if err != nil {
		return err
	}
	err = s.DeleteUniqueIndexCache(ctx, list)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiCacheByIDS 根据IDS删除多条数据并清理缓存
func (s *SysLogRepo) DeleteMultiCacheByIDSTx(ctx context.Context, tx *gorm_gen_dao.Query, IDS []string) error {
	dao := tx.SysLog
	list, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Find()
	if err != nil {
		return err
	}
	if len(list) == 0 {
		return nil
	}
	_, err = dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Delete()
	if err != nil {
		return err
	}
	err = s.DeleteUniqueIndexCache(ctx, list)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByIDS 根据IDS删除多条数据
func (s *SysLogRepo) DeleteMultiByIDS(ctx context.Context, IDS []string) error {
	dao := gorm_gen_dao.Use(s.db).SysLog
	_, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultiByIDS 根据IDS删除多条数据
func (s *SysLogRepo) DeleteMultiByIDSTx(ctx context.Context, tx *gorm_gen_dao.Query, IDS []string) error {
	dao := tx.SysLog
	_, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// DeleteUniqueIndexCache 删除唯一索引存在的缓存
func (s *SysLogRepo) DeleteUniqueIndexCache(ctx context.Context, data []*gorm_gen_model.SysLog) error {
	keys := make([]string, 0)
	for _, v := range data {
		keys = append(keys, s.cache.Key(cacheSysLogByIDPrefix, v.ID))

	}
	err := s.cache.DelBatch(ctx, keys)
	if err != nil {
		return err
	}
	return nil
}

// FindOneCacheByID 根据ID查询一条数据并设置缓存
func (s *SysLogRepo) FindOneCacheByID(ctx context.Context, ID string) (*gorm_gen_model.SysLog, error) {
	resp := new(gorm_gen_model.SysLog)
	key := s.cache.Key(cacheSysLogByIDPrefix, ID)
	cacheValue, err := s.cache.Fetch(ctx, key, func() (string, error) {
		dao := gorm_gen_dao.Use(s.db).SysLog
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
func (s *SysLogRepo) FindOneByID(ctx context.Context, ID string) (*gorm_gen_model.SysLog, error) {
	dao := gorm_gen_dao.Use(s.db).SysLog
	result, err := dao.WithContext(ctx).Where(dao.ID.Eq(ID)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return result, nil
}

// FindMultiCacheByIDS 根据IDS查询多条数据并设置缓存
func (s *SysLogRepo) FindMultiCacheByIDS(ctx context.Context, IDS []string) ([]*gorm_gen_model.SysLog, error) {
	resp := make([]*gorm_gen_model.SysLog, 0)
	keys := make([]string, 0)
	keyToParam := make(map[string]string)
	for _, v := range IDS {
		key := s.cache.Key(cacheSysLogByIDPrefix, v)
		keys = append(keys, key)
		keyToParam[key] = v
	}
	cacheValue, err := s.cache.FetchBatch(ctx, keys, func(miss []string) (map[string]string, error) {
		params := make([]string, 0)
		for _, v := range miss {
			params = append(params, keyToParam[v])
		}
		dao := gorm_gen_dao.Use(s.db).SysLog
		result, err := dao.WithContext(ctx).Where(dao.ID.In(params...)).Find()
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
			value[s.cache.Key(cacheSysLogByIDPrefix, v.ID)] = string(marshal)
		}
		return value, nil
	})
	if err != nil {
		return nil, err
	}
	for _, v := range cacheValue {
		tmp := new(gorm_gen_model.SysLog)
		err := json.Unmarshal([]byte(v), tmp)
		if err != nil {
			return nil, err
		}
		resp = append(resp, tmp)
	}
	return resp, nil
}

// FindMultiByIDS 根据IDS查询多条数据
func (s *SysLogRepo) FindMultiByIDS(ctx context.Context, IDS []string) ([]*gorm_gen_model.SysLog, error) {
	dao := gorm_gen_dao.Use(s.db).SysLog
	result, err := dao.WithContext(ctx).Where(dao.ID.In(IDS...)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}