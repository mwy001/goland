package mongodb

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	mongoconf "github.com/mwy001/goland/pkg/conf/mongodb"
	"github.com/mwy001/goland/pkg/log"
)

var (
	mg        *mongoConn
	onceMongo sync.Once
)

type mongoConn struct {
	Client *mongo.Client
	DB     *mongo.Database
}

// Conn returns single instance of neo4j connection
func Conn() *mongoConn {
	onceMongo.Do(func() {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoconf.Config().Mongo.Address))

		if err != nil {
			log.L().Errorf("Mongodb connection error: %v", err)
		}

		mg = &mongoConn{
			Client: client,
			DB:     client.Database(mongoconf.Config().Mongo.DB),
		}
	})

	return mg
}
