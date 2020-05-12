/**
 * @Author: xiaoxiao
 * @Description:
 * @File:  controllers
 * @Version: 1.0.0
 * @Date: 2020/5/12 2:11 下午
 */
package mvc_template

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
