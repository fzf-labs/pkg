package cache

import (
	"bufio"
	"context"
	"strings"
	"time"

	"github.com/go-redis/redis/extra/redisotel/v8"
	"github.com/go-redis/redis/v8"
)

type GoRedisConfig struct {
	Addr         string        `json:"addr"`
	Password     string        `json:"password"`
	DB           int           `json:"db"`
	DialTimeout  time.Duration `json:"dialTimeout"`
	WriteTimeout time.Duration `json:"writeTimeout"`
	ReadTimeout  time.Duration `json:"readTimeout"`
}

// NewGoRedis 初始化go-redis客户端
func NewGoRedis(cfg GoRedisConfig) (*redis.Client, error) {
	Client := redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		Password:     cfg.Password,
		DB:           cfg.DB,
		DialTimeout:  cfg.DialTimeout,
		WriteTimeout: cfg.WriteTimeout,
		ReadTimeout:  cfg.ReadTimeout,
	})
	//添加链路追踪
	Client.AddHook(redisotel.NewTracingHook())
	//ping 检测一下
	_, err := Client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return Client, nil
}

// RedisInfo Redis服务信息
func RedisInfo(redis *redis.Client, sections ...string) (res map[string]string) {
	infoStr, err := redis.Info(context.Background(), sections...).Result()
	res = map[string]string{}
	if err != nil {
		return res
	}
	// string拆分多行
	lines, err := stringToLines(infoStr)
	if err != nil {
		return res
	}
	// 解析成Map
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" || strings.HasPrefix(lines[i], "# ") {
			continue
		}
		k, v := stringToKV(lines[i])
		res[k] = v
	}
	return res
}

// stringToLines string拆分多行
func stringToLines(s string) (lines []string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
}

// DBSize 当前数据库key数量
func DBSize(redis *redis.Client) int64 {
	size, err := redis.DBSize(context.Background()).Result()
	if err != nil {
		return 0
	}
	return size
}

// stringToKV string拆分key和val
func stringToKV(s string) (string, string) {
	ss := strings.Split(s, ":")
	if len(ss) < 2 {
		return s, ""
	}
	return ss[0], ss[1]
}
