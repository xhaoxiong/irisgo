/**
 * @Author: xiaoxiao
 * @Description:
 * @File:  commom
 * @Version: 1.0.0
 * @Date: 2020/5/12 2:12 下午
 */
package commands


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
