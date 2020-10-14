package route

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"test/web/controllers"
)

func InitRouter(app *iris.Application) {

	mvc.New(app.Party("/")).Handle(controllers.NewTestController())

}