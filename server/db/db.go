package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DBClient *DB

type DB struct {
	Conn    *mongo.Database
	Name    string
	Product *ProductRepo
}

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("DB connection failed")
	}

	DBClient = &DB{
		Conn:    client.Database("ecom"),
		Name:    "ecom",
		Product: NewProductRepo(),
	}

}

func GetDBClient() *DB {
	return DBClient
}
