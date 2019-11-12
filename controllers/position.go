package controllers

import "github.com/kataras/iris/v12/mvc"

type PositionController struct{}

var positionView = mvc.View{
	Name: "layout.html",
}

func (c *PositionController) Get() mvc.Result {
	return positionView
}
