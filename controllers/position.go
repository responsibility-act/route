package controllers

import "github.com/kataras/iris/mvc"

type PositionController struct {
	mvc.C
}

func (c *PositionController) Get() mvc.Result {
	return mvc.View{
		Name: "layout.html",
	}
}
