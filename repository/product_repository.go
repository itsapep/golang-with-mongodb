package repository

import (
	"github.com/itsapep/golang-with-mongodb/model"
	"github.com/itsapep/golang-with-mongodb/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository interface {
	Add(newProduct *model.Product) error
	Retrieve() ([]model.Product, error)
}

// 1. FindAllProduct with Pagination
// 2. UpdateProduct
// 3. DeleteProduct
// 4. GetByIdProduct
// 5. GetByProductCategory

type productRepository struct {
	db *mongo.Database
}

// Add implements ProductRepository
func (p *productRepository) Add(newProduct *model.Product) error {
	ctx, cancel := utils.InitContext()
	defer cancel()
	newProduct.Id = primitive.NewObjectID()
	_, err := p.db.Collection("products").InsertOne(ctx, newProduct)
	if err != nil {
		return err
	}
	return nil
}

// Retrieve implements ProductRepository
func (p *productRepository) Retrieve() ([]model.Product, error) {
	var products []model.Product
	ctx, cancel := utils.InitContext()
	defer cancel()
	cursor, err := p.db.Collection("products").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var product model.Product
		err = cursor.Decode(&product)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func NewProductRepository(db *mongo.Database) ProductRepository {
	repo := new(productRepository)
	repo.db = db
	return repo
}
