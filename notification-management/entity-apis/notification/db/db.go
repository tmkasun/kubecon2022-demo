package db

import (
	"context"
	"os"
	"time"

	"github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func init() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("URI")))
	if err != nil {
		panic(err)
	}

	logrus.Info("db connected")

	db = client.Database("notification-service")
}

func GetDB() *mongo.Database {
	return db
}
