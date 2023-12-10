package handlers

import (
	"camera-truth/app/database"
	"camera-truth/app/schemas"
	"context"
	// "fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateCamera (name string) {

	camera:= schemas.Camera{
		Name: name,
	}

	ctx := context.TODO()
	database.UsersCollections.InsertOne(ctx, camera)



	// return nil
}


func FindCamera (name string) (*schemas.Camera, error) {
	ctx := context.TODO()

	var cam schemas.Camera
	filter := bson.M{"name": name}
	// Decode function decodes the retrieved document's contents into the provided struct.
	// To do this, it needs the actual memory address of the struct variable so that it can modify the variable directly in memory.
	err:= database.UsersCollections.FindOne(ctx, filter).Decode(&cam)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
	}
	// fmt.Println("THIS IS X ", x)
	return &cam, nil 
	
}