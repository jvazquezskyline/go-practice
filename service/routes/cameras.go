package routes

import (
	"fmt"
	"net/http"

	"camera-truth/app/handlers"

	"github.com/labstack/echo/v4"
	// "context"
	// "go.mongodb.org/mongo-driver/bson"
)



func FindCamera (c echo.Context) (error) {
	req:= c.Request()
	
	headers:= req.Header
	
	name:= headers.Get("name")
	fmt.Println("Retrieving camera........... ", name)

	camera, err := handlers.FindCamera(name)


	if err != nil {
		return c.String(http.StatusNotFound, "Camera not found.")

	}


	return c.JSON(http.StatusOK, camera)

}

func CreateCamera (c echo.Context) error {
	req:= c.Request()

	headers:= req.Header

	
	name:= headers.Get("name")

	fmt.Println("Name of the camera headers ",  name)
	// camera:=controllers.Camera{Name: name}


	
	fmt.Println("About to create a camera")
	handlers.CreateCamera(name)
	// controllers.CreateCamera(camera)
	return nil
}
