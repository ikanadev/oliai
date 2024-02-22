package rest

import "github.com/labstack/echo/v4"

func ServeRoutes() {
	app := echo.New()
	app.GET("/status", func(c echo.Context) error {
		return c.JSON(200, "running")
	})
	app.Logger.Fatal(app.Start(":4000"))
}
