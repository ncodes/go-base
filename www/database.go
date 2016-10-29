package www

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
	"time"
)

// Establish a redis connection
func GetRedisPool(addr, password string, db int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr, redis.DialDatabase(db), redis.DialPassword(password))
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

// Get postgres connection
func GetPostgresConn(database, host, username, password, sslmode string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s", host, username, database, sslmode, password))
	if err != nil {
		return nil, err
	}

	db.LogMode(false)
	db.SetLogger(log.New(os.Stdout, "\r\n", 0))

	// // setup postgis and functions
	// err = db.Debug().Raw(`
	// 	CREATE EXTENSION IF NOT EXISTS postgis;
	// 	` + models.POSTGRES_FUNCTIONS_SQL + `
	// `).Error

	// if err != nil {
	// 	return nil, err
	// }

	return db, nil
}
