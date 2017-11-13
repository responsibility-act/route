package controllers

import "github.com/kataras/iris/mvc"

type MainController struct {
	mvc.C
}

func (c *MainController) Get() mvc.Result {
	return mvc.View{
		Name: "layout.html",
	}
}
