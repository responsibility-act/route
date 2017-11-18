package controllers

import "github.com/kataras/iris/mvc"

type FolderController struct {
	mvc.C
}

func (c *FolderController) Get() mvc.Result {
	return mvc.View{
		Name: "layout.html",
	}
}
