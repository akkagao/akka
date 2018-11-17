package common

import (
	"time"

	"github.com/akkagao/akka/template/service"
	"github.com/akkagao/goutil"
	"github.com/figoxu/Figo"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	db_mail *gorm.DB
)

var (
	GenesisService *service.GenesisService
)

/**
* 启动数据库
 */
func StartDb() {
	pg_mall, err := gorm.Open("postgres", Config.GetString("database.db_mall"))
	goutil.ChkErr(err)
	pg_mall.DB().SetConnMaxLifetime(time.Minute * 5)
	pg_mall.DB().SetMaxIdleConns(0)
	pg_mall.DB().SetMaxOpenConns(5)
	pg_mall.SetLogger(&Figo.GormLog{})
	pg_mall.LogMode(true)
	pg_mall.SingularTable(true)
	pg_mall.Debug().AutoMigrate(mallEntities()...)
	db_mail = pg_mall

}

/**
* 初始化mall 数据库
 */
func mallEntities() (values []interface{}) {
	values = append(values, &service.Genesis{})
	return
}

func StartService() {
	GenesisService = service.NewGenesisService(db_mail)
}

/**
* 启动redis
 */
func StartRedis() {

}
