package repository

import (
	"ginblog/pkg/log"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
	// rdb    *redis.Client
	logger *log.Logger
}

func NewRepository(logger *log.Logger, db *gorm.DB) *Repository {
	return &Repository{
		db: db,
		// rdb:    rdb,
		logger: logger,
	}
}

func NewDb(conf *viper.Viper) *gorm.DB {
	// TODO: init db
	db, err := gorm.Open(mysql.Open(conf.GetString("data.mysql.user")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
	// return &gorm.DB{}
}
