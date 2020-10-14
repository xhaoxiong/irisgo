package main

import (
	"github.com/kataras/iris/v12"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	
	"test/web/middleware"
	"test/config"
	"test/models"
	"test/route"
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

	app.Use(middleware.Cors) //是否启用跨域中间件
	app.HandleDir("/public", "./web/views/static") //是否指定静态目录
	//app.RegisterView(iris.HTML("./web/views/admin", ".html")) //是否注册模板
	app.AllowMethods(iris.MethodOptions)
	app.Use(middleware.GetJWT().Serve) //是否启用jwt中间件
	app.Use(middleware.LogMiddle)      //是否启用logrus中间件
	app.Configure(iris.WithOptimizations)

	return app
}

