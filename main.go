package main

import (
        "github.com/labstack/echo"
	"fin_test/handler"
        echoMiddleware "github.com/labstack/echo/middleware"
)

func main() {
        // Echo instance
        e := echo.New()

        // Middleware
        e.Use(echoMiddleware.Logger())
        e.Use(echoMiddleware.Recover())

        e.GET("/", func(c echo.Context) error {return HealthCheck(c) })
        e.PUT("/login", func(c echo.Context) error {return handler.UserLogin(c) })

        e.Logger.Fatal(e.Start(":9000"))

}
func HealthCheck(c echo.Context) error {
        c.JSON(200, "health check is ok!")
        return nil
}
