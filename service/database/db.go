package database

import (
	"camera-truth/app/utils"
	"context"
	"log"
	"os"

	// "os"
	"fmt"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/options"
)


// type UsersCollections struct {
// 	Collection *mongo.Collection
// }

var DBClient *mongo.Client
var DB *mongo.Database
var UsersCollections *mongo.Collection

var (
	DBName  string = utils.ReturnSafeValue(os.Getenv("DB_NAME"), "camera-truth").(string)
	DBHost  string = utils.ReturnSafeValue(os.Getenv("DB_HOST"), "mongo").(string)
	DBPort  string = utils.ReturnSafeValue(os.Getenv("DB_PORT"), "27017").(string)

)

func getDBConnString () string {
	connectionString:= fmt.Sprintf("mongodb://%s:%s", DBHost, DBPort)

	fmt.Println("db conn: " + connectionString)

	return connectionString

}



func returnDatabase (name string) (*mongo.Database){
	return DBClient.Database(name)
}


func GetCollection (name string) (*mongo.Collection){
   return DB.Collection(name)
}




func ConnectDB () (*mongo.Client) {

	connectionString:= getDBConnString()

	ctx := context.TODO()
	clientOptions:= options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal("Failed to connect to mongo database.")
	}

	fmt.Println("Successfully connected to Database.")

	DBClient = client

	DB = returnDatabase(DBName)

	UsersCollections = GetCollection("users")
	// UsersCollections{Collection: usersCollections}

	return DBClient
}