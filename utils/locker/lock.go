package locker

import (
	"errors"
	"time"

	"github.com/go-redis/redis"
)

var ErrNotHeld = errors.New("not held anymore")

var unlockScript = redis.NewScript(`
	if redis.call("get", KEYS[1]) == ARGV[1] then
		return redis.call("del", KEYS[1])
	else
		return 0
	end`)

// 1: lock success if ok is true and err is nil
// 2: lock failed if ok is false or err is nil
func Lock(client *redis.Client, key string, value interface{}, timeout time.Duration) (ok bool, err error) {
	ok, err = client.SetNX(key, value, timeout).Result()
	return
}

func Unlock(client *redis.Client, key string, value interface{}) (err error) {
	var ret interface{}
	ret, err = unlockScript.Run(client, []string{key}, value).Result()
	if err != nil {
		if err == redis.Nil {
			return ErrNotHeld
		}
		return
	}

	v, ok := ret.(int64)
	if !ok || v != 1 {
		return ErrNotHeld
	}
	return nil
}
