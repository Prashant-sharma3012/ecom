package db

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Prashant-sharma3012/ecom/tree/main/server/domain/entity"

	"github.com/Prashant-sharma3012/ecom/tree/main/server/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collectionName = "product"

type ProductRepo struct{}

func NewProductRepo() *ProductRepo {
	return &ProductRepo{}
}

var _ repository.ProductRepository = &ProductRepo{}

func (pr *ProductRepo) SaveProduct(p *entity.Product) (*entity.Product, error) {

	collection := dbClient.Conn.Collection(collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, p)
	if err != nil {
		return &entity.Product{}, err
	}
	p.ID = fmt.Sprintf("%v", res.InsertedID)

	return p, nil
}

func (pr *ProductRepo) GetAllProduct(skip, limit int64) ([]entity.Product, error) {

	collection := dbClient.Conn.Collection(collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	productCur, err := collection.Find(ctx, bson.M{}, options.Find().SetSkip(skip).SetLimit(limit))
	defer productCur.Close(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var products []entity.Product

	for productCur.Next(nil) {
		elem := &bson.D{}
		err := productCur.Decode(elem)
		if err != nil {
			log.Fatal("Decode error ", err)
		}

		m := elem.Map()

		product := entity.Product{
			ID:          m["id"].(string),
			Name:        m["name"].(string),
			Price:       m["price"].(int64),
			Description: m["description"].(string),
			Rating:      m["rating"].(int),
			Images:      m["images"].([]string),
			Seller:      m["seller"].(string),
			CreatedBy:   m["createdBy"].(string),
			UpdatedBy:   m["updatedBy"].(string),
			UpdatedAt:   m["updatedAt"].(time.Time),
			CreatedAt:   m["createdAt"].(time.Time),
		}
		products = append(products, product)
	}
	return products, nil

}

func (pr *ProductRepo) GetProduct(productId string) (*entity.Product, error) {
	collection := dbClient.Conn.Collection(collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var product entity.Product
	filter := bson.M{"id": productId}

	err := collection.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		return &entity.Product{}, err
	}

	return &product, err
}

func (pr *ProductRepo) UpdateProduct(p *entity.Product) ([]byte, error) {

	collection := dbClient.Conn.Collection(collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.UpdateOne(ctx, bson.D{{"id", p.ID}}, bson.D{{"$set", p}})

	if err != nil {
		return nil, err
	}
	return []byte(p.ID + "updated successfully"), nil
}

func (pr *ProductRepo) DeleteProduct(productId string) ([]byte, error) {

	collection := dbClient.Conn.Collection(collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := collection.DeleteOne(ctx, bson.D{{"id", productId}})
	if err != nil {
		return nil, err
	}

	return []byte(strconv.Itoa(int(res.DeletedCount))), nil
}
