package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"p2-ip-kartikajls/config"
)

type WhatsAppService interface {
	Send(phone string, message string) error
}

type whatsappService struct {
	apiURL string
	apiKey string
}

func NewWhatsAppService() WhatsAppService {
	return &whatsappService{
		apiURL: config.GetWasenderAPIURL(),
		apiKey: config.GetWasenderAPIKey(),
	}
}

func (s *whatsappService) Send(phone string, message string) error {

	payload := map[string]interface{}{
		"to":   phone,
		"text": message,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		s.apiURL,
		bytes.NewBuffer(body),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(
		"Authorization",
		"Bearer "+s.apiKey,
	)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf(
			"wasender API returned status %d",
			resp.StatusCode,
		)
	}

	return nil
}
