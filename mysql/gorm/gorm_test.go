package gorm

import (
	"fmt"
	"testing"
)

type logger0 struct {
}

func (l *logger0) Print(v ...interface{}) {
	fmt.Println(v...)
}

func TestNewGorm(t *testing.T) {

	conf := &Config{
		Driver:  "mysql",
		Host:    "127.0.0.1",
		Port:    "3306",
		Name:    "mysql",
		User:    "root",
		Pwd:     "123456",
		LogMode: true,
	}
	var log logger0
	db, err := New(&log, conf)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	v := make([]interface{}, 0)
	db.Raw("select * from user;").Scan(&v)
}
