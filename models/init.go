package models

import (
	"log"

	"getcharzp.cn/define"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB = Init()

var RDB = InitRedisDB()

func Init() *gorm.DB {
	dsn := define.MysqlDNS + "/oj?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("gorm Init Error : ", err)
	}
	return db
}

func InitRedisDB() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
