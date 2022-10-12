package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/nuwand/kubecon2022-demo/notification-service/db"
	"github.com/nuwand/kubecon2022-demo/notification-service/handlers"
)

func NotificationRoutes(r chi.Router) {

	r.Post("/", NotificationCreate)
	r.Get("/", NotificationList)
	r.Delete("/{id}", NotificationList)
	r.Get("/{id}", NotificationGet)
	r.Put("/{id}", NotificationUpdate)

}

// @Summary Create Notification
// @Tags Notification
// @Accept json
// @Produce json
// @Param data body db.Notification	true	"data"
// @Success 200 {object} db.Notification	"Okay"
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /notification [post]
func NotificationCreate(w http.ResponseWriter, r *http.Request) {
	body := db.Notification{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondWithError(w, 400, err.Error())
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	obj, err := handlers.NotificationCreate(ctx, body)
	if err != nil {
		RespondWithError(w, 500, err.Error())
		return
	}
	RespondwithJSON(w, 200, obj)
}

// @Summary Get Notification List
// @Tags Notification
// @Accept json
// @Produce json
// @Param 	userEmail 	query 	string 	false 	"User Email"
// @Param 	userId 	query 	string 	false 	"User ID"
// @Success 200 {object} []db.Notification	"Okay"
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /notification [get]
func NotificationList(w http.ResponseWriter, r *http.Request) {
	var query db.Notification
	if err := decoder.Decode(&query, r.URL.Query()); err != nil {
		RespondWithError(w, 400, err.Error())
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	objList, err := handlers.NotificationList(ctx, query)
	if err != nil {
		RespondWithError(w, 404, err.Error())
		return
	}
	RespondwithJSON(w, 200, objList)

}

// @Summary Delete Notification
// @Tags Notification
// @Accept json
// @Produce json
// @Param	id	path	string	true	"id"
// @Success 200 {object} string	"Okay"
// @Failure 404 {string} string
// @Router /notification/{id} [delete]
func NotificationDelete(w http.ResponseWriter, r *http.Request) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	obj, err := handlers.NotificationDelete(ctx, chi.URLParam(r, "id"))
	if err != nil {
		RespondWithError(w, 404, err.Error())
		return
	}
	RespondwithJSON(w, 200, obj)
}

// @Summary Update Notification
// @Tags Notification
// @Accept json
// @Produce json
// @Param	id	path	string	true	"id"
// @Param data body db.Notification	true	"data"
// @Success 200 {object} db.Notification	"Okay"
// @Failure 404 {string} string
// @Router /notification/{id} [put]
func NotificationUpdate(w http.ResponseWriter, r *http.Request) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	body := db.Notification{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondWithError(w, 400, err.Error())
		return
	}

	obj, err := handlers.NotificationUpdate(ctx, chi.URLParam(r, "id"), body)
	if err != nil {
		RespondWithError(w, 404, err.Error())
		return
	}
	RespondwithJSON(w, 200, obj)
}

// @Summary Get Notification
// @Tags Notification
// @Accept json
// @Produce json
// @Param	id	path	string	true	"id"
// @Success 200 {object} string	"Okay"
// @Failure 404 {string} string
// @Router /notification/{id} [get]
func NotificationGet(w http.ResponseWriter, r *http.Request) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	obj, err := handlers.NotificationGet(ctx, chi.URLParam(r, "id"))
	if err != nil {
		RespondWithError(w, 404, err.Error())
		return
	}
	RespondwithJSON(w, 200, obj)
}
