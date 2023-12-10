package main

import (
	"fmt"
	"net/http"
    // "service/routes"
    "camera-truth/app/routes"
    "camera-truth/app/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// "net/http"
)

func main() {
    e:= echo.New()
    
    // DB conn
    database.ConnectDB()

    



    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.GET("/", hello)
    e.GET("/name", routes.FindCamera)

    e.POST("/", routes.CreateCamera)

    fmt.Println("Hello, World!")
    e.Logger.Fatal(e.Start(":1323"))
}

func hello(c echo.Context) error {
    // routes.CreateCamera()

    fmt.Println("Created camera")

    return c.String(http.StatusOK, "Hello From GO Lang")
}