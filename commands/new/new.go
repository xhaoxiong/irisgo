/**
*@Author: haoxiongxiao
*@Date: 2019/4/3
*@Description: CREATE GO FILE commands
 */
package commands

import (
	"github.com/xhaoxiong/irisgo/utils"
	"log"
	"os"
	"path"
	"strings"
)

var conf = `gormlog: true
mysql:
  username: "root"
  password: ""
  addr: "127.0.0.1:3306"
  name: ""
log:
  writers: file,stdout
  logger_level: DEBUG
  logger_file: log
  log_format_text: false
  rollingPolicy: size
  log_rotate_date: 1
  log_rotate_size: 1
  log_backup_count: 7
addr: ":8080"`

var config = `package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	Name string
}

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	if err := c.initConfig(); err != nil {
		return err
	}
	c.initLog()

	c.watchConfig()
	return nil

}

func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name)
	} else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}

	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("irisgo")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func (c *Config) initLog() {
	passLagerCfg := log.PassLagerCfg{
		Writers:        viper.GetString("log.writers"),
		LoggerLevel:    viper.GetString("log.logger_level"),
		LoggerFile:     viper.GetString("log.logger_file"),
		LogFormatText:  viper.GetBool("log.log_format_text"),
		RollingPolicy:  viper.GetString("log.rollingPolicy"),
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}

	log.InitWithConfig(&passLagerCfg)
}

func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("Config file changed: %s\n", e.Name)
	})
}
`

var mysqlInit = `package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/lexkong/log"
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
		log.Errorf(err, "Database connection failed. Database name: %s", name)
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
		log.Error("自动建表失败", err)
	}
}

func keepAlive(dbc *gorm.DB) {
	for {
		dbc.DB().Ping()
		time.Sleep(60 * time.Second)
	}
}
`

var repositories = `package repositories

import (
	"{{.Appname}}/models"
	"github.com/jinzhu/gorm"
)

type TestRepositories struct {
	db *gorm.DB
}

func NewTestRepositories() *TestRepositories {
	return &TestRepositories{db: models.DB.Mysql}
}`

var service = `package service

import (
	"{{.Appname}}/models"
	"{{.Appname}}/repositories"
)

type TestService struct {
	repo *repositories.UserRepositories
}

func NewTestService() *TestService {
	return &TestService{repo: repositories.NewTestRepositories()}
}`

var controllers = ``

var route = ``

var main = `package main

import (
	"{{.AppName}}/config"
	"{{.AppName}}/models"
	"github.com/spf13/pflag"
)

var (
	cfg = pflag.StringP("config", "c", "", "./config.yaml")
)

func main() {
	pflag.Parse()

	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	models.DB.Init()
}`

func CreatedApp(appPath, appName string) {
	log.Println("Creating application...")
	os.MkdirAll(appName, 0755)
	os.Mkdir(path.Join(appName, "conf"), 0755)
	os.Mkdir(path.Join(appName, "config"), 0755)
	os.Mkdir(path.Join(appName, "models"), 0755)
	os.Mkdir(path.Join(appName, "repositories"), 0755)
	os.Mkdir(path.Join(appName, "service"), 0755)
	utils.WriteToFile(path.Join(appName, "conf", "config.yaml"), conf)
	utils.WriteToFile(path.Join(appName, "config", "config.go"), config)
	utils.WriteToFile(path.Join(appName, "models", "init.go"), mysqlInit)
	utils.WriteToFile(path.Join(appName, "service", "TestService.go"), strings.Replace(service, "{{.Appname}}", appName, -1))
	utils.WriteToFile(path.Join(appName, "repositories", "TestRepo.go"), strings.Replace(repositories, "{{.Appname}}", appName, -1))
	utils.WriteToFile(path.Join(appName, "main.go"), strings.Replace(main, "{{.AppName}}", appName, -1))

	log.Println("new application successfully created!")
}
