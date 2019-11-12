package controllers

import "github.com/kataras/iris/v12/mvc"

type MainController struct{}

var mainView = mvc.View{
	Name: "layout.html",
}

func (c *MainController) Get() mvc.Result {
	return mainView
}
