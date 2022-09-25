package storage

import (
	"context"
	"github.com/gomodule/redigo/redis"
	"github.com/novikov-ai/rate-limiter/configs"
)

type RedisClient struct {
	conf configs.Config
	conn redis.Conn
}

const (
	CommandAdd    = "SAAD"
	CommandRemove = "SREM"
)

func New(conf configs.Config) *RedisClient {
	return &RedisClient{conf: conf}
}

func (rc *RedisClient) Connect(ctx context.Context) error {
	conn, err := redis.DialContext(ctx, "tcp", rc.conf.Server.Host+":"+rc.conf.Server.Port)
	if err != nil {
		return err
	}

	rc.conn = conn

	return nil
}

func (rc *RedisClient) Close() error {
	return rc.conn.Close()
}

func (rc *RedisClient) Add(set, key, value string) error {
	_, err := rc.conn.Do(CommandAdd, set, key, value)
	if err != nil {
		return err
	}

	return nil
}

func (rc *RedisClient) Remove(set, key string) error {
	_, err := rc.conn.Do(CommandRemove, set, key)
	if err != nil {
		return err
	}

	return nil
}
