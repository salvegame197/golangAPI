package main

import (
	"fmt"
	"log"
	"salvegame197/golangApi/src/infra/common/config"
	"salvegame197/golangApi/src/infra/common/persistence/model"
	"salvegame197/golangApi/src/infra/common/services/middleware"
	"salvegame197/golangApi/src/infra/common/util"
	"salvegame197/golangApi/src/infra/user"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getPostgresDBClient(config config.Interface) *gorm.DB {
	host := config.Database().Host()
	user := config.Database().User()
	pass := config.Database().Password()
	database := config.Database().Name()
	dsn := fmt.Sprintf("host=%v user=%v password=%v database=%v", host, user, pass, database)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}

func getRedisClient(config config.Interface) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.Redis().Address(),
		Password: config.Redis().Password(),
	})
}

func main() {

	config := config.Initialize()
	db := getPostgresDBClient(config)
	redisClient := getRedisClient(config)

	model.Initialize(db)
	utils := util.Initialize()
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	router.Group("/api/v1/", middleware.AuthMiddleware())
	// User only can be added by authorized person
	user.Initialize(router, utils, db, redisClient)
	log.Fatal(router.Run(":" + config.Server().Port()))

}
