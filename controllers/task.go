package controllers

import "github.com/kataras/iris/mvc"

type TaskController struct {
	mvc.C
}

func (c *TaskController) Get() mvc.Result {
	return mvc.View{
		Name: "layout.html",
	}
}
