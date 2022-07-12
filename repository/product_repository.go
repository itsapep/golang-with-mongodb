package repository

import (
	"github.com/itsapep/golang-with-mongodb/model"
	"github.com/itsapep/golang-with-mongodb/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository interface {
	Add(entity interface{}) error
	Retrieve() (interface{}, error)
	Get(by interface{}) (interface{}, error)
	Update(by interface{}, entity interface{}) error
	Delete(entity interface{}) error
}

// 1. FindAllProduct with Pagination
// 2. UpdateProduct
// 3. DeleteProduct
// 4. GetByIdProduct
// 5. GetByProductCategory

type productRepository struct {
	db *mongo.Database
}

// Delete implements ProductRepository
func (p *productRepository) Delete(entity interface{}) error {
	deletedProduct, ok := entity.(*model.Product)
	if ok {
		ctx, cancel := utils.InitContext()
		defer cancel()
		_, err := p.db.Collection("products").DeleteOne(ctx, deletedProduct)
		if err != nil {
			return err
		}
		return nil
	}
	return utils.NewCastInterfaceError("product")
}

// Get implements ProductRepository
func (p *productRepository) Get(by interface{}) (interface{}, error) {
	ctx, cancel := utils.InitContext()
	defer cancel()
	product, err := p.db.Collection("products").Find(ctx, by)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// Update implements ProductRepository
func (p *productRepository) Update(by interface{}, entity interface{}) error {
	updatedProduct, ok := entity.(*model.Product)
	if ok {
		ctx, cancel := utils.InitContext()
		defer cancel()
		_, err := p.db.Collection("products").UpdateOne(ctx, by, updatedProduct)
		if err != nil {
			return err
		}
		return nil
	}
	return utils.NewCastInterfaceError("product")
}

// Add implements ProductRepository
func (p *productRepository) Add(entity interface{}) error {
	newProduct, ok := entity.(*model.Product)
	if ok {
		ctx, cancel := utils.InitContext()
		defer cancel()
		newProduct.Id = primitive.NewObjectID()
		_, err := p.db.Collection("products").InsertOne(ctx, newProduct)
		if err != nil {
			return err
		}
		return nil
	}
	return utils.NewCastInterfaceError("product")
}

// Retrieve implements ProductRepository
func (p *productRepository) Retrieve() (interface{}, error) {
	var products = []model.Product{}
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
