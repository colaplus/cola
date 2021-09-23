package session

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"sync"
)

const (
	SessionFlagNone = iota
	SessionFlagModify
	SessionFlagLoad
)

type RedisSession struct {
	id  string
	pool       *redis.Pool
	data map[string]interface{}
	rwlock     sync.RWMutex
	flag       int
}

func NewRedisSession(id string, pool *redis.Pool) Session {
	s := &RedisSession{
		id:   id,
		data: make(map[string]interface{}, 8),
		flag: SessionFlagNone,
		pool: pool,
	}

	return s
}

func (r *RedisSession) Id() string {
	return r.id
}

func (r *RedisSession) Set(key string, value interface{}) (err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()

	r.data[key] = value
	r.flag = SessionFlagModify
	return
}

func (r *RedisSession) loadFromRedis() (err error) {
	conn := r.pool.Get()
	reply, err := conn.Do("GET", r.id)
	if err != nil {
		return
	}

	data, err := redis.String(reply, err)
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(data), &r.data)
	if err != nil {
		return
	}

	return
}

func (r *RedisSession) Get(key string) (result interface{}, err error) {
	r.rwlock.RLock()
	defer r.rwlock.Unlock()

	// 实现了一个延迟加载的功能
	if r.flag == SessionFlagNone {
		// 该session还没有加载，那么就从redis中加载数据
		err = r.loadFromRedis()
		if err != nil {
			return
		}
	}

	result, ok := r.data[key]
	if !ok {
		err = ErrSessionNotExist
		return
	}

	return
}
func (r *RedisSession) Del(key string) (err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()

	r.flag = SessionFlagModify
	delete(r.data, key)
	return
}
func (r *RedisSession) Save() (err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()

	if r.flag != SessionFlagModify {
		return
	}

	data, err := json.Marshal(r.data)
	if err != nil {
		return
	}

	conn := r.pool.Get()
	_, err = conn.Do("SET", r.id, string(data))
	if err != nil {
		return
	}
	return
}

func (r *RedisSession) IsModify() bool {
	if r.flag == SessionFlagModify {
		return true
	}

	return false
}