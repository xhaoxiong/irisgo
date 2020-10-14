package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"github.com/spf13/viper"
	"time"
)

type Database struct {
	Mysql *gorm.DB
}

var DB *Database

func (db *Database) Init() {
	DB = &Database{
		Mysql: GetDB(),
	}
}

func GetDB() *gorm.DB {
	return openMysqlDB(viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.addr"),
		viper.GetString("mysql.name"))
}

func openMysqlDB(username, password, addr, name string) *gorm.DB {

	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		//"Asia/Shanghai"),
		"Local")
	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Println(err, "Database connection failed. Database name: %s", name)
	}

	// set for db connection
	setupDB(db)
	go keepAlive(db)
	return db
}

func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	//db.DB().SetMaxOpenConns(20000) // 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	db.DB().SetMaxIdleConns(2) // 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	db.SingularTable(true)     //设置表名不为负数

	autoMigrate(db)

}

func autoMigrate(db *gorm.DB) {

	if err := db.AutoMigrate().Error;
		err != nil {
		log.Println("自动建表失败", err)
	}
}

func keepAlive(dbc *gorm.DB) {
	for {
		dbc.DB().Ping()
		time.Sleep(60 * time.Second)
	}
}
