package redispkg

import (
	"github.com/gomodule/redigo/redis"
)

// RedisPkg :
type RedisPkg struct {
	Host      string `json:"host"`
	Auth      string `json:"auth"`
	DB        int    `json:"db"`
	MaxIdle   int    `json:"max_idle"`
	MaxActive int    `json:"max_active"`
	Pool      *redis.Pool
}

// NewRedisPkg :
func NewRedisPkg(host, auth string, dbIndex, maxIdle, maxActive int) (*RedisPkg, error) {
	redisPkg := RedisPkg{
		Host:      host,
		Auth:      auth,
		DB:        dbIndex,
		MaxIdle:   maxIdle,
		MaxActive: maxActive,
	}

	err := redisPkg.initRedisPool()
	if err != nil {
		return nil, err
	}

	return &redisPkg, nil
}

// initRedisPool :
func (p *RedisPkg) initRedisPool() error {
	conn, err := redis.Dial("tcp", p.Host)
	if err != nil {
		return err
	}

	defer conn.Close()

	if _, err := conn.Do("AUTH", p.Auth); err != nil {
		conn.Close()
		return err
	}

	if _, err := conn.Do("SELECT", p.DB); err != nil {
		conn.Close()
		return err
	}

	p.Pool = &redis.Pool{
		MaxIdle:   p.MaxIdle,
		MaxActive: p.MaxActive, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", p.Host)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", p.Auth); err != nil {
				c.Close()
				return nil, err
			}

			if _, err := c.Do("SELECT", p.DB); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
	}

	return nil
}
