package cache

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"silent-cxl.top/app/vars"
	"silent-cxl.top/library/hlog"
	libredis "silent-cxl.top/library/redis"
	"strings"
	"time"
)

type DBCache struct {
	db     *gorm.DB
	driver *libredis.Redis
	debug  bool
	expire time.Duration
}

func NewCache(db *gorm.DB) *DBCache {
	return &DBCache{
		driver: vars.Redis,
		db:     db,
		expire: time.Second * 3600 * 24,
		debug:  vars.YmlConfig.GetBool("Debug"),
	}
}

// QueryRows 条件查找 v 必需是指针类型参数
func (dc *DBCache) QueryRows(k string, v interface{}, f func(_db *gorm.DB, _v interface{}, _bys ...interface{}) error, bys ...interface{}) (err error) {
	if dc.debug {
		return f(dc.db, v, bys...)
	}
	var keys int64 = 0
	cacheKey := dc.autoBuildKey(k, bys...)
	if keys, err = dc.driver.Exists(dc.driver.Ctx, cacheKey).Result(); err != nil {
		return err
	} else if keys == 0 {
		if err = f(dc.db, v, bys...); err != nil {
			return err
		}
		_ = dc.setCache(cacheKey, v)
		return nil
	}
	result, err := dc.driver.Get(dc.driver.Ctx, cacheKey).Bytes()
	if err != nil {
		return err
	}
	if err = json.Unmarshal(result, v); err != nil {
		return err
	}
	return nil
}

func (dc *DBCache) RemoveBy(key string, by interface{}) (err error) {
	_, err = dc.driver.Del(dc.driver.Ctx, dc.autoBuildKey(key, by)).Result()
	return err
}

func (dc *DBCache) setCache(key string, v interface{}) (err error) {
	marshal, err := json.Marshal(v)
	if err != nil {
		return err
	}
	_, err = dc.driver.Set(dc.driver.Ctx, key, string(marshal), dc.expire).Result()
	return err
}

func (dc *DBCache) autoBuildKey(key string, ks ...interface{}) string {
	s := []string{"%s:%s"}
	vs := []interface{}{dc.driver.Prefix, key}
	for _, k := range ks {
		s = append(s, "%v")
		vs = append(vs, k)
	}
	return fmt.Sprintf(strings.Join(s, ":"), vs...)
}

// SetExpire 手动设置过期时间，单位 秒
func (dc *DBCache) SetExpire(sec int) {
	dc.expire = time.Second * time.Duration(sec)
}

// ExpireKeys 过期被修改的缓存
func (dc *DBCache) ExpireKeys(keys []string) {
	if dc.debug {
		return
	}
	_keys := make([]string, len(keys))
	for i, key := range keys {
		_keys[i] = dc.autoBuildKey(key)
	}
	hlog.NewLog(logrus.InfoLevel, "cache-clear-many").Log(map[string]interface{}{
		"keys": strings.Join(_keys, ","),
	}, "redis expire key")
	dc.driver.Del(dc.driver.Ctx, _keys...)
}

// ExpireKey 过期被修改的缓存（单个）
func (dc *DBCache) ExpireKey(key string) {
	if dc.debug {
		return
	}
	hlog.NewLog(logrus.InfoLevel, "cache-clear-one").Log(map[string]interface{}{"key": key}, "redis expire key")
	dc.driver.Del(dc.driver.Ctx, dc.autoBuildKey(key))
}
