package controllers

import "github.com/kataras/iris/v12/mvc"

type CandidateController struct{}

var candidateView = mvc.View{
	Name: "layout.html",
}

func (c *CandidateController) Get() mvc.Result {
	return candidateView
}
