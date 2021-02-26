package initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"ms-inventory/config"
	modelsV1 "ms-inventory/pkg/models/inventory/v1"
)

// Establishing mysql connection and returning mysql conn object
func GormMysql() *gorm.DB {
	m := config.Mysql{
		Path:         "127.0.0.1:3306",
		Config:       "charset=utf8mb4&parseTime=True&loc=Local",
		Dbname:       "first_go",
		Username:     "testuser",
		Password:     "password123",
		MaxIdleConns: 10,
		MaxOpenConns: 100,
	}
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN: dsn,
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig()); err != nil {
		panic("failed to connect database")
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

// Gorm configuration
func gormConfig() *gorm.Config {
	return &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
}

// Mysql migration
func MysqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		modelsV1.InventoryCategory{},
	)
	if err != nil {
		log.Fatalf("Failed to register table: %v", err)
	}
	log.Printf("Tables registered")
}
