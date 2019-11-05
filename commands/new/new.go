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

`

var mysqlInit = `package models

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
	"{{.Appname}}/repositories"
)

type TestService struct {
	repo *repositories.TestRepositories
}

func NewTestService() *TestService {
	return &TestService{repo: repositories.NewTestRepositories()}
}`

var controllers = `package controllers

import "github.com/kataras/iris/v12"

type TestController struct {
	Ctx iris.Context
	Common
}

func NewTestController() *TestController {
	return &TestController{}
}

func (this *TestController) Get() {
	this.ReturnSuccess()
}`

var common = `
package controllers

import "github.com/kataras/iris/v12"

type Common struct {
	Ctx iris.Context
}

func (this *Common) ReturnJson(status int, message string, args ...interface{}) {
	result := make(map[string]interface{})
	result["code"] = status
	result["message"] = message

	key := ""

	for _, arg := range args {
		switch arg.(type) {
		case string:
			key = arg.(string)

		default:
			result[key] = arg
		}
	}
	this.Ctx.JSON(result)
	this.Ctx.StopExecution()
	return
}

func (this *Common) ReturnSuccess(args ...interface{}) {
	result := make(map[string]interface{})
	result["code"] = 10000
	result["message"] = "success"
	key := ""
	for _, arg := range args {
		switch arg.(type) {
		case string:
			key = arg.(string)
		default:
			result[key] = arg
		}
	}
	this.Ctx.JSON(result)
	this.Ctx.StopExecution()
	return
}

`

var jwt = `
package middleware

import (
"github.com/dgrijalva/jwt-go"
jwtmiddleware "github.com/iris-contrib/middleware/jwt"
"github.com/kataras/iris/v12/context"

"fmt"
"time"
)

var JwtAuthMiddleware = jwtmiddleware.New(jwtmiddleware.Config{
	ValidationKeyGetter: validationKeyGetterFuc,
	SigningMethod:       jwt.SigningMethodHS256,
	Expiration:          true,
	Extractor:           extractor,
}).Serve

const jwtKey = "{{.Appname}}"

var validationKeyGetterFuc = func(token *jwt.Token) (interface{}, error) {
	return []byte(jwtKey), nil
}

var extractor = func(ctx context.Context) (string, error) {
	authHeader := ctx.GetHeader("token")
	if authHeader == "" {
		return "", nil // No error, just no token
	}

	return authHeader, nil
}

//注册jwt中间件
func GetJWT() *jwtmiddleware.Middleware {
	jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
		//这个方法将验证jwt的token
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			//自己加密的秘钥或者说盐值
			return []byte(jwtKey), nil
		},
		//加密的方式
		SigningMethod: jwt.SigningMethodHS256,
		//验证未通过错误处理方式
		//ErrorHandler: func(context.Context, string)
		ErrorHandler: func(ctx context.Context, e error) {
			ctx.Next()
		},
	})
	return jwtHandler
}

//生成token
func GenerateToken(msg string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"msg": msg,                                                      //openid
		"iss": "iris_{{.Appname}}",                                      //签发者
		"iat": time.Now().Unix(),                                        //签发时间
		"jti": "9527",                                                   //jwt的唯一身份标识，主要用来作为一次性token,从而回避重放攻击。
		"exp": time.Now().Add(10 * time.Hour * time.Duration(1)).Unix(), //过期时间
	})

	tokenString, _ := token.SignedString([]byte(jwtKey))
	fmt.Println("签发时间：", time.Now().Unix())
	fmt.Println("到期时间：", time.Now().Add(10*time.Hour*time.Duration(1)).Unix())
	return tokenString
}
`

var route = `package route

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"{{.Appname}}/web/controllers"
)

func InitRouter(app *iris.Application) {

	mvc.New(app.Party("/")).Handle(controllers.NewTestController())

}`

var main = `package main

import (
	"github.com/kataras/iris/v12"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	
	"{{.Appname}}/web/middleware"
	"{{.Appname}}/config"
	"{{.Appname}}/models"
	"{{.Appname}}/route"
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
	app := newApp()

	route.InitRouter(app)

	app.Run(iris.Addr(viper.GetString("addr")))
}

func newApp() *iris.Application {
	app := iris.New()
	//crs := cors.New(cors.Options{
	//	AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
	//	AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	//	AllowCredentials: true,
	//	AllowedHeaders:   []string{"*"},
	//})
	//
	//app.Use(crs) //
	//app.StaticWeb("/assets", "./web/views/admin/assets")
	//app.RegisterView(iris.HTML("./web/views/admin", ".html"))
	app.AllowMethods(iris.MethodOptions)
	app.Use(middleware.GetJWT().Serve)//是否启用jwt中间件
	app.Configure(iris.WithOptimizations)

	return app
}

`

func CreatedApp(appPath, appName string) {
	log.Println("Creating application...")
	os.MkdirAll(appName, 0755)
	os.Mkdir(path.Join(appName, "conf"), 0755)
	os.Mkdir(path.Join(appName, "config"), 0755)
	os.Mkdir(path.Join(appName, "models"), 0755)
	os.Mkdir(path.Join(appName, "route"), 0755)
	os.Mkdir(path.Join(appName, "repositories"), 0755)
	os.Mkdir(path.Join(appName, "service"), 0755)
	os.MkdirAll(path.Join(appName, "/web/controllers"), 0755)
	os.MkdirAll(path.Join(appName, "/web/middleware"), 0755)
	utils.WriteToFile(path.Join(appName, "conf", "config.yaml"), conf)
	utils.WriteToFile(path.Join(appName, "config", "config.go"), config)
	utils.WriteToFile(path.Join(appName, "models", "init.go"), mysqlInit)
	utils.WriteToFile(path.Join(appName, "service", "TestService.go"), strings.Replace(service, "{{.Appname}}", appName, -1))
	utils.WriteToFile(path.Join(appName, "repositories", "TestRepo.go"), strings.Replace(repositories, "{{.Appname}}", appName, -1))
	utils.WriteToFile(path.Join(appName, "route", "route.go"), strings.Replace(route, "{{.Appname}}", appName, -1))
	utils.WriteToFile(path.Join(appName, "/web/controllers", "TestController.go"), controllers)
	utils.WriteToFile(path.Join(appName, "/web/controllers", "Common.go"), common)
	utils.WriteToFile(path.Join(appName, "/web/middleware", "jwt.go"), strings.Replace(jwt, "{{.Appname}}", appName, -1))

	utils.WriteToFile(path.Join(appName, "main.go"), strings.Replace(main, "{{.Appname}}", appName, -1))

	log.Println("new application successfully created!")
}
