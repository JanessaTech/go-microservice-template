package database

import (
	"errors"
	"fmt"
	"time"

	"github.com/hi-supergirl/go-microservice-template/config"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Create a new database connection
func NewDataBase(logger *zap.Logger, config *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DataBaseConfig.UserName, config.DataBaseConfig.Password, config.DataBaseConfig.Host, config.DataBaseConfig.Schema)
	logger.Sugar().Infoln("dsn =", dsn)
	//dsn := "templateuser:templatepwd@tcp(192.168.0.200:3306)/gin_micro_template?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		return nil, errors.New("cannot connect to mysql")
	} else {
		logger.Sugar().Infoln("mysql is connected")
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, errors.New("failed to return DB()")
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(config.DataBaseConfig.MaxIdleConns)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(config.DataBaseConfig.MaxOpenConns)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
