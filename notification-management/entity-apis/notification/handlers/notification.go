package handlers

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/nuwand/kubecon2022-demo/notification-service/db"
	"github.com/nuwand/kubecon2022-demo/notification-service/helpers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NotificationCreate(ctx context.Context, req db.Notification) (*db.Notification, error) {

	if v, err := helpers.IsValid(req); !v {
		return nil, err
	}
	obj := db.Notification{}
	if err := copier.Copy(&obj, &req); err != nil {
		return nil, err
	}

	if err := obj.Save(ctx); err != nil {
		return nil, err
	}

	return &obj, nil
}

func NotificationList(ctx context.Context, req db.Notification) ([]db.Notification, error) {
	obj := db.Notification{}
	if err := copier.Copy(&obj, &req); err != nil {
		return nil, err
	}

	return db.GetNotificationList(ctx, obj)

}

func NotificationDelete(ctx context.Context, id string) (*db.Notification, error) {

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	obj := db.Notification{ID: objID}

	if err := obj.Get(ctx); err != nil {
		return nil, err
	}

	if err := obj.Delete(ctx); err != nil {
		return nil, err
	}

	return &obj, nil
}

func NotificationGet(ctx context.Context, id string) (*db.Notification, error) {

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	obj := db.Notification{ID: objID}

	if err := obj.Get(ctx); err != nil {
		return nil, err
	}

	return &obj, nil
}

func NotificationUpdate(ctx context.Context, id string, req db.Notification) (*db.Notification, error) {

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	obj := db.Notification{ID: objID}
	req.ID = obj.ID

	if err := obj.Get(ctx); err != nil {
		return nil, err
	}

	if err := copier.Copy(&obj, &req); err != nil {
		return nil, err
	}

	if err := obj.Update(ctx); err != nil {
		return nil, err
	}
	return &obj, nil
}
