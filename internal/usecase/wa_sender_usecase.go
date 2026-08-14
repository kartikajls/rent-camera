package usecase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"p2-ip-kartikajls/config"
	"p2-ip-kartikajls/internal/dto"
)

type WasenderUsecase interface {
	SendMessage(req dto.SendWARequest) (*dto.SendWAResponse, error)
}

type wasenderUsecase struct{}

func NewWasenderUsecase() WasenderUsecase {
	return &wasenderUsecase{}
}

type wasenderRequest struct {
	To   string `json:"to"`
	Text string `json:"text"`
}

type wasenderResponse struct {
	Success bool `json:"success"`
	Data    struct {
		MsgID  int64  `json:"msgId"`
		JID    string `json:"jid"`
		Status string `json:"status"`
	} `json:"data"`
}

func (u *wasenderUsecase) SendMessage(req dto.SendWARequest) (*dto.SendWAResponse, error) {

	payload := wasenderRequest{
		To:   req.Phone,
		Text: req.Message,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest(
		http.MethodPost,
		config.GetWasenderAPIURL(),
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set(
		"Authorization",
		"Bearer "+config.GetWasenderAPIKey(),
	)

	httpReq.Header.Set(
		"Content-Type",
		"application/json",
	)

	client := &http.Client{}

	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf(
			"WasenderAPI returned status %d: %s",
			resp.StatusCode,
			string(responseBody),
		)
	}

	var result wasenderResponse

	if err := json.Unmarshal(responseBody, &result); err != nil {
		return nil, err
	}

	if !result.Success {
		return nil, fmt.Errorf(
			"WasenderAPI failed to send message",
		)
	}

	return &dto.SendWAResponse{
		MsgID:  result.Data.MsgID,
		JID:    result.Data.JID,
		Status: result.Data.Status,
	}, nil
}
