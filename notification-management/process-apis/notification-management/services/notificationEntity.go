package services

import (
	"encoding/json"
	"errors"
	"github.com/parnurzeal/gorequest"
	"github.com/sirupsen/logrus"
	"os"
)

type NotificationCreateRequest struct {
	UserEmail string `json:"userEmail,omitempty"`
	Message   string `json:"message,omitempty"`
	Type      string `json:"type,omitempty"`
}

type NotificationCreateResponse struct {
	Data NotificationCreateResponseData `json:"data,omitempty"`
}

type NotificationCreateResponseData struct {
	ID        string `json:"_id,omitempty"`
	UserEmail string `json:"userEmail,omitempty"`
	Message   string `json:"message,omitempty"`
	Type      string `json:"type,omitempty"`
	Read      *bool  `json:"read,omitempty"`
}

func NECreateNotification(req NotificationCreateRequest) (*NotificationCreateResponseData, error) {

	request := gorequest.New()

	resp, body, _ := request.Post(os.Getenv("notificationURL") + "/api/v1/notification").
		Send(req).
		End()

	if resp.StatusCode >= 300 {
		logrus.Infof("NECreateNotification:error %d %s", resp.StatusCode, body)
		return nil, errors.New(body)
	}

	res := NotificationCreateResponse{}
	if err := json.Unmarshal([]byte(body), &res); err != nil {
		return nil, err
	}

	response := res.Data
	return &response, nil
}
