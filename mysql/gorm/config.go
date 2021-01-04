package gorm

import "fmt"

const MysqlDialectStr = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"

type Config struct {
	Dialect string `json:"dialect"` // 填写dialect后，其他参数可以不用填

	Driver string `json:"driver"`
	Host   string `json:"host"`
	Port   string `json:"port"`
	Name   string `json:"name"`
	User   string `json:"user"`
	Pwd    string `json:"pwd"`

	LogMode bool `json:"log_mode"`
}

func (c *Config) GetDialect() string {
	if c.Dialect == "" {
		return fmt.Sprintf(MysqlDialectStr, c.User, c.Pwd, c.Host, c.Port, c.Name)
	}
	return c.Dialect
}
