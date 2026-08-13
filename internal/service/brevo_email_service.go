package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type brevoEmailService struct {
	apiKey      string
	senderEmail string
	senderName  string
}

// =====================================================
// CONSTRUCTOR
// =====================================================

func NewBrevoEmailService() EmailService {
	return &brevoEmailService{
		apiKey:      os.Getenv("BREVO_API_KEY"),
		senderEmail: os.Getenv("BREVO_SENDER_EMAIL"),
		senderName:  os.Getenv("BREVO_SENDER_NAME"),
	}
}

// =====================================================
// BREVO REQUEST
// =====================================================

type brevoEmailRequest struct {
	Sender      brevoSender `json:"sender"`
	To          []brevoTo   `json:"to"`
	Subject     string      `json:"subject"`
	HTMLContent string      `json:"htmlContent"`
}

type brevoSender struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type brevoTo struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// =====================================================
// BREVO RESPONSE
// =====================================================

type brevoEmailResponse struct {
	MessageID string `json:"messageId"`
}

// =====================================================
// SEND EMAIL
// =====================================================

func (s *brevoEmailService) SendEmail(to string, subject string, htmlContent string) (string, error) {

	// -------------------------------------------------
	// Validate configuration
	// -------------------------------------------------

	if s.apiKey == "" {
		return "", fmt.Errorf("BREVO_API_KEY belum diset")
	}

	if s.senderEmail == "" {
		return "", fmt.Errorf("BREVO_SENDER_EMAIL belum diset")
	}

	// -------------------------------------------------
	// Validate input
	// -------------------------------------------------

	if to == "" {
		return "", fmt.Errorf("recipient email is required")
	}

	if subject == "" {
		return "", fmt.Errorf("email subject is required")
	}

	if htmlContent == "" {
		return "", fmt.Errorf("email content is required")
	}

	// -------------------------------------------------
	// Request body
	// -------------------------------------------------

	requestBody := brevoEmailRequest{
		Sender: brevoSender{
			Email: s.senderEmail,
			Name:  s.senderName,
		},
		To: []brevoTo{
			{
				Email: to,
			},
		},
		Subject:     subject,
		HTMLContent: htmlContent,
	}

	// -------------------------------------------------
	// Encode JSON
	// -------------------------------------------------

	jsonBody, err := json.Marshal(requestBody)

	if err != nil {
		return "", fmt.Errorf(
			"failed to marshal Brevo request: %w",
			err,
		)
	}

	// -------------------------------------------------
	// Create HTTP request
	// -------------------------------------------------

	req, err := http.NewRequest(
		http.MethodPost,
		"https://api.brevo.com/v3/smtp/email",
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return "", fmt.Errorf(
			"failed to create Brevo request: %w",
			err,
		)
	}

	// -------------------------------------------------
	// Headers
	// -------------------------------------------------

	req.Header.Set("accept", "application/json")
	req.Header.Set("api-key", s.apiKey)
	req.Header.Set("content-type", "application/json")

	// -------------------------------------------------
	// Send request
	// -------------------------------------------------

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return "", fmt.Errorf(
			"failed to send email: %w",
			err,
		)
	}

	defer resp.Body.Close()

	// -------------------------------------------------
	// Handle error response
	// -------------------------------------------------

	if resp.StatusCode < http.StatusOK ||
		resp.StatusCode >= http.StatusMultipleChoices {

		var errorResponse map[string]interface{}

		if err := json.NewDecoder(resp.Body).Decode(
			&errorResponse,
		); err != nil {

			return "", fmt.Errorf(
				"Brevo API returned status code %d",
				resp.StatusCode,
			)
		}

		return "", fmt.Errorf(
			"Brevo API error: status %d, response: %v",
			resp.StatusCode,
			errorResponse,
		)
	}

	// -------------------------------------------------
	// Decode successful response
	// -------------------------------------------------

	var response brevoEmailResponse

	if err := json.NewDecoder(resp.Body).Decode(
		&response,
	); err != nil {

		return "", fmt.Errorf(
			"failed to decode Brevo response: %w",
			err,
		)
	}

	// -------------------------------------------------
	// Validate message ID
	// -------------------------------------------------

	if response.MessageID == "" {
		return "", fmt.Errorf(
			"Brevo returned empty message ID",
		)
	}

	return response.MessageID, nil
}
