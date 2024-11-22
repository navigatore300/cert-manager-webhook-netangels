// netangels_dns.go

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type NetangelsAPI struct {
	APIKey string
}

func (n *NetangelsAPI) CreateDNSRecord(domain string, recordType string, recordValue string) error {
	// Создание DNS-записи
	req, err := http.NewRequest("POST", "https://api.netangels.ru/api/v1/dns/records", nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+n.APIKey)
	req.Header.Set("Content-Type", "application/json")

	record := map[string]string{
		"domain": domain,
		"type":   recordType,
		"value":  recordValue,
	}

	jsonRecord, err := json.Marshal(record)
	if err != nil {
		return err
	}

	req.Body = ioutil.NopCloser(bytes.NewReader(jsonRecord))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("Failed to create DNS record: %d", resp.StatusCode)
	}

	return nil
}

func (n *NetangelsAPI) GetToken(domain string) (string, error) {
	// Получение токена
	req, err := http.NewRequest("GET", "https://api.netangels.ru/api/v1/dns/tokens", nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+n.APIKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Failed to get token: %d", resp.StatusCode)
	}

	var token struct {
		Token string `json:"token"`
	}

	err = json.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		return "", err
	}

	return token.Token, nil
}

func (n *NetangelsAPI) DeleteDNSRecord(domain string, recordType string, recordValue string) error {
	// Удаление DNS-записи
	req, err := http.NewRequest("DELETE", "https://api.netangels.ru/api/v1/dns/records", nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+n.APIKey)
	req.Header.Set("Content-Type", "application/json")

	record := map[string]string{
		"domain": domain,
		"type":   recordType,
		"value":  recordValue,
	}

	jsonRecord, err := json.Marshal(record)
	if err != nil {
		return err
	}

	req.Body = ioutil.NopCloser(bytes.NewReader(jsonRecord))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("Failed to delete DNS record: %d", resp.StatusCode)
	}

	return nil
}
