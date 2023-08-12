package database

import (
	"fmt"
	"github.com/systoms/wegin/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase(databaseConfig config.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		databaseConfig.User,
		databaseConfig.Password,
		databaseConfig.Host,
		databaseConfig.Port,
		databaseConfig.DatabaseName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
