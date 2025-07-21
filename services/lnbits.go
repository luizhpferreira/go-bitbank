package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var baseURL = "http://localhost:5000"

func CreateUser(username, password, passwordRepeat string) error {

	if password != passwordRepeat {
		return errors.New("passwords do not match")
	}

	userData := map[string]string{
		"username":        username,
		"password":        password,
		"password_repeat": passwordRepeat,
	}

	requestBody, err := json.Marshal(userData)
	if err != nil {
		return fmt.Errorf("error marshaling user data: %w", err)
	}

	url := fmt.Sprintf("%s/users/api/v1/user", baseURL)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJsaHBmIiwiYXV0aF90aW1lIjoxNzUzMTA2MDg1LCJhcGlfdG9rZW5faWQiOiIxN2EwOTZjZTE4NDE0YjEzOGI1NGQ2YmU2MmJhNmQ0NSIsImV4cCI6MTc1NjYwOTE4NX0.yZzqOeqTiMIap5e45l46g9FHHJzFTDqFi8C-VZEY9K0")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error creating user: received status %s", resp.Status)
	}

	return nil
}

func CreateInvoice(invoiceData struct {
	Out    bool   `json:"out"`
	Amount int    `json:"amount"`
	Unit   string `json:"unit"`
	Memo   string `json:"memo"`
}) error {

	requestBody, err := json.Marshal(invoiceData)
	if err != nil {
		return fmt.Errorf("error marshaling invoice data: %w", err)
	}

	url := fmt.Sprintf("%s/api/v1/payments", baseURL)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", "25b4ec3a3c1a40978fcbc11fde123370")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error creating invoice: received status %s", resp.Status)
	}

	return nil
}
