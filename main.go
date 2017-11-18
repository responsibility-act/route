package main

import (
	"github.com/grayzone/route/controllers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()
	app.Configure(iris.WithoutServerError(iris.ErrServerClosed), iris.WithCharset("UTF-8"), iris.WithOptimizations)
	templates := iris.HTML("./views", ".html")
	templates.Layout("layout.html")
	templates.Reload(true)
	app.RegisterView(templates)
	app.StaticWeb("/static", "./static")

	app.Use(recover.New())
	app.Use(logger.New())

	app.Logger().SetLevel("debug")

	app.Controller("/", new(controllers.MainController))
	app.Controller("/company", new(controllers.CompanyController))
	app.Controller("/position", new(controllers.PositionController))
	app.Controller("/folder", new(controllers.FolderController))
	app.Controller("/candidate", new(controllers.CandidateController))
	app.Controller("/task", new(controllers.TaskController))

	app.Run(iris.Addr(":8088"))
}
