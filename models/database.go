package models

import (
	"fmt"
	"time"

	"github.com/importcjj/neighbor/utils/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	username = "root"
	password = "jiaju"
	host     = "localhost"
	port     = 3306
	dbname   = "neighbor"
	maxIdle  = 50
	maxOpen  = 100
)

var (
	db *gorm.DB
)

func init() {
	db = CreateDB(username, password, host, port, dbname, maxIdle, maxOpen)

	// 自动建表.
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Location{})

}

// CreateDB 初始化MYSQL实例
func CreateDB(username, password, host string, port int, dbname string, maxIdle, maxOpen int) *gorm.DB {
	connstr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		dbname,
	)

	log.Warnf("try to connect to MYSQL %s:%d", host, port)
	database, err := gorm.Open("mysql", connstr)
	if err != nil {
		log.Errorf("failed to connect MYSQL %s:%d/%s: %s", host, port, dbname, err.Error())
		return database
	}
	log.Infof("connected to MYSQL %s:%d/%s", host, port, dbname)

	database.DB().SetMaxIdleConns(maxIdle)
	database.DB().SetMaxOpenConns(maxOpen)
	// database.DB().SetConnMaxLifetime(d)

	return database
}

// Session 返回一个Transation
// Must remember to commit or rollback
func Session() *gorm.DB {
	return db.Begin()
}

// DB 返回一个MYSQL的空闲连接
func DB() *gorm.DB {
	return db
}

// TimeMixin mixin
type TimeMixin struct {
	UpdatedAt time.Time `gorm:"column:updated_at;type:TIMESTAMP"`
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP"`
}
