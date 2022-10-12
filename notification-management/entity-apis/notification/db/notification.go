package db

import (
	"context"

	"github.com/nuwand/kubecon2022-demo/notification-service/helpers"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "notifications"

var collection *mongo.Collection

type Notification struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	UserEmail string             `bson:"userEmail,omitempty" json:"userEmail,omitempty"`
	Message   string             `bson:"message,omitempty" json:"message,omitempty"`
	Type      string             `bson:"type,omitempty" json:"type,omitempty"`
	Read      *bool              `bson:"read,omitempty" json:"read,omitempty"`
}

func init() {
	collection = GetDB().Collection(collectionName)
}

func (obj *Notification) Save(ctx context.Context) error {
	if v, err := helpers.IsValid(obj); !v {
		return err
	}
	readBool := false
	obj.Read = &readBool

	result, err := collection.InsertOne(ctx, &obj)
	if err != nil {
		return err
	}

	obj.ID = (result.InsertedID).(primitive.ObjectID)

	logrus.Infof("Notification:Inserted %s", obj.ID)
	return nil
}

func (obj *Notification) Update(ctx context.Context) error {
	if v, err := helpers.IsValid(obj); !v {
		return err
	}

	filter := Notification{ID: obj.ID}

	update := bson.D{{"$set", obj}}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	logrus.Infof("Notification:Updated %v", obj.ID)
	return nil
}

func (obj *Notification) Get(ctx context.Context) error {
	if v, err := helpers.IsValid(obj); !v {
		return err
	}

	if err := collection.FindOne(ctx, obj).Decode(&obj); err != nil {
		return err
	}
	logrus.Infof("Notification:Retrieved %v", obj.ID)

	return nil
}

func (obj *Notification) Delete(ctx context.Context) error {
	_, err := collection.DeleteOne(ctx, obj)
	if err != nil {
		return err
	}
	logrus.Infof("Notification:Removed %v", obj.ID)
	return nil
}

func GetNotificationList(ctx context.Context, notification Notification) ([]Notification, error) {
	filterCursor, err := collection.Find(ctx, notification)
	if err != nil {
		return nil, err
	}
	var list []Notification
	if err = filterCursor.All(ctx, &list); err != nil {
		return nil, err
	}
	return list, err
}
