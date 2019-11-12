package controllers

import "github.com/kataras/iris/v12/mvc"

type FolderController struct{}

var folderView = mvc.View{
	Name: "layout.html",
}

func (c *FolderController) Get() mvc.Result {
	return folderView
}
