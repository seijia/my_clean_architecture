package database

import (
	"api_client/config"
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var rdb *redis.Client
var err error

func NewDB() *gorm.DB {

	connect := fmt.Sprintf("%s:%s@%s(%s)/%s?charset=utf8mb4&parseTime=%s&loc=Local",
		config.C.Database.User,
		config.C.Database.Password,
		config.C.Database.Net,
		config.C.Database.Host,
		config.C.Database.DBName,
		config.C.Database.Params.ParseTime,
	)

	// close transaction mode
	DB, err = gorm.Open(mysql.Open(connect), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
	})
	// if not connect,replace with mock
	if err != nil {
		db, _, _ := sqlmock.New()
		DB, _ = gorm.Open(mysql.New(mysql.Config{
			Conn: db,
		}), &gorm.Config{})
		return DB
	}
	return DB
}

func NewRedis() *redis.Client {
	addr := fmt.Sprintf("%s:%s", config.C.Redis.Host, config.C.Redis.Port)
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	return rdb
}
