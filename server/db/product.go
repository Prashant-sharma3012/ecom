package db

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Prashant-sharma3012/ecom/tree/main/server/domain/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collectionName = "product"

type ProductRepo struct{}

func NewProductRepo() *ProductRepo {
	return &ProductRepo{}
}

// why should i do this!!!!
// var _ repository.ProductRepository = &ProductRepo{}

func (pr *ProductRepo) SaveProduct(p *entity.Product) (*entity.Product, error) {

	collection := DBClient.Conn.Collection(collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	p.ID = primitive.NewObjectID()

	_, err := collection.InsertOne(ctx, p)
	if err != nil {
		return &entity.Product{}, err
	}

	return p, nil
}

func (pr *ProductRepo) GetAllProduct(skip, limit int64) ([]entity.Product, error) {

	collection := DBClient.Conn.Collection(collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	productCur, err := collection.Find(ctx, bson.M{}, options.Find().SetSkip(skip).SetLimit(limit))
	defer productCur.Close(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var result []bson.M
	var products []entity.Product

	if err = productCur.All(ctx, &result); err != nil {
		log.Fatal(err)
	}

	for _, value := range result {
		var product entity.Product
		bsonBytes, _ := bson.Marshal(value)
		bson.Unmarshal(bsonBytes, &product)
		products = append(products, product)
	}

	return products, nil
}

func (pr *ProductRepo) GetProduct(productId string) (*entity.Product, error) {
	collection := DBClient.Conn.Collection(collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectId, _ := primitive.ObjectIDFromHex(productId)

	var product entity.Product
	filter := bson.M{"id": objectId}

	err := collection.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		return &entity.Product{}, err
	}

	return &product, err
}

func (pr *ProductRepo) UpdateProduct(p *entity.Product) ([]byte, error) {

	collection := DBClient.Conn.Collection(collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.UpdateOne(ctx, bson.D{{"id", p.ID}}, bson.D{{"$set", p}})

	if err != nil {
		return nil, err
	}
	return []byte(p.ID.Hex() + "updated successfully"), nil
}

func (pr *ProductRepo) DeleteProduct(productId string) ([]byte, error) {

	collection := DBClient.Conn.Collection(collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := collection.DeleteOne(ctx, bson.D{{"id", productId}})
	if err != nil {
		return nil, err
	}

	return []byte(strconv.Itoa(int(res.DeletedCount))), nil
}
