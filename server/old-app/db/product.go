package db

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collectionName = "product"

type Product struct {
	ID          string   `json:"id" bson:"id"`
	Name        string   `json:"name"  bson:"name"`
	Price       string   `json:"price"  bson:"price"`
	Description string   `json:"description"  bson:"description"`
	Rating      int      `json:"rating"  bson:"rating"`
	Images      []string `json:"images"  bson:"images"`
}

func (db *DB) createproduct(p Product) (Product, error) {

	collection := dbClient.Conn.Database(dbClient.Name).Collection(collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, p)
	if err != nil {
		return Product{}, err
	}
	p.ID = fmt.Sprintf("%v", res.InsertedID)

	return p, nil
}

func (db *DB) listProduct(skip, limit int64) ([]Product, error) {

	collection := dbClient.Conn.Database(dbClient.Name).Collection(collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	productCur, err := collection.Find(ctx, bson.M{}, options.Find().SetSkip(skip).SetLimit(limit))
	defer productCur.Close(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var products []Product

	for productCur.Next(nil) {
		elem := &bson.D{}
		err := productCur.Decode(elem)
		if err != nil {
			log.Fatal("Decode error ", err)
		}

		m := elem.Map()

		product := Product{
			ID:          m["id"].(string),
			Name:        m["name"].(string),
			Price:       m["price"].(string),
			Description: m["description"].(string),
			Rating:      m["rating"].(int),
			Images:      m["images"].([]string),
		}
		products = append(products, product)
	}
	return products, nil

}

func (db *DB) productDetailByID(productId string) (Product, error) {
	collection := dbClient.Conn.Database(dbClient.Name).Collection(collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var product Product
	filter := bson.M{"id": productId}

	err := collection.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		return Product{}, err
	}

	return product, err
}

func UpdateProduct(p Product) ([]byte, error) {

	collection := dbClient.Conn.Database(dbClient.Name).Collection(collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.UpdateOne(ctx, bson.D{{"id", p.ID}}, bson.D{{"$set", p}})

	if err != nil {
		return nil, err
	}
	return []byte(p.ID + "updated successfully"), nil
}

func DeleteProduct(productId string) ([]byte, error) {

	collection := dbClient.Conn.Database(dbClient.Name).Collection(collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := collection.DeleteOne(ctx, bson.D{{"id", productId}})
	if err != nil {
		return nil, err
	}

	return []byte(strconv.Itoa(int(res.DeletedCount))), nil
}
