package redis

import (
	"context"
	"errors"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	DefaultTTL   = 3600 * time.Second
	DefaultValue = "lock"
)

// var rdb *redis.Client
// var ctx = context.Background()
var (
	Store *persist.RedisStore
	RDB   *redis.Client
	ctx   = context.Background()
)

// Lock 实现了 Redis 原子锁.
type Lock struct {
	client    *redis.Client
	lockKey   string
	lockValue string
	ttl       time.Duration
	ctx       context.Context
}

// NewLock 初始化一个 Redis 锁.
func NewLock(lockKey string) *Lock {
	return &Lock{
		client:    RDB,
		lockKey:   lockKey,
		lockValue: DefaultValue,
		ttl:       DefaultTTL,
		ctx:       ctx,
	}
}

// SetLockValue 设置锁的值.
func (r *Lock) SetLockValue(value string) *Lock {
	r.lockValue = value
	return r
}

// SetTTL 设置锁的过期时间.
func (r *Lock) SetTTL(ttl time.Duration) *Lock {
	r.ttl = ttl
	return r
}

// CheckAndSetLock 检查并设置锁.
func (r *Lock) CheckAndSetLock() (bool, error) {
	ok, err := r.client.SetNX(r.ctx, r.lockKey, r.lockValue, r.ttl).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return false, err
	}
	return ok, nil
}

// SetLock 强制设置锁.
func (r *Lock) SetLock() error {
	_, err := r.client.Set(r.ctx, r.lockKey, r.lockValue, r.ttl).Result()
	return err
}

// Unlock 解锁.
func (r *Lock) Unlock() error {
	return r.client.Eval(r.ctx, `
		if redis.call("get", KEYS[1]) == ARGV[1] then
			return redis.call("del", KEYS[1])
		else
			return 0
		end
	`, []string{r.lockKey}, r.lockValue).Err()
}

// IsLockExist 检查锁是否存在.
func (r *Lock) IsLockExist() (bool, error) {
	exists, err := r.client.Exists(r.ctx, r.lockKey).Result()
	if err != nil {
		return false, err
	}
	return exists == 1, nil
}
