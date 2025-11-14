package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"user-service-ucd/src/app/dto"
)

type NotificationService struct {
	BaseURL string
	Client  *http.Client
}

func NewNotificationService() *NotificationService {
	return &NotificationService{
		BaseURL: os.Getenv("EMAIL_SERVICE_URL"),
		Client:  &http.Client{},
	}
}

func (ns *NotificationService) SendWelcomeEmail(technicianEmail string, loginData dto.LoginRequest, token string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/technicians/register", ns.BaseURL)

	request := dto.SendWelcomeEmailRequest{
		To:       technicianEmail,
		Username: loginData.Username,
		Password: loginData.Password,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := ns.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("notification-service responded with %s", resp.Status)
	}

	var data map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}
