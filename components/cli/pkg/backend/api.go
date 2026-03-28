package backend

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ApiBackend struct {
	BaseURL string
}

func (b *ApiBackend) GetStatus(ctx context.Context) (*MachineStatus, error) {
	resp, err := http.Get(fmt.Sprintf("%s/status", b.BaseURL))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s", resp.Status)
	}

	body, _ := io.ReadAll(resp.Body)
	var status MachineStatus
	err = json.Unmarshal(body, &status)
	return &status, err
}

func (b *ApiBackend) PushButton(ctx context.Context) (string, error) {
	resp, err := http.Post(fmt.Sprintf("%s/push", b.BaseURL), "application/json", nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API error: %s", resp.Status)
	}

	body, _ := io.ReadAll(resp.Body)
	return string(body), nil
}
