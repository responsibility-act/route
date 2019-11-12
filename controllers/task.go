package controllers

import "github.com/kataras/iris/v12/mvc"

type TaskController struct{}

var taskView = mvc.View{
	Name: "layout.html",
}

func (c *TaskController) Get() mvc.Result {
	return taskView
}
