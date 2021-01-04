package gorm

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// 数据库的logger接口，如果想查看，gorm的sql日志，倾实现logger
type logger interface {
	Print(v ...interface{})
}

func New(logger logger, c *Config) (*gorm.DB, error) {
	dialStr := c.GetDialect()
	if c.Driver == "" {
		c.Driver = "mysql"
	}
	db, err := gorm.Open(c.Driver, dialStr)
	if err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(100)
	db.DB().SetConnMaxLifetime(120)
	db.DB().SetMaxOpenConns(10000)

	if c.LogMode {
		db.LogMode(c.LogMode)
		db.SetLogger(logger)
	}

	return db, nil
}
