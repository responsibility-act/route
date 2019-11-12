package main

import (
	"github.com/grayzone/route/controllers"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	app.RegisterView(iris.HTML("./views", ".html").
		Layout("layout.html").
		Reload(true))

	app.HandleDir("/static", "./static")

	app.Use(recover.New())
	app.Use(logger.New())

	mvc.New(app).Configure(func(mvcApp *mvc.Application) {
		mvcApp.Handle(new(controllers.MainController))

		// All the below could be replaced with Controller's methods follows the names:
		// - GetCompany
		// - GetPosition
		// - GetFolder
		// - GetCandidate
		// - GetTask
		mvcApp.Party("/company").Handle(new(controllers.CompanyController))
		mvcApp.Party("/position").Handle(new(controllers.PositionController))
		mvcApp.Party("/folder").Handle(new(controllers.FolderController))
		mvcApp.Party("/candidate").Handle(new(controllers.CandidateController))
		mvcApp.Party("/task").Handle(new(controllers.TaskController))
	})

	app.Run(iris.Addr(":8088"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithCharset("UTF-8"),
		iris.WithOptimizations,
	)
}
