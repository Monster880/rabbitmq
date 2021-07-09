package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.RegisterView(iris.HTML("./web/views", ".html"))

	app.Run(
		iris.Addr("localhost:8080"),
	)
}
