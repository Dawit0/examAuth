package sms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Sender interface {
	SendSMS(phone, message string) error
}

type StubSender struct{}

func (s *StubSender) SendSMS(phone, message string) error {
	fmt.Printf("[SMS] to=%s message=%s\n", phone, message)
	return nil
}

type SendSMSGateSender struct {
	APIKey string
	URL    string
	Client *http.Client
}

func NewSendSMSGateSender(apiKey, url string) *SendSMSGateSender {
	return &SendSMSGateSender{APIKey: apiKey, URL: url, Client: &http.Client{Timeout: 10 * time.Second}}
}

func (s *SendSMSGateSender) SendSMS(phone, message string) error {
	payload := map[string]string{
		"api_key": s.APIKey,
		"to":      phone,
		"message": message,
	}
	b, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", s.URL, bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	resp, err := s.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return fmt.Errorf("sms provider error: %s", resp.Status)
	}
	return nil
}
