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

func NewBrevoEmailService() EmailService {
	return &brevoEmailService{
		apiKey:      os.Getenv("BREVO_API_KEY"),
		senderEmail: os.Getenv("BREVO_SENDER_EMAIL"),
		senderName:  os.Getenv("BREVO_SENDER_NAME"),
	}
}

// =========================
// Brevo Request
// =========================

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

// =========================
// Brevo Response
// =========================

type brevoEmailResponse struct {
	MessageID string `json:"messageId"`
}

// =========================
// Send Email to Brevo
// =========================

func (s *brevoEmailService) send(
	toEmail string,
	toName string,
	subject string,
	htmlContent string,
) (string, error) {

	if s.apiKey == "" {
		return "", fmt.Errorf("BREVO_API_KEY belum diset")
	}

	if s.senderEmail == "" {
		return "", fmt.Errorf("BREVO_SENDER_EMAIL belum diset")
	}

	requestBody := brevoEmailRequest{
		Sender: brevoSender{
			Email: s.senderEmail,
			Name:  s.senderName,
		},
		To: []brevoTo{
			{
				Email: toEmail,
				Name:  toName,
			},
		},
		Subject:     subject,
		HTMLContent: htmlContent,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal email request: %w", err)
	}

	req, err := http.NewRequest(
		http.MethodPost,
		"https://api.brevo.com/v3/smtp/email",
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return "", fmt.Errorf("failed to create email request: %w", err)
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("api-key", s.apiKey)
	req.Header.Set("content-type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send email: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf(
			"brevo API returned status code: %d",
			resp.StatusCode,
		)
	}

	var response brevoEmailResponse

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", fmt.Errorf(
			"failed to decode Brevo response: %w",
			err,
		)
	}

	return response.MessageID, nil
}

// =========================
// Registration Email
// =========================

func (s *brevoEmailService) SendRegistrationEmail(
	toEmail string,
	username string,
) (string, error) {

	subject := "Registration Confirmation - Rent Camera"

	htmlContent := fmt.Sprintf(`
		<html>
		<body>
			<h2>Welcome to Rent Camera, %s!</h2>

			<p>Your registration has been successfully completed.</p>

			<p>
				Thank you for joining Rent Camera.
			</p>

			<p>
				Regards,<br>
				Rent Camera Team
			</p>
		</body>
		</html>
	`, username)

	return s.send(
		toEmail,
		username,
		subject,
		htmlContent,
	)
}

// =========================
// Booking Confirmation Email
// =========================

func (s *brevoEmailService) SendBookingConfirmationEmail(
	toEmail string,
	username string,
	orderID int64,
) (string, error) {

	subject := "Booking Confirmation - Rent Camera"

	htmlContent := fmt.Sprintf(`
		<html>
		<body>
			<h2>Booking Confirmation</h2>

			<p>Hello %s,</p>

			<p>
				Your camera rental booking has been successfully created.
			</p>

			<p>
				<strong>Rental Order ID:</strong> %d
			</p>

			<p>
				Please check your account for complete booking details.
			</p>

			<p>
				Regards,<br>
				Rent Camera Team
			</p>
		</body>
		</html>
	`, username, orderID)

	return s.send(
		toEmail,
		username,
		subject,
		htmlContent,
	)
}

// =========================
// Payment Confirmation Email
// =========================

func (s *brevoEmailService) SendPaymentConfirmationEmail(
	toEmail string,
	username string,
	orderID int64,
) (string, error) {

	subject := "Payment Confirmation - Rent Camera"

	htmlContent := fmt.Sprintf(`
		<html>
		<body>
			<h2>Payment Confirmation</h2>

			<p>Hello %s,</p>

			<p>
				Your payment has been successfully confirmed.
			</p>

			<p>
				<strong>Rental Order ID:</strong> %d
			</p>

			<p>
				Your rental order is now being processed.
			</p>

			<p>
				Regards,<br>
				Rent Camera Team
			</p>
		</body>
		</html>
	`, username, orderID)

	return s.send(
		toEmail,
		username,
		subject,
		htmlContent,
	)
}
