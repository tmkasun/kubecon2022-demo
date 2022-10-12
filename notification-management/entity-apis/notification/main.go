package main

import (
	"net/http"

	"github.com/nuwand/kubecon2022-demo/notification-service/api"

	"github.com/sirupsen/logrus"
)

// @title Notification Entity API documentation
// @version 1.0.0
// @BasePath /api/v1
func main() {

	routes, err := api.LoadAPI()
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("http server started")
	http.ListenAndServe(":3000", routes)
}
