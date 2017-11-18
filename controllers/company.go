package controllers

import "github.com/kataras/iris/mvc"

type CompanyController struct {
	mvc.C
}

func (c *CompanyController) Get() mvc.Result {
	return mvc.View{
		Name: "layout.html",
	}
}
