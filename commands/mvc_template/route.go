/**
 * @Author: xiaoxiao
 * @Description:
 * @File:  route
 * @Version: 1.0.0
 * @Date: 2020/5/12 2:12 下午
 */
package mvc_template

var route = `package route

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"{{.Appname}}/web/controllers"
)

func InitRouter(app *iris.Application) {

	mvc.New(app.Party("/")).Handle(controllers.NewTestController())

}`
