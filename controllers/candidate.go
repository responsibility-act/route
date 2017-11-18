package controllers

import "github.com/kataras/iris/mvc"

type CandidateController struct {
	mvc.C
}

func (c *CandidateController) Get() mvc.Result {
	return mvc.View{
		Name: "layout.html",
	}
}
