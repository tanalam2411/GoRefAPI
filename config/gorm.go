package config

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

// Mysql connection configuration
type Mysql struct {
	Path         string `json:"path" yaml:"path"`
	Config       string `json:"config" yaml:"config"`
	Dbname       string `json:"dbname" yaml:"db-name"`
	Username     string `json:"username" yaml:"username"`
	Password     string `json:"password" yaml:"password"`
	MaxIdleConns int    `json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `json:"maxOpenConns" yaml:"max-open-conns"`
}

// Dsn for establishing mysql connection
func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
}

func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(c.Query("page"))
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(c.Query("page_size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
