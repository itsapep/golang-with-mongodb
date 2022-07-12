package manager

import (
	"context"
	"fmt"
	"time"

	"github.com/itsapep/golang-with-mongodb/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type InfraManager interface {
	DBConn() *mongo.Database
}

type infraManager struct {
	db  *mongo.Database
	cfg config.Config
}

func (i *infraManager) initDb() {
	credential := options.Credential{
		Username: i.cfg.User,
		Password: i.cfg.Password,
	}

	mongoURI := fmt.Sprintf("mongodb://%s:%s", i.cfg.Host, i.cfg.Port)
	clientOptions := options.Client()
	clientOptions.ApplyURI(mongoURI).SetAuth(credential)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	i.db = client.Database(i.cfg.DbName)
}

func (i *infraManager) DBConn() *mongo.Database {
	return i.db
}

func NewInfraManager(config config.Config) InfraManager {
	infra := infraManager{
		cfg: config,
	}
	infra.initDb()
	return &infra
}
