package models

import (
	"fmt"
	"link-gin-db/config"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func Init(mysqlConf *config.MysqlConfig) {
	var (
		err error
	)

	str := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", mysqlConf.User, mysqlConf.Password, mysqlConf.Addr, mysqlConf.Db)
	DB, err = gorm.Open("mysql", str)
	if err != nil {
		panic(err)
	}

	DB.LogMode(true)
	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(mysqlConf.MaxIdleConnS)
	DB.DB().SetMaxOpenConns(mysqlConf.MaxOpenConnS)

	DB.AutoMigrate(&TBUser{})

	return
}

type UserList struct {
	Lock  *sync.Mutex
	IdMap map[uint]*TBUser
}

func checkPageSize(offset, limit int32) (int32, int32) {
	if offset < 0 {
		offset = 0
	}
	if limit > 20 {
		limit = 20
	}
	return offset, limit
}
