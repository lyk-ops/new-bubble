package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"new-bubble/settings"
)

var (
	DB *gorm.DB
)

func InitMysql(cfg *settings.MysqlConfig) (err error) {
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	//fmt.Printf(dsn)
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}
	return DB.DB().Ping()
}
func Close() {
	DB.Close()
}
