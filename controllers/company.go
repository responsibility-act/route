package controllers

import "github.com/kataras/iris/v12/mvc"

type CompanyController struct{}

var companyView = mvc.View{
	Name: "layout.html",
}

func (c *CompanyController) Get() mvc.Result {
	return companyView
}
