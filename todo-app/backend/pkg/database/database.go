package database

import (
	"demo-backend-app/pkg/database/schema"
	"demo-backend-app/pkg/env"
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Initialize(db *gorm.DB) {
	db.AutoMigrate(&schema.Todo{})
}

func ConnectMySQL() (*gorm.DB, error) {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&tls=%s",
		env.DBUser, env.DBPass, env.DBHost, env.DBPort, env.DBName, env.TLSOption)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect mysql: %w", err)
	}
	return db, nil
}

func ConnectSQLite() (*gorm.DB, error) {
	// refer https://gorm.io/ja_JP/docs/index.html for details
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db, nil
}
