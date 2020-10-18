/**
 * @Author: xiaoxiao
 * @Description:
 * @File:  init
 * @Version: 1.0.0
 * @Date: 2020/5/12 2:11 下午
 */
package api_template

var mysqlInit = `
package models

 import (
    "database/sql"
    "fmt"
    "github.com/casbin/casbin/v2"
    gormadapter "github.com/casbin/gorm-adapter/v3"
    "github.com/sirupsen/logrus"
    "github.com/spf13/viper"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "log"
    "time"
 )

 type Database struct {
    Mysql *gorm.DB
 }

 var DB *Database
 var Enforcer *casbin.Enforcer

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
    db, err := gorm.Open(mysql.Open(config), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Silent),
    })
    if err != nil {
        logrus.Println(err, "Database connection failed. Database name: %s", name)
		return db
    }
    dbc, err := db.DB()
    setupDB(dbc)
    go keepAlive(dbc)
    //autoMigrate(db)
    a, err := gormadapter.NewAdapterByDB(db)
    if err != nil {
        logrus.Println("新建适配器失败", err)
    }
    Enforcer, _ = casbin.NewEnforcer("conf/rbac_model.conf", a)
    Enforcer.LoadPolicy()
    return db
 }

 func setupDB(db *sql.DB) {
    db.SetMaxOpenConns(-1) // 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
    db.SetMaxIdleConns(2)  // 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
 }

 func autoMigrate(db *gorm.DB) {

    if err := db.Migrator().CreateTable().Error();
        err != "" {
        log.Println("自动建表失败", err)
    }
 }

 func keepAlive(dbc *sql.DB) {
    for {
        dbc.Ping()
        time.Sleep(60 * time.Second)
    }
 }

`

var traceLogger = `package models

import "time"

type TraceLogger struct {
	ReqTime   time.Time
	ReqUri    string
	ReqMethod string
	Proto     string
	UserAgent string
	Referer   string
	Length int64
}

func NewLogger() *TraceLogger {
	return &TraceLogger{}
}`
