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
	app.Use(middleware.LogMiddle)//是否启用logrus中间件
	app.Configure(iris.WithOptimizations)

	return app
}

