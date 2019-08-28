package db

import "github.com/go-redis/redis"

type Database interface {
	Insert(key, value string) error
	Find(key string) (string, error)
}

type RedisDatabase struct {
	client *redis.Client
}

func NewRedisDatabase(options *redis.Options) Database {
	return &RedisDatabase{
		client: redis.NewClient(options),
	}
}

func (db *RedisDatabase) Insert(key, value string) error {
	if st := db.client.Set(key, value, 0); st.Err() != nil {
		return st.Err()
	}

	return nil
}

func (db *RedisDatabase) Find(key string) (string, error) {
	st := db.client.Get(key)
	if st.Err() != nil {
		return "", st.Err()
	}

	return st.Val(), nil
}
