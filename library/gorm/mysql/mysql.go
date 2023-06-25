package libmysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func NewMysql(addr, username, pwd, dbname, charset string, port int, life time.Duration) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local", username, pwd, addr, port, dbname, charset)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		//Logger:               nil,
	})
	if err != nil {
		return nil, err
	}
	_ = db.Callback().Query().Before("gorm:query").Register("disable_raise_record_not_found", func(gormDB *gorm.DB) {
		gormDB.Statement.RaiseErrorOnNotFound = false
	})
	d, err := db.DB()
	if err != nil {
		return nil, err
	}

	d.SetConnMaxIdleTime(time.Second * 30) // 最大空闲时间
	d.SetConnMaxLifetime(life)             // 最大连接时间
	d.SetMaxIdleConns(10)                  // 最大空闲连接
	d.SetMaxOpenConns(128)                 // 最大连接
	return
}
