package misc

import (
	"time"

	log "code.google.com/p/log4go"
	"github.com/garyburd/redigo/redis"
)

type MessageQueue struct {
	Queue *redis.Pool
}

var MQ *MessageQueue

func InitQueue() {
	MQ = &MessageQueue{
		Queue: &redis.Pool{
			MaxActive:   Conf.Redis.MaxActive,
			MaxIdle:     Conf.Redis.MaxIdle,
			IdleTimeout: Conf.Redis.IdleTimeout * time.Second,
			Dial: func() (redis.Conn, error) {
				c, err := redis.DialTimeout(
					"tcp", Conf.Redis.Addr,
					Conf.Redis.ConnectTimeout*time.Second,
					Conf.Redis.ReadTimeout*time.Second,
					Conf.Redis.WriteTimeout*time.Second)
				if err != nil {
					log.Warn("failed to connect Redis, (%s)", err)
					return nil, err
				}
				if Conf.Redis.Passwd != "" {
					if _, err := c.Do("AUTH", Conf.Redis.Passwd); err != nil {
						log.Warn("failed to auth Redis, (%s)", err)
						return nil, err
					}
				}
				//log.Debugf("connected with Redis (%s)", server)
				return c, err
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
		},
	}
}

func (this *MessageQueue) do(commandName string, args ...interface{}) (interface{}, error) {
	var conn redis.Conn
	i := DEFAULT_RETRY
	for ; i > 0; i-- {
		conn = this.StatStore.Get()
		err := conn.Err()
		if err == nil {
			break
		} else {
		}
		time.Sleep(DEFAULT_RETRY_INTERVAL)
	}

	if i == 0 || conn == nil {
		return nil, fmt.Errorf("failed to find a useful redis conn")
	} else {
		ret, err := conn.Do(commandName, args...)
		conn.Close()
		return ret, err
	}
}
