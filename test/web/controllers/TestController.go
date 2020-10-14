package controllers

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
}